package netdwn

import (
	"runtime"
	"time"
)

const defaultTimeout = 30 * time.Second

type config struct {
	WorkersNumber int
	Timeout       time.Duration
}

func NewConfig(workersNumber int, timeout time.Duration) config {
	if workersNumber <= 0 {
		workersNumber = runtime.NumCPU()
	}
	if timeout <= 0 {
		timeout = defaultTimeout
	}
	return config{
		WorkersNumber: workersNumber,
		Timeout:       timeout,
	}
}
