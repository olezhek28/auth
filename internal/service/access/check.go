package access

import (
	"context"
	"errors"
	"strings"

	"github.com/olezhek28/auth/internal/utils"
	"google.golang.org/grpc/metadata"
)

const authPrefix = "Bearer "

var accessibleRoles map[string]string

func (s *service) Check(ctx context.Context, endpointAddress string) (bool, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false, errors.New("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return false, errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return false, errors.New("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	claims, err := utils.VerifyToken(accessToken, s.authConfig.AccessTokenSecretKey())
	if err != nil {
		return false, errors.New("access token is invalid")
	}

	accessibleMap, err := s.accessibleRoles(ctx)
	if err != nil {
		return false, errors.New("failed to get accessible roles")
	}

	role, ok := accessibleMap[endpointAddress]
	if !ok {
		return true, nil
	}

	if role == claims.Role {
		return true, nil
	}

	return false, errors.New("access denied")
}

func (s *service) accessibleRoles(ctx context.Context) (map[string]string, error) {
	if accessibleRoles == nil {
		accessibleRoles = make(map[string]string)

		accessInfo, err := s.accessRepository.GetList(ctx)
		if err != nil {
			return nil, err
		}

		for _, info := range accessInfo {
			accessibleRoles[info.EndpointAddress] = info.Role
		}
	}

	return accessibleRoles, nil
}
