package config

import (
	"os"

	"github.com/pkg/errors"
)

var _ HTTPConfig = (*httpConfig)(nil)

const (
	httpHostEnvName = "HTTP_HOST"
)

type HTTPConfig interface {
	Host() string
}

type httpConfig struct {
	host string
}

func NewHTTPConfig() (*httpConfig, error) {
	host := os.Getenv(httpHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("http host not found")
	}

	return &httpConfig{
		host: host,
	}, nil
}

func (cfg *httpConfig) Host() string {
	return cfg.host
}
