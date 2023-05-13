package note_v1

import (
	"context"

	converter "github.com/olezhek28/auth/internal/converter/note"
	desc "github.com/olezhek28/auth/pkg/note_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := i.noteService.Create(ctx, converter.ToInfo(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
