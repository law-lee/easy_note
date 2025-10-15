package service

import (
	"context"

	"github.com/law-lee/easy_note/cmd/note/dal/db"
	"github.com/law-lee/easy_note/kitex_gen/demonote"
)

type CreateNoteService struct {
	ctx context.Context
}

// NewCreateNoteService new CreateNoteService
func NewCreateNoteService(ctx context.Context) *CreateNoteService {
	return &CreateNoteService{ctx: ctx}
}

// CreateNote create note
func (s *CreateNoteService) CreateNote(req *demonote.CreateNoteRequest) error {
	noteModel := &db.Note{
		UserID:  req.UserId,
		Title:   req.Title,
		Content: req.Content,
	}
	return db.CreateNote(s.ctx, []*db.Note{noteModel})
}
