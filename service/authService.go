package service

import (
	"github.com/FlyCynomys/gear/models/auth"
	"github.com/FlyCynomys/tools/format"
)

const (
	ErrorUserNotExist  = 400001
	ErrorPasswordError = 400002
	ErrorUserHasExist  = 400003
)

type AuthService struct {
}

func (a *AuthService) Login(account string, password string) *Result {
	re := new(Result)
	au := auth.NewAuth()
	au.Email = account
	au.Password = password
	ok, err := au.Get()
	if err != nil {
		re.Status = -1
		re.Data = ok
		re.Description = err.Error()
	} else {
		re.Status = 1
		re.Data = au.UID
		re.Description = "ok"
	}
	return re
}

func (a *AuthService) Logout(account string, cookieinfo string) *Result {
	re := new(Result)
	au := auth.NewAuth()
	au.Email = account
	ok, err := au.Update()
	if err != nil {
		re.Status = -1
		re.Data = ok
		re.Description = err.Error()
		return re
	} else {
		re.Status = 1
		re.Data = ok
		re.Description = "ok"
	}
	return re
}

func (a *AuthService) Register(account, password, nickname string) *Result {
	re := new(Result)
	au := auth.NewAuthWithSalt()
	au.Email = account
	au.Password = password
	data, ok := format.TransToPinyin(nickname)
	if ok != true {
		au.UrlToken = ""
	} else {
		au.UrlToken = data
	}
	ok, err := au.Insert()
	if err != nil {
		re.Status = -1
		re.Data = ok
		re.Description = err.Error()
	} else {
		re.Status = 1
		re.Data = au.UID
		re.Description = "ok"
	}
	return re
}
