package note

import (
	"github.com/olezhek28/auth/internal/model"
	desc "github.com/olezhek28/auth/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToInfo(info *desc.NoteInfo) *model.Info {
	switch info.GetList().(type) {
	case *desc.NoteInfo_Todo:
	case *desc.NoteInfo_Marked:
	}

	return &model.Info{
		Title:     info.GetTitle(),
		Content:   info.GetContent(),
		CreatedAt: info.GetCreatedAt().AsTime(),
	}
}

func ToNoteDesc(note *model.Note) *desc.Note {
	var updatedAt *timestamppb.Timestamp
	if note.UpdatedAt.Valid {
		updatedAt = timestamppb.New(note.UpdatedAt.Time)
	}

	return &desc.Note{
		Id: note.ID,
		Info: &desc.NoteInfo{
			Title:     note.Info.Title,
			Content:   note.Info.Content,
			CreatedAt: timestamppb.New(note.Info.CreatedAt),
		},
		UpdatedAt: updatedAt,
	}
}

func ToNoteListDesc(notes []*model.Note) []*desc.Note {
	noteList := make([]*desc.Note, 0, len(notes))

	for _, note := range notes {
		noteList = append(noteList, ToNoteDesc(note))
	}

	return noteList
}
