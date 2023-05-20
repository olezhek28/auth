package auth

import (
	"context"

	"github.com/olezhek28/auth/internal/utils"
	"github.com/pkg/errors"
)

func (s *service) GetRefreshToken(ctx context.Context, username string, password string) (string, error) {
	// Лезем в кеш

	//Лезем в базу за данными пользователя
	// Проверяем пароль
	// Генерируем токен

	userInfo, err := s.userRepository.Get(ctx, username)
	if err != nil {
		return "", err
	}

	if !utils.VerifyPassword(userInfo.Password, password) {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateToken(userInfo, s.authConfig.RefreshTokenSecretKey(), s.authConfig.RefreshTokenExpiration())
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
