package license

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

type License struct {
	LicenseID   int64  `json:"licenseid,omitempty" orm:"column(licenseid);pk;auto"`
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

func (l *License) Insert() (int64, bool, error) {
	o := orm.NewOrm()
	index, err := o.Insert(l)
	if err != nil {
		return -1, false, errors.New("create License failed")
	}
	return index, true, nil
}

func (l *License) Get() (bool, error) {
	o := orm.NewOrm()
	err := o.Read(l, "licenseid")
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
	_, err := o.Update(l, colms...)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (l *License) Delete() (bool, error) {
	o := orm.NewOrm()
	l.Deleted = true
	_, err := o.Update(l, "licenseid", "deleted")
	if err != nil {
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
