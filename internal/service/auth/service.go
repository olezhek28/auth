package auth

import (
	"context"

	"github.com/olezhek28/auth/internal/config"
	userRepository "github.com/olezhek28/auth/internal/repository/user"
)

type Service interface {
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
	GetRefreshToken(ctx context.Context, username string, password string) (string, error)
}

type service struct {
	authConfig config.AuthConfig

	userRepository userRepository.Repository
}

func NewService(authConfig config.AuthConfig, userRepository userRepository.Repository) *service {
	return &service{
		authConfig:     authConfig,
		userRepository: userRepository,
	}
}
