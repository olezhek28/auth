package access

import "context"

type Service interface {
	Check(ctx context.Context, endpointAddress string) (bool, error)
}

type service struct {
}

func NewService() *service {
	return &service{}
}
