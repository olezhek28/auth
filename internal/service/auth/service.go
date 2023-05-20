package auth

import (
	"context"

	userRepository "github.com/olezhek28/auth/internal/repository/user"
)

type Service interface {
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
	GetRefreshToken(ctx context.Context, username string, password string) (string, error)
}

type service struct {
	userRepository userRepository.Repository
}

func NewService(userRepository userRepository.Repository) *service {
	return &service{
		userRepository: userRepository,
	}
}
