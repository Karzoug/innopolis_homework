package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"assesment_1/internal/app"

	"github.com/rs/zerolog"
)

var loggerLevel = zerolog.InfoLevel

func init() {
	ls := flag.String("log", "info", "log level (trace/debug/info/warn/error/fatal/panic)")
	flag.Parse()

	l, err := zerolog.ParseLevel(*ls)
	if err != nil {
		log.Fatal(err)
	}
	loggerLevel = l
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM)
	defer stop()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(loggerLevel)

	logger := zerolog.New(os.Stderr).
		Level(loggerLevel).
		With().
		Timestamp().
		Logger()

	if err := app.Run(ctx, logger); err != nil {
		logger.Error().Err(err).Msg("app run error")
	}
}
