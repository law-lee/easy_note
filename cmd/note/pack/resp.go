package pack

import (
	"errors"
	"time"

	"github.com/law-lee/easy_note/kitex_gen/demonote"
	"github.com/law-lee/easy_note/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *demonote.BaseResp {
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
func baseResp(err errno.ErrNo) *demonote.BaseResp {
	return &demonote.BaseResp{
		StatusCode:    err.ErrCode,
		StatusMessage: err.ErrMsg,
		ServiceTime:   time.Now().Unix(),
	}
}
