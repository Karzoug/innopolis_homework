package app

import (
	"context"

	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"

	"assesment_1/internal/config"
	"assesment_1/internal/delivery/http"
	"assesment_1/internal/model"
	"assesment_1/internal/repo/mapqueue"
	"assesment_1/internal/repo/mxmap"
	"assesment_1/internal/service"
	"assesment_1/internal/token/whitelist"
)

func Run(ctx context.Context, logger zerolog.Logger) error {
	cfg := new(config.Config)
	if err := env.Parse(cfg); err != nil {
		return err
	}

	set := mxmap.New[string, struct{}]()
	tokenValidator := whitelist.New(set)
	tokenValidator.Add("hDFJ!4&5MVg*bTDX") // TODO: move to config or delete

	cache := mapqueue.New[string, model.Message]()

	frService := service.New(cfg.Service, cache, logger)

	srv := http.New(
		cfg.HTTP,
		tokenValidator,
		frService,
		logger,
	)

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return frService.Run(ctx)
	})
	eg.Go(func() error {
		return srv.Run(ctx)
	})

	return eg.Wait()
}
