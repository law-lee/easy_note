package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/law-lee/easy_note/cmd/user/dal/db"
	"github.com/law-lee/easy_note/kitex_gen/demouser"
	"github.com/law-lee/easy_note/pkg/errno"
)

type CreateUserService struct {
	ctx context.Context
}

func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

func (s *CreateUserService) CreateUser(req *demouser.CreateUserRequest) error {
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return err
	}

	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	return db.CreateUser(s.ctx, []*db.User{{
		Username: req.Username,
		Password: password,
	}})
}
