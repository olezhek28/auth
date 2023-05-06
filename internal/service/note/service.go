package note

import (
	"context"

	"github.com/olezhek28/auth/internal/model"
	"github.com/olezhek28/auth/internal/repository/note"
)

var _ Service = (*service)(nil)

type Service interface {
	Create(ctx context.Context, info *model.Info) (int64, error)
	GetList(ctx context.Context) ([]*model.Note, error)
}

type service struct {
	noteRepository note.Repository
}

func NewService(noteRepository note.Repository) *service {
	return &service{
		noteRepository: noteRepository,
	}
}
