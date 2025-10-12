package pack

import (
	"errors"
	"time"

	"github.com/law-lee/easy_note/kitex_gen/demouser"
	"github.com/law-lee/easy_note/pkg/errno"
)

func BuildBaseResp(err error) *demouser.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *demouser.BaseResp {
	return &demouser.BaseResp{
		StatusCode:    err.ErrCode,
		StatusMessage: err.ErrMsg,
		ServiceTime:   time.Now().Unix(),
	}
}
