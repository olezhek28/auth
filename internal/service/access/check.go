package access

import "context"

func (s *service) Check(ctx context.Context, endpointAddress string) (bool, error) {
	return false, nil
}
