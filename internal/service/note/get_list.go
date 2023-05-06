package note

import (
	"context"

	"github.com/olezhek28/auth/internal/model"
)

func (s *service) GetList(ctx context.Context) ([]*model.Note, error) {
	res, err := s.noteRepository.GetList(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}
