package config

import (
	"os"

	"github.com/pkg/errors"
)

var _ SwaggerConfig = (*swaggerConfig)(nil)

const (
	swaggerHostEnvName = "SWAGGER_HOST"
)

type SwaggerConfig interface {
	Host() string
}

type swaggerConfig struct {
	host string
}

func NewSwaggerConfig() (*swaggerConfig, error) {
	host := os.Getenv(swaggerHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("swagger host not found")
	}

	return &swaggerConfig{
		host: host,
	}, nil
}

func (cfg *swaggerConfig) Host() string {
	return cfg.host
}
