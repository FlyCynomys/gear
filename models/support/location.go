package support

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

type Location struct {
	LocationID   int64  `json:"id,omitempty" orm:"column(locid);pk;auto"`
	Name         string `json:"name,omitempty" orm:"column(name)"`
	URL          string `json:"url,omitempty" orm:"column(url)"`
	AvatarUrl    string `json:"avatar_url,omitempty" orm:"column(avatar_url)"`
	Introduction string `json:"introduction,omitempty" orm:"column(introduction);type(text)"`
	Type         string `json:"type,omitempty" orm:"column(type)"`
	Excerpt      string `json:"excerpt,omitempty" orm:"column(excerpt);type(text)"`

	Deleted bool      `json:"deleted,omitempty" orm:"column(deleted)"`
	Created time.Time `json:"created"  orm:"auto_now_add;type(datetime)"`
	Updated time.Time `json:"updated" orm:"auto_now;type(datetime)"`
}

func NewLocation() *Location {
	return &Location{
		Deleted: false,
	}
}

func (c *Location) Insert() (int64, bool, error) {
	o := orm.NewOrm()
	index, err := o.Insert(c)
	if err != nil {
		return -1, false, errors.New("create Location failed")
	}
	return index, true, nil
}

func (c *Location) Get() (bool, error) {
	o := orm.NewOrm()
	err := o.Read(c, "locid")
	if err != nil {
		if err == orm.ErrNoRows {
			return false, errors.New("Location not exist")
		}
		return false, err
	}
	return true, nil
}

func (c *Location) Update(colms ...string) (bool, error) {
	if len(colms) <= 0 {
		return true, nil
	}
	o := orm.NewOrm()
	_, err := o.Update(c, colms...)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *Location) Delete() (bool, error) {
	o := orm.NewOrm()
	c.Deleted = true
	_, err := o.Update(c, "locid", "deleted")
	if err != nil {
		return false, err
	}
	return true, nil
}
