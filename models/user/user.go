package user

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	UserID    int64  `json:"userid,omitempty" orm:"column(userid);pk;auto"`
	UrlToken  string `json:"url_token,omitempty" orm:"column(url_token)"`
	NickName  string `json:"nick_name,omitempty" orm:"column(nick_name);charset(utf8)"`
	UserType  string `json:"user_type,omitempty" orm:"column(user_type)"`
	Gender    int    `json:"gender,omitempty" orm:"column(gender)"`
	RealName  string `json:"real_name,omitempty" orm:"column(real_name)"`
	AvatarUrl string `json:"avatar_url,omitempty" orm:"column(avatar_url)"`

	Email string `json:"email,omitempty" orm:"column(email)"`
	Phone string `json:"phone,omitempty" orm:"column(phone)"`

	Deleted bool      `json:"deleted,omitempty" orm:"column(deleted)"`
	Created time.Time `json:"created"  orm:"auto_now_add;type(datetime)"`
	Updated time.Time `json:"updated" orm:"auto_now;type(datetime)"`
}

func NewUser() *User {
	return &User{
		Deleted: false,
	}
}

func (u *User) Insert() (int64, bool, error) {
	o := orm.NewOrm()
	index, err := o.Insert(u)
	if err != nil {
		return -1, false, err
	}
	return index, true, nil
}

func (u *User) Get() (bool, error) {
	o := orm.NewOrm()
	err := o.Read(u, "userid", "deleted")
	if err == orm.ErrNoRows {
		return false, errors.New("user not exist")
	}
	return true, nil
}

func (u *User) Update(colms ...string) (bool, error) {
	if len(colms) <= 0 {
		return true, nil
	}
	o := orm.NewOrm()
	_, err := o.Update(u, colms...)
	if err != nil {
		return false, errors.New("update failed")
	}
	return true, nil
}

func (u *User) Delete() (bool, error) {
	o := orm.NewOrm()
	u.Deleted = true
	_, err := o.Update(u, "deleted")
	if err != nil {
		return false, errors.New("delete user failed")
	}
	return true, nil
}

func (u *User) GetByCondition(condition string) (bool, error) {
	o := orm.NewOrm()
	err := o.Read(u, "condition", "deleted")
	if err == orm.ErrNoRows {
		return false, errors.New("user not exist")
	}
	return true, nil
}

func GetUserByLocation(query string) ([]*User, error) {
	return nil, nil
}

func GetUserByJob(query string) ([]*User, error) {
	return nil, nil
}

func GetUserByCompany(query string) ([]*User, error) {
	return nil, nil
}

func GetUserByNmae(query string) ([]*User, error) {
	return nil, nil
}

func DeleteUserByCondition(query string) (bool, error) {
	return true, nil
}
