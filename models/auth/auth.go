package auth

import (
	"errors"

	"github.com/astaxie/beego/orm"
)

var o orm.Ormer

func Init() {
	o = orm.NewOrm()
}

type Auth struct {
	ID       string `json:"id,omitempty"`
	UID      string `json:"uid,omitempty"`
	UrlToken string `json:"url_token,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Salt     string `json:"salt,omitempty"`
}

func (a *Auth) Insert() (bool, error) {
	err := o.Read(a, "email")
	if err == orm.ErrNoRows {
		a.Password = EncodePassword(a.Password, a.Salt)
		_, err = o.Insert(a)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, errors.New("user exist")
}

func (a *Auth) Get() (bool, error) {
	password := a.Password

	err := o.Read(a, "email")
	if err == orm.ErrNoRows {
		return false, errors.New("user not exist")
	}
	temp := DecodePassword(a.Password, a.Salt)
	if password != temp {
		return false, errors.New("auth failed")
	}
	return true, nil
}

func DecodePassword(password string, salt string) string {
	return password
}

func EncodePassword(password string, salt string) string {
	return password
}
