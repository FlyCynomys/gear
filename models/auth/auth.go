package auth

import (
	"time"

	ErrCode "github.com/FlyCynomys/errors"
	Err "github.com/FlyCynomys/tools/err"
	"github.com/FlyCynomys/tools/randomstring"
	"github.com/FlyCynomys/tools/uuid"
	"github.com/astaxie/beego/orm"
)

var IdGenerator, _ = uuid.NewIdWorker(1)

type Auth struct {
	AuthID   int64  `json:"authid,omitempty" orm:"column(authid);pk;auto"`
	UserID   int64  `json:"userid,omitempty" orm:"column(userid)"`
	UrlToken string `json:"url_token,omitempty" orm:"column(url_token)"`
	Password string `json:"password,omitempty" orm:"column(password)"`
	Email    string `json:"email,omitempty" orm:"column(email)"`
	Nickname string `json:"nickname,omitempty" orm:"column(nickname);charset(utf8)"`
	Phone    string `json:"phone,omitempty" orm:"column(phone)"`
	Salt     string `json:"salt,omitempty" orm:"column(salt)"`

	Actived     bool   `json:"active,omitempty" orm:"column(actived)"`
	ActiveEmail string `json:"Active_email,omitempty" orm:"column(active_email)"`

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
	o := orm.NewOrm()
	a.Password = EncodePassword(a.Password, a.Salt)
	_, err := o.Insert(a)
	if err != nil {
		return false, Err.New(-1, err.Error())
	}
	return true, nil
}

func (a *Auth) Get() (bool, *Err.ErrorCode) {
	o := orm.NewOrm()
	password := a.Password
	err := o.Read(a, "email", "deleted")
	if err == orm.ErrNoRows {
		return false, ErrCode.ErrorUserNotExist
	}
	temp := DecodePassword(a.Password, a.Salt)
	if password != temp {
		return false, ErrCode.ErrorAuthPasswordNotMatch
	}
	return true, nil
}

func (a *Auth) GetByAccount(email string) (bool, *Err.ErrorCode) {
	o := orm.NewOrm()
	password := a.Password
	err := o.Read(a, "email", "deleted")
	if err == orm.ErrNoRows {
		return false, ErrCode.ErrorUserNotExist
	}
	temp := DecodePassword(a.Password, a.Salt)
	if password != temp {
		return false, ErrCode.ErrorAuthPasswordNotMatch
	}
	return true, nil
}

func (a *Auth) GetByUid() (bool, *Err.ErrorCode) {
	o := orm.NewOrm()
	password := a.Password
	err := o.Read(a, "userid", "deleted")
	if err == orm.ErrNoRows {
		return false, ErrCode.ErrorUserNotExist
	}
	temp := DecodePassword(a.Password, a.Salt)
	if password != temp {
		return false, ErrCode.ErrorAuthPasswordNotMatch
	}
	return true, nil
}

func (a *Auth) Update() (bool, *Err.ErrorCode) {
	o := orm.NewOrm()
	_, err := o.Update(a, "auto_now")
	if err != nil {
		return false, ErrCode.ErrorAuthInfoUpdateFailed
	}
	return true, nil
}

func (a *Auth) Delete() (bool, *Err.ErrorCode) {
	o := orm.NewOrm()
	_, err := o.Update(a, "authid", "deleted")
	if err != nil {
		if err == orm.ErrNoRows {
			return false, ErrCode.ErrorAuthInfoUpdateFailed
		}
		return false, ErrCode.ErrorAuthInfoDeleteFailed
	}
	return true, nil
}

func DecodePassword(password string, salt string) string {
	return password
}

func EncodePassword(password string, salt string) string {
	return password
}
