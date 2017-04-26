package auth

import (
	"time"

	Err "github.com/FlyCynomys/tools/err"
	"github.com/FlyCynomys/tools/randomstring"
	"github.com/astaxie/beego/orm"
)

var o orm.Ormer

func Init() {
	o = orm.NewOrm()
	o.Using("default")
}

type Auth struct {
	ID       string `json:"id,omitempty"`
	UID      string `json:"uid,omitempty"`
	UrlToken string `json:"url_token,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Salt     string `json:"salt,omitempty"`

	Deleted bool `json:"deleted,omitempty"`

	Created time.Time `json:"created"  orm:"auto_now_add;type(datetime)"`
	Updated time.Time `json:"updated" orm:"auto_now;type(datetime)"`
}

func NewAuth() *Auth {
	return &Auth{
		Deleted: false,
	}
}

func NewAuthWithSalt() *Auth {
	return &Auth{
		Salt:    string(randomstring.RandomString()),
		Deleted: false,
	}
}

func (a *Auth) Insert() (bool, *Err.ErrorCode) {
	err := o.Read(a, "email", "deleted")
	if err == orm.ErrNoRows {
		a.Password = EncodePassword(a.Password, a.Salt)
		_, err = o.Insert(a)
		if err != nil {
			return false, Err.New(-1, err.Error())
		}
		return true, nil
	}
	return false, Err.New(-1, "user is exist")
}

func (a *Auth) Get() (bool, *Err.ErrorCode) {
	password := a.Password
	err := o.Read(a, "email", "deleted")
	if err == orm.ErrNoRows {
		return false, Err.New(-1, "user not exist")
	}
	temp := DecodePassword(a.Password, a.Salt)
	if password != temp {
		return false, Err.New(-1, "password not right")
	}
	return true, nil
}

func (a *Auth) Update() (bool, *Err.ErrorCode) {
	_, err := o.Update(a, "auto_now")
	if err != nil {
		return false, Err.New(-1, "user update failed")
	}
	return true, nil
}

func (a *Auth) Delete() (bool, *Err.ErrorCode) {
	_, err := o.Update(a, "aid", "deleted")
	if err != nil {
		if err == orm.ErrNoRows {
			return false, Err.New(-1, "user not exist")
		}
		return false, Err.New(-1, "user delete failed")
	}
	return true, nil
}

func DecodePassword(password string, salt string) string {
	return password
}

func EncodePassword(password string, salt string) string {
	return password
}
