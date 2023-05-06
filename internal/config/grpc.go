package config

import (
	"os"

	"github.com/pkg/errors"
)

var _ GRPCConfig = (*grpcConfig)(nil)

const (
	grpcHostEnvName = "GRPC_HOST"
)

type GRPCConfig interface {
	Host() string
}

type grpcConfig struct {
	host string
}

func NewGRPCConfig() (*grpcConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	return &grpcConfig{
		host: host,
	}, nil
}

func (cfg *grpcConfig) Host() string {
	return cfg.host
}
