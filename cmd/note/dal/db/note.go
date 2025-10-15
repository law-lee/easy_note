package db

import (
	"context"

	"github.com/law-lee/easy_note/pkg/consts"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	UserID  int64  `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (n *Note) TableName() string {
	return consts.NoteTableName
}

// CreateNote create note info
func CreateNote(ctx context.Context, notes []*Note) error {
	if err := DB.WithContext(ctx).Create(notes).Error; err != nil {
		return err
	}
	return nil
}

// MGetNotes multiple get list of note info
func MGetNotes(ctx context.Context, noteIDs []int64) ([]*Note, error) {
	var res []*Note
	if len(noteIDs) == 0 {
		return res, nil
	}
	if err := DB.WithContext(ctx).Where("id IN ?", noteIDs).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}
