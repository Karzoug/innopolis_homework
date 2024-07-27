package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog"

	"assesment_1/internal/model"
	"assesment_1/internal/token"
)

const (
	shutdownTimeout   = 5 * time.Second
	readHeaderTimeout = 2 * time.Second
	writeTimeout      = 5 * time.Second
)

type FileMessageService interface {
	Create(ctx context.Context, msg model.Message) error
}

type server struct {
	srv    *http.Server
	logger zerolog.Logger
}

func New(cfg Config, tv token.Validator, service FileMessageService, logger zerolog.Logger) *server {
	logger = logger.With().Str("from", "http server").Logger()

	handler := &handler{
		tv:      tv,
		service: service,
		logger:  logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /file-records", handler.CreateFileMessage)

	return &server{
		srv: &http.Server{
			Addr:              fmt.Sprint(":", cfg.Port),
			Handler:           mux,
			ReadHeaderTimeout: readHeaderTimeout,
			WriteTimeout:      writeTimeout,
		},
		logger: logger,
	}
}

func (s *server) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()

		ctxShutdown, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := s.srv.Shutdown(ctxShutdown); err != nil {
			s.logger.Error().Err(err).Msg("shutdown server error")
		}
	}()

	s.logger.Info().Str("address", s.srv.Addr).Msg("starting listening")
	if err := s.srv.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	s.logger.Info().Str("address", s.srv.Addr).Msg("stopped listening")
	return nil
}
