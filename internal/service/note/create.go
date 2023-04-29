package note

import (
	"context"

	"github.com/olezhek28/auth/internal/model"
)

func (s *service) Create(ctx context.Context, info *model.Info) (int64, error) {
	id, err := s.noteRepository.Create(ctx, info)
	if err != nil {
		return 0, err
	}

	return id, nil
}
