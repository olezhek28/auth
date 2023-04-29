package note

import (
	"github.com/olezhek28/auth/internal/model"
	desc "github.com/olezhek28/auth/pkg/note_v1"
)

func ToInfo(info *desc.NoteInfo) *model.Info {
	return &model.Info{
		Title:     info.GetTitle(),
		Content:   info.GetContent(),
		CreatedAt: info.GetCreatedAt().AsTime(),
	}
}
