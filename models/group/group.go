package group

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

var o orm.Ormer

func Init() {
	o = orm.NewOrm()
	o.Using("default")
}

type Group struct {
	ID          int64
	GID         int64
	NickName    string
	Headline    string
	Description string

	Deleted bool `json:"deleted,omitempty"`

	Created time.Time `json:"created"  orm:"auto_now_add;type(datetime)"`
	Updated time.Time `json:"updated" orm:"auto_now;type(datetime)"`
}

func NewGroup() *Group {
	return &Group{
		Deleted: false,
	}
}

func (g *Group) Insert() (bool, error) {
	err := o.Read(g, "gid")
	if err == orm.ErrNoRows {
		_, err = o.Insert(g)
		if err != nil {
			return false, errors.New("create group failed")
		}
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return false, errors.New("create group failed")
}

func (g *Group) Get() (bool, error) {
	err := o.Read(g, "gid")
	if err != nil {
		if err == orm.ErrNoRows {
			return false, errors.New("group not exist")
		}
		return false, err
	}
	return true, nil
}

func (g *Group) Update(colms ...string) (bool, error) {
	if len(colms) <= 0 {
		return true, nil
	}
	temp := NewGroup()
	temp.GID = g.GID
	err := o.Read(temp, "gid", "deleted")
	if err != nil {
		if err == orm.ErrNoRows {
			return false, errors.New("group not exist")
		}
		return false, err
	}
	_, err = o.Update(g, colms...)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (g *Group) Delete() (bool, error) {
	_, err := o.Update(g, "gid", "deleted")
	if err != nil {
		if err == orm.ErrNoRows {
			return false, errors.New("group not exist")
		}
		return false, err
	}
	return true, nil
}

func (g *Group) GetGroupByCondition(condition string) (bool, error) {
	err := o.Read(g, condition, "deleted")
	if err == orm.ErrNoRows {
		return false, errors.New("group not exist")
	}
	return true, nil
}

func GetGroupByCondition(query string) ([]*Group, error) {
	return nil, nil
}

func DeleteGroupByCondition(query string) (bool, error) {
	return true, nil
}
