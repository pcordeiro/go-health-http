package http

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Name    string
	URL     string
	Timeout time.Duration
}

const defaultTimeout = 5 * time.Second

func NewHttpCheck(cfg *Config) func(ctx context.Context) error {
	if cfg.Timeout == 0 {
		cfg.Timeout = defaultTimeout
	}

	return func(ctx context.Context) error {
		req, err := http.NewRequest(http.MethodGet, cfg.URL, nil)
		if err != nil {
			return fmt.Errorf("%s health check failed creating the request: %w", cfg.Name, err)
		}

		ctx, cancel := context.WithTimeout(ctx, cfg.Timeout)
		defer cancel()

		// Inform remote service to close the connection after the transaction is complete
		req.Header.Set("Connection", "close")
		req = req.WithContext(ctx)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("%s health check failed making the request: %w", cfg.Name, err)
		}
		defer res.Body.Close()

		if res.StatusCode >= http.StatusInternalServerError {
			return fmt.Errorf("%s remote service is not available at the moment", cfg.Name)
		}

		return nil
	}
}
