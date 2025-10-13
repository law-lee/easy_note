package pack

import (
	"github.com/law-lee/easy_note/cmd/user/dal/db"
	"github.com/law-lee/easy_note/kitex_gen/demouser"
)

func User(u *db.User) *demouser.User {
	if u == nil {
		return nil
	}
	return &demouser.User{
		UserId:   int64(u.ID),
		Username: u.Username,
		Avatar:   "test",
	}
}

func Users(us []*db.User) []*demouser.User {
	res := make([]*demouser.User, 0)
	for _, u := range us {
		if temp := User(u); temp != nil {
			res = append(res, User(u))
		}
	}
	return res
}
