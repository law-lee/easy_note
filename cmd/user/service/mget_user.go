package service

import (
	"context"

	"github.com/law-lee/easy_note/cmd/user/dal/db"
	"github.com/law-lee/easy_note/cmd/user/pack"
	"github.com/law-lee/easy_note/kitex_gen/demouser"
)

type MGetUserService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

// MGetUser get user info by ids
func (s *MGetUserService) MGetUser(req *demouser.MGetUserRequest) ([]*demouser.User, error) {
	modelUser, err := db.MGetUsers(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}

	return pack.Users(modelUser), nil
}
