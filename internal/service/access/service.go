package access

import (
	"context"

	"github.com/olezhek28/auth/internal/config"
	accessRepository "github.com/olezhek28/auth/internal/repository/access"
)

type Service interface {
	Check(ctx context.Context, endpointAddress string) (bool, error)
}

type service struct {
	authConfig config.AuthConfig

	accessRepository accessRepository.Repository
}

func NewService(authConfig config.AuthConfig, accessRepository accessRepository.Repository) *service {
	return &service{
		authConfig:       authConfig,
		accessRepository: accessRepository,
	}
}
