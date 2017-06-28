package user

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

//UserSetting setting by UserSetting set
//用户的配置只有从属关系，没有删除功能，但是保留删除字段
type UserSetting struct {
	ID     int64 `json:"id" orm:"column(id);pk;auto"`
	UserID int64 `json:"userid" orm:"column(userid)`

	IsOrg       bool   `json:"is_org,omitempty" orm:"column(is_org)"`
	Description string `json:"description,omitempty" orm:"column(description)"`
	CoverUrl    string `json:"cover_url,omitempty" orm:"column(cover_url)"`
	Headline    string `json:"headline,omitempty" orm:"column(headline)"`

	Deleted bool      `json:"deleted,omitempty" orm:"column(deleted)"`
	Created time.Time `json:"created"  orm:"auto_now_add;type(datetime)"`
	Updated time.Time `json:"updated" orm:"auto_now;type(datetime)"`
}

func NewUserSetting() *UserSetting {
	return &UserSetting{
		Deleted: false,
	}
}

func (u *UserSetting) Insert() (bool, error) {
	o := orm.NewOrm()
	_, err := o.Insert(u)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *UserSetting) Get() (bool, error) {
	o := orm.NewOrm()
	err := o.Read(u, "uid", "deleted")
	if err == orm.ErrNoRows {
		return false, errors.New("UserSetting not exist")
	}
	return true, nil
}

func (u *UserSetting) Update(colms ...string) (bool, error) {
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

func (u *UserSetting) Delete() (bool, error) {
	o := orm.NewOrm()
	u.Deleted = true
	_, err := o.Update(u, "deleted")
	if err != nil {
		return false, errors.New("delete UserSetting failed")
	}
	return true, nil
}

func (u *UserSetting) GetByCondition(condition string) (bool, error) {
	o := orm.NewOrm()
	err := o.Read(u, "condition", "deleted")
	if err == orm.ErrNoRows {
		return false, errors.New("UserSetting not exist")
	}
	return true, nil
}
