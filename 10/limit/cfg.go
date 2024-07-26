package limit

import (
	"time"

	r "golang.org/x/time/rate"
)

const (
	defaultCleanupTimeout  time.Duration = 5 * time.Minute
	defaultLimiterLifetime time.Duration = 1 * time.Minute
)

type config struct {
	rate            r.Limit
	burst           int
	cleanupTimeout  time.Duration
	limiterLifetime time.Duration
}

func NewConfig(rate, burst int, opts ...Option) config {
	cfg := config{
		rate:            r.Limit(rate),
		burst:           burst,
		cleanupTimeout:  defaultCleanupTimeout,
		limiterLifetime: defaultLimiterLifetime,
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	return cfg
}

type Option func(*config)

func WithCleanupTimeout(t time.Duration) Option {
	return func(c *config) {
		c.cleanupTimeout = t
	}
}

func WithLimiterLifetime(t time.Duration) Option {
	return func(c *config) {
		c.limiterLifetime = t
	}
}
