package active

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

type Active struct {
	ID          int64  `json:"id,omitempty" orm:"column(id);pk;auto"`
	ActiveID    int64  `json:"activeid,omitempty" orm:"column(activeid)"`
	Topic       string `json:"topic,omitempty" orm:"column(topic)"`
	Headline    string `json:"headline,omitempty" orm:"column(headline)"`
	Description string `json:"description,omitempty" orm:"column(descriptions)"`

	Deleted bool      `json:"deleted,omitempty" orm:"column(deleted)"`
	Created time.Time `json:"created"  orm:"auto_now_add;type(datetime)"`
	Updated time.Time `json:"updated" orm:"auto_now;type(datetime)"`
}

func NewActive() *Active {
	return &Active{
		Deleted: false,
	}
}

func (a *Active) Insert() (bool, error) {
	o := orm.NewOrm()
	o.Using("default")
	err := o.Read(a, "activeid")
	if err == orm.ErrNoRows {
		_, err = o.Insert(a)
		if err != nil {
			return false, errors.New("create active failed")
		}
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return false, errors.New("create active failed")
}

func (a *Active) Get() (bool, error) {
	o := orm.NewOrm()
	o.Using("default")
	err := o.Read(a, "activeid", "deleted")
	if err != nil {
		if err == orm.ErrNoRows {
			return false, errors.New("active not exist")
		}
		return false, err
	}
	return true, nil
}

func (a *Active) Update(colms ...string) (bool, error) {
	if len(colms) <= 0 {
		return true, nil
	}
	temp := NewActive()
	temp.ActiveID = a.ActiveID
	o := orm.NewOrm()
	o.Using("default")
	err := o.Read(temp, "activeid", "deleted")
	if err != nil {
		if err == orm.ErrNoRows {
			return false, errors.New("active not exist")
		}
		return false, err
	}
	_, err = o.Update(a, colms...)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *Active) Delete() (bool, error) {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Update(a, "activeid", "deleted")
	if err != nil {
		if err == orm.ErrNoRows {
			return false, errors.New("active not exist")
		}
		return false, err
	}
	return true, nil
}

func (a *Active) GetActiveByCondition(condition string) (bool, error) {
	o := orm.NewOrm()
	o.Using("default")
	err := o.Read(a, condition, "deleted")
	if err == orm.ErrNoRows {
		return false, errors.New("active not exist")
	}
	return true, nil
}

func GetActiveByCondition(query string) ([]*Active, error) {
	return nil, nil
}

func DeleteActiveByCondition(query string) (bool, error) {
	return true, nil
}
