package config

import (
	"encoding/base64"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

const (
	accessTokenSecretKeyEnvName  = "ACCESS_TOKEN_SECRET_KEY"
	accessTokenExpirationEnvName = "ACCESS_TOKEN_EXPIRATION_MINUTES"

	refreshTokenSecretKeyEnvName  = "REFRESH_TOKEN_SECRET_KEY"
	refreshTokenExpirationEnvName = "REFRESH_TOKEN_EXPIRATION_MINUTES"
)

type AuthConfig interface {
	RefreshTokenSecretKey() []byte
	AccessTokenSecretKey() []byte

	RefreshTokenExpiration() time.Duration
	AccessTokenExpiration() time.Duration
}

type authConfig struct {
	refreshTokenSecretKey []byte
	accessTokenSecretKey  []byte

	refreshTokenExpiration time.Duration
	accessTokenExpiration  time.Duration
}

func NewAuthConfig() (*authConfig, error) {
	refreshTokenSecretKey, err := decode(os.Getenv(refreshTokenSecretKeyEnvName))
	if err != nil {
		return nil, err
	}

	accessTokenSecretKey, err := decode(os.Getenv(accessTokenSecretKeyEnvName))
	if err != nil {
		return nil, err
	}

	refreshTokenExpiration, err := strconv.Atoi(os.Getenv(refreshTokenExpirationEnvName))
	if err != nil {
		return nil, errors.Wrap(err, "refresh token expired invalid")
	}

	accessTokenExpiration, err := strconv.Atoi(os.Getenv(accessTokenExpirationEnvName))
	if err != nil {
		return nil, errors.Wrap(err, "access token expired invalid")
	}

	return &authConfig{
		refreshTokenSecretKey:  refreshTokenSecretKey,
		accessTokenSecretKey:   accessTokenSecretKey,
		refreshTokenExpiration: time.Minute * time.Duration(refreshTokenExpiration),
		accessTokenExpiration:  time.Minute * time.Duration(accessTokenExpiration),
	}, nil
}

func decode(key string) ([]byte, error) {
	if len(key) == 0 {
		return nil, errors.New("access token private key uri not found")
	}

	decodeKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode access token private key")
	}

	return decodeKey, nil
}

func (cfg *authConfig) RefreshTokenSecretKey() []byte {
	return cfg.refreshTokenSecretKey
}

func (cfg *authConfig) AccessTokenSecretKey() []byte {
	return cfg.accessTokenSecretKey
}

func (cfg *authConfig) RefreshTokenExpiration() time.Duration {
	return cfg.refreshTokenExpiration
}

func (cfg *authConfig) AccessTokenExpiration() time.Duration {
	return cfg.accessTokenExpiration
}
