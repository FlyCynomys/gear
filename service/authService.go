package service

import (
	"github.com/FlyCynomys/gear/models/auth"
	"github.com/FlyCynomys/gear/models/user"
	"github.com/FlyCynomys/tools/format"
	"github.com/FlyCynomys/tools/log"
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
		re.Data = au.UserID
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
	au.Nickname = nickname
	urltoken, ok := format.TransToPinyin(nickname)
	if ok != true {
		au.UrlToken = ""
	} else {
		au.UrlToken = urltoken
	}
	newuser := user.NewUser()
	CopyAuthInfo2UserInfo(au, newuser)
	uid, createUserOk, err := newuser.Insert()
	if createUserOk == false {
		log.Error("create user error : ", err)
		re.Status = -1
		re.Data = ok
		re.Description = err.Error()
		return re
	}
	au.UserID = uid
	ok, err = au.Insert()
	if !ok {
		re.Status = -1
		re.Data = ok
		re.Description = err.Error()
		return re
	}
	re.Status = 1
	re.Data = au.UserID
	re.Description = "ok"
	return re
}

func CopyAuthInfo2UserInfo(authinfo *auth.Auth, userinfo *user.User) {
	userinfo.Email = authinfo.Email
	userinfo.NickName = authinfo.Nickname
	userinfo.Gender = 0 //1,2,3,4
	userinfo.UrlToken = authinfo.UrlToken
}
