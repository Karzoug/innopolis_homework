package service

import (
	"context"
	"runtime"
	"sync"
	"time"

	"github.com/rs/zerolog"

	"assesment_1/internal/model"
	"assesment_1/internal/repo"
	"assesment_1/internal/repo/mxmap"
)

type cacheState int

const (
	cacheOK cacheState = iota
	cacheBusy
	cacheHeavy
)

var (
	cacheKeysOKLen   int = 20 * runtime.NumCPU() // TODO: need to be configurable, need benchmarks
	cacheKeysBusyLen int = 50 * runtime.NumCPU()
)

type service struct {
	cfg           Config
	cache         repo.MapQueues[string, model.Message]
	lockedKeysSet repo.CMap[string, struct{}] // key is the path to the file
	logger        zerolog.Logger
}

func New(cfg Config, cache repo.MapQueues[string, model.Message], logger zerolog.Logger) *service {
	logger = logger.With().Str("from", "service").Logger()
	fixConfig(&cfg, logger)

	return &service{
		cfg:           cfg,
		cache:         cache,
		lockedKeysSet: mxmap.New[string, struct{}](),
		logger:        logger,
	}
}

func (s *service) Run(ctx context.Context) error {
	wg := &sync.WaitGroup{}
	s.logger.Info().
		Int("workers count", s.cfg.MinWorkerCount).
		Msg("starting workers")
	for i := 0; i < s.cfg.MinWorkerCount; i++ {
		wg.Add(1)
		go s.doWorker(ctx, wg)
	}

	go s.doScaling(ctx, wg)

	wg.Wait()
	s.logger.Info().Msg("stopped all workers")
	return nil
}

// doScaling scales worker count if queue state changes.
// Algorithm based on count of files in cache, not messages.
func (s *service) doScaling(ctx context.Context, wg *sync.WaitGroup) {
	prevCacheState := cacheOK
	cancels := make([]context.CancelFunc, 0, s.cfg.MaxWorkerCount-s.cfg.MinWorkerCount)

	t := time.NewTicker(s.cfg.WorkerInterval / 2)
	for {
		select {
		case <-ctx.Done():
			// any worker that is running in this function will be canceled,
			// because parent ctx is canceled,
			// so it's safe to return
			return
		case <-t.C:
		}

		currState := s.getCacheState()
		if prevCacheState == currState {
			continue
		}

		// how many workers should be added or removed
		diffWorkerCount := int(currState-prevCacheState) * ((s.cfg.MaxWorkerCount - s.cfg.MinWorkerCount) / (int(cacheHeavy - cacheOK)))
		s.logger.Info().
			Int("prev worker count", len(cancels)+s.cfg.MinWorkerCount).
			Int("new worker count", len(cancels)+s.cfg.MinWorkerCount+diffWorkerCount).
			Msg("run scaling workers")
		if diffWorkerCount < 0 {
			left := len(cancels) + diffWorkerCount
			for i := len(cancels) - 1; i >= left; i-- {
				cancels[i]()
			}
			cancels = cancels[:left]
		} else {
			for i := 0; i < diffWorkerCount; i++ {
				ctx, cancel := context.WithCancel(ctx)
				cancels = append(cancels, cancel)

				wg.Add(1)
				go s.doWorker(ctx, wg)
			}
		}
		prevCacheState = currState
	}
}

func (s *service) getCacheState() cacheState {
	count := s.cache.Len()

	switch {
	case count <= cacheKeysOKLen:
		return cacheOK
	case count <= cacheKeysBusyLen:
		return cacheBusy
	default:
		return cacheHeavy
	}
}
