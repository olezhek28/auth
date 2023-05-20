package access_v1

import (
	"github.com/olezhek28/auth/internal/service/access"
	desc "github.com/olezhek28/auth/pkg/access_v1"
)

type Implementation struct {
	desc.UnimplementedAccessV1Server

	accessService access.Service
}

func NewImplementation(accessService access.Service) *Implementation {
	return &Implementation{
		accessService: accessService,
	}
}
