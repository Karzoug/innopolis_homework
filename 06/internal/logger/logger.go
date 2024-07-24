package logger

import "time"

var TimeFieldFormat = time.RFC3339

type Logger interface {
	Print(...any)
}
