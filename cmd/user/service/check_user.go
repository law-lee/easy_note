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

type CheckUserService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{ctx: ctx}
}

// CheckUser check user by username and password
func (s *CheckUserService) CheckUser(req *demouser.CheckUserRequest) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.Username
	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.AuthorizationFailedErr
	}
	u := users[0]
	if u.Password != password {
		return 0, errno.AuthorizationFailedErr
	}
	return int64(u.ID), nil
}
