package service

import (
	"context"
	"os"
	"path/filepath"
	"sync"
	"time"

	"assesment_1/internal/model"

	"github.com/cenkalti/backoff/v4"
)

const defaultInitialInterval = 50 * time.Millisecond

// doWorker processes messages from the queue
// and writes the data to the file
func (s *service) doWorker(ctx context.Context, wg *sync.WaitGroup) {
	defer func() {
		s.logger.Trace().Msg("worker stopped")
		wg.Done()
	}()
	s.logger.Trace().Msg("worker started")

	t := time.NewTicker(s.cfg.WorkerInterval)
	doneFlag := false

	for {
		select {
		case <-ctx.Done():
			doneFlag = true
		case <-t.C:
		}

		writeFn := func(key string, arr []model.Message) error {
			defer func() {
				s.lockedKeysSet.Delete(key)
			}()

			if len(arr) == 0 {
				return nil
			}

			// create dir if not exists
			dir := filepath.Dir(arr[0].FileID)
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				if err := os.MkdirAll(dir, 0750); err != nil {
					return err
				}
			}

			file, err := os.OpenFile(arr[0].FileID, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				return err
			}
			defer file.Close()

			for _, v := range arr {
				// try to write to the file, retry if necessary,
				// if error is still occurs log it
				operation := func() error {
					if _, err := file.WriteString(v.Data); err != nil {
						return err
					}
					_, _ = file.Write([]byte("\n"))
					return nil
				}
				err = backoff.Retry(operation,
					backoff.NewExponentialBackOff(
						backoff.WithInitialInterval(defaultInitialInterval),
						backoff.WithMaxElapsedTime(s.cfg.MaxWaitTimeoutToWriteFile),
					),
				)
				if err != nil {
					// here we lost one message to write to the file
					// TODO: try to save it
					return err
				}
			}
			return nil
		}

		for _, key := range s.cache.Keys() {
			if _, ok := s.lockedKeysSet.LoadOrStore(key, struct{}{}); ok {
				continue
			}
			arr, _ := s.cache.LRange(key, -1)
			if err := writeFn(key, arr); err != nil {
				s.logger.Error().Err(err).Msg("processing write file message error")
			}
		}

		if doneFlag {
			return
		}
	}
}
