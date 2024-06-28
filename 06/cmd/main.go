package main

import (
	"06/internal/logger/best"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var TimeFieldFormat = time.RFC3339

func main() {
	numberInt := int(32)
	numberFloat := float64(64.3333333333)

	l := best.NewLogger(os.Stdout)
	l.Print("Hello, World!", numberInt, numberFloat)
	l.Debug().Int("age", numberInt).Float64("weight", numberFloat).Msg("Hello, World!")

	zerolog.TimeFieldFormat = TimeFieldFormat
	lz := zerolog.New(os.Stdout).With().Timestamp().Logger()
	lz.Print("Hello, World!", numberInt, numberFloat)
	lz.Debug().Int("age", numberInt).Float64("weight", numberFloat).Msg("Hello, World!")
}
