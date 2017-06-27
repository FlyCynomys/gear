package license

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

type License struct {
	ID          int64  `json:"id,omitempty" orm:"column(id);pk;auto"`
	LID         int64  `json:"lid,omitempty" orm:"column(lid)"`
	LicenseName string `json:"license_name,omitempty" orm:"column(license_name);charset(utf8)"`
	Label       string `json:"label,omitempty" orm:"column(label);charset(utf8)"`
	Description string `json:"description,omitempty" orm:"column(description);charset(utf8)"`
	Content     string `json:"content,omitempty" orm:"column(content);charset(utf8)"`

	CreatorId       int64 `json:"creator_id,omitempty" orm:"column(creator_id)"`
	ParentLicenseId int64 `json:"parent_license_id,omitempty" orm:"column(parent_license_id)"`
	IsFork          bool  `json:"is_fork,omitempty" orm:"column(is_fork)"`
	ForkAble        bool  `json:"fork_able,omitempty" orm:"column(fork_able)"`

	Deleted bool `json:"deleted,omitempty" orm:"column(deleted)"`

	Created time.Time `json:"created"  orm:"auto_now_add;type(datetime)"`
	Updated time.Time `json:"updated" orm:"auto_now;type(datetime)"`
}

func NewLicense() *License {
	return &License{
		Deleted: false,
	}
}

func (l *License) Insert() (bool, error) {
	o := orm.NewOrm()
	err := o.Read(l, "lid")
	if err == orm.ErrNoRows {
		_, err = o.Insert(l)
		if err != nil {
			return false, errors.New("create License failed")
		}
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return false, errors.New("create License failed")
}

func (l *License) Get() (bool, error) {
	o := orm.NewOrm()
	err := o.Read(l, "lid")
	if err != nil {
		if err == orm.ErrNoRows {
			return false, errors.New("License not exist")
		}
		return false, err
	}
	return true, nil
}

func (l *License) Update(colms ...string) (bool, error) {
	if len(colms) <= 0 {
		return true, nil
	}
	o := orm.NewOrm()
	temp := NewLicense()
	temp.LID = l.LID
	err := o.Read(temp, "lid", "deleted")
	if err != nil {
		if err == orm.ErrNoRows {
			return false, errors.New("License not exist")
		}
		return false, err
	}
	_, err = o.Update(l, colms...)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (l *License) Delete() (bool, error) {
	o := orm.NewOrm()
	_, err := o.Update(l, "lid", "deleted")
	if err != nil {
		if err == orm.ErrNoRows {
			return false, errors.New("License not exist")
		}
		return false, err
	}
	return true, nil
}

func (l *License) GetGroupByCondition(condition string) (bool, error) {
	o := orm.NewOrm()
	err := o.Read(l, condition, "deleted")
	if err == orm.ErrNoRows {
		return false, errors.New("License not exist")
	}
	return true, nil
}
