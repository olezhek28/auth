package interceptor

import (
	"context"

	"github.com/olezhek28/auth/internal/metric"
	"google.golang.org/grpc"
)

func MetricsInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	res, err := handler(ctx, req)
	if err != nil {
		metric.IncRequestTotal(info.FullMethod, "error")
	} else {
		metric.IncRequestTotal(info.FullMethod, "success")
	}

	return res, err
}
