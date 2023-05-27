package auth

import (
	"context"

	"github.com/olezhek28/auth/internal/sys"
	"github.com/olezhek28/auth/internal/sys/codes"
	"github.com/olezhek28/auth/internal/utils"
)

func (s *service) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := utils.VerifyToken(refreshToken, s.authConfig.RefreshTokenSecretKey())
	if err != nil {
		return "", sys.NewCommonError("invalid refresh token", codes.Aborted)
	}

	userInfo, err := s.userRepository.Get(ctx, claims.Username)
	if err != nil {
		return "", err
	}

	accessToken, err := utils.GenerateToken(userInfo, s.authConfig.AccessTokenSecretKey(), s.authConfig.AccessTokenExpiration())
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
