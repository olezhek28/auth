package note_v1

import (
	"context"

	converter "github.com/olezhek28/auth/internal/converter/note"
	desc "github.com/olezhek28/auth/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) GetList(ctx context.Context, _ *emptypb.Empty) (*desc.GetListResponse, error) {
	res, err := i.noteService.GetList(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.GetListResponse{
		Notes: converter.ToNoteListDesc(res),
	}, nil
}
