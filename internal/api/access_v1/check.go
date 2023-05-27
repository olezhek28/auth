package access_v1

import (
	"context"

	desc "github.com/olezhek28/auth/pkg/access_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Check(ctx context.Context, req *desc.CheckRequest) (*desc.CheckResponse, error) {
	isAllowed, err := i.accessService.Check(ctx, req.GetEndpointAddress())
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "access denied")
	}

	return &desc.CheckResponse{IsAllowed: isAllowed}, nil
}
