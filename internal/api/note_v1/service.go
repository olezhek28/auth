package note_v1

import (
	"github.com/olezhek28/auth/internal/service/note"
	desc "github.com/olezhek28/auth/pkg/note_v1"
)

type Implementation struct {
	desc.UnimplementedNoteV1Server

	noteService note.Service
}

func NewImplementation(noteService note.Service) *Implementation {
	return &Implementation{
		noteService: noteService,
	}
}
