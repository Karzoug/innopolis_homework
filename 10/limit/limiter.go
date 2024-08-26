package limit

import (
	"sync"
	"time"

	r "golang.org/x/time/rate"
)

type visitor struct {
	limiter  *r.Limiter
	lastSeen time.Time
}

type limiter struct {
	visitors map[string]*visitor
	mu       sync.Mutex
	cfg      config
}

func newLimiter(cfg config) *limiter {
	l := &limiter{
		visitors: make(map[string]*visitor),
		cfg:      cfg,
	}
	go l.cleanup()

	return l
}

func (l *limiter) getVisitor(ip string) *r.Limiter {
	l.mu.Lock()
	defer l.mu.Unlock()

	v, exists := l.visitors[ip]
	if !exists {
		limiter := r.NewLimiter(l.cfg.rate, l.cfg.burst)
		l.visitors[ip] = &visitor{limiter, time.Now()}
		return limiter
	}

	v.lastSeen = time.Now()
	return v.limiter
}

func (l *limiter) cleanup() {
	for {
		time.Sleep(l.cfg.cleanupTimeout)

		l.mu.Lock()
		for ip, v := range l.visitors {
			if time.Since(v.lastSeen) > l.cfg.limiterLifetime {
				delete(l.visitors, ip)
			}
		}
		l.mu.Unlock()
	}
}
