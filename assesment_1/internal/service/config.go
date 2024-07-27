package service

import (
	"runtime"
	"time"

	"github.com/rs/zerolog"
)

type Config struct {
	MinWorkerCount            int           `env:"MIN_WORKER_COUNT"`
	MaxWorkerCount            int           `env:"MAX_WORKER_COUNT"`
	WorkerInterval            time.Duration `env:"WORKER_INTERVAL" envDefault:"1s"`
	MaxWaitTimeoutToWriteFile time.Duration `env:"MAX_WAIT_TIMEOUT_TO_WRITE_FILE" envDefault:"10s"`
}

func fixConfig(cfg *Config, logger zerolog.Logger) {
	if cfg.MinWorkerCount <= 0 {
		cfg.MinWorkerCount = runtime.NumCPU()
		logger.Info().
			Int("min worker count", cfg.MinWorkerCount).
			Msg("config: set to default value")
	}

	if cfg.MaxWorkerCount <= 0 {
		cfg.MaxWorkerCount = runtime.NumCPU() * 2
		logger.Info().
			Int("max worker count", cfg.MaxWorkerCount).
			Msg("service config: set to default value")
	}
	if cfg.WorkerInterval <= 0 {
		cfg.WorkerInterval = 1 * time.Second
		logger.Warn().
			Dur("worker interval", cfg.WorkerInterval).
			Msg("service config: set to default value")
	}
	if cfg.MaxWaitTimeoutToWriteFile <= 0 {
		cfg.MaxWaitTimeoutToWriteFile = 10 * time.Second
		logger.Warn().
			Dur("max wait timeout to write file", cfg.MaxWaitTimeoutToWriteFile).
			Msg("service config: set to default value")
	}
}
