package netdwn

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

var ErrBadUrl = errors.New("bad url")

func Download(ctx context.Context, urls []string, cfg config) error {
	var errExists bool
	for _, u := range urls {
		_, err := url.ParseRequestURI(u)
		if err != nil {
			errExists = true
			log.Err(err).Msg("bad url")
		}
	}
	if errExists {
		return ErrBadUrl
	}

	var eg errgroup.Group
	eg.SetLimit(cfg.WorkersNumber)

	client := &http.Client{}

	for _, url := range urls {
		eg.Go(func() error {
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()

			logger := log.With().Str("url", url).Logger()

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				logger.Warn().
					Err(err).
					Msg("failed to create http request")
				return nil
			}

			filename := path.Base(url)
			f, err := createFile(filename, logger)
			if err != nil {
				logger.Warn().
					Err(err).
					Msg("failed to create file")
				return nil
			}
			defer f.Close()

			logger.Info().Msg("download started")

			resp, err := client.Do(req)
			if err != nil {
				logger.Warn().
					Err(err).
					Msg("download failed")
				return nil
			}
			defer resp.Body.Close()

			time.AfterFunc(cfg.Timeout, cancel)

			if _, err := io.Copy(f, resp.Body); err != nil {
				logger.Warn().
					Err(err).
					Msg("download failed")
				return nil
			}
			logger.Info().Msg("download finished")

			return nil
		})
	}
	return eg.Wait()
}

func createFile(filename string, logger zerolog.Logger) (*os.File, error) {
	if _, err := os.Stat(filename); nil == err {
		logger.Warn().
			Str("filename", filename).
			Msg("file already exists")

		// try to create a new file with a different name
		for i := 1; ; i++ {
			ext := path.Ext(filename)
			newFilename := fmt.Sprintf("%s-%d%s", filename[:len(filename)-len(ext)], i, ext)
			if _, err := os.Stat(newFilename); nil == err {
				continue
			}
			f, err := os.Create(newFilename)
			if err != nil {
				return nil, err
			}
			return f, nil
		}
	}
	return os.Create(filename)
}
