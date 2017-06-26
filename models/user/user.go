package user

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	ID        int64  `json:"id,omitempty" orm:"column(id);pk;auto"`
	UID       int64  `json:"uid,omitempty" orm:"column(uid)"`
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

func (u *User) Insert() (bool, error) {
	if u == nil {
		return false, errors.New("empty object")
	}
	o := orm.NewOrm()
	u.Deleted = false
	err := o.Read(u, "uid", "deleted")
	if err == orm.ErrNoRows {
		_, err := o.Insert(u)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return true, nil
}

func (u *User) Get() (bool, error) {
	o := orm.NewOrm()
	err := o.Read(u, "uid", "deleted")
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
	var temp = NewUser()
	temp.UID = u.UID
	err := o.Read(temp, "uid", "deleted")
	if err == orm.ErrNoRows {
		return false, errors.New("user not exist")
	}
	_, err = o.Update(u, colms...)
	if err != nil {
		return false, errors.New("update failed")
	}
	return true, nil
}

func (u *User) Delete() (bool, error) {
	o := orm.NewOrm()
	err := o.Read(u, "uid", "deleted")
	if err == orm.ErrNoRows {
		return true, nil
	}
	_, err = o.Update(u, "deleted")
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
