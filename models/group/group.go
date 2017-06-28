package group

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

type Group struct {
	GroupID       int64  `json:"groupid,omitempty" orm:"column(groupid);pk;auto"`
	GroupNickName string `json:"group_nick_name,omitempty" orm:"column(group_nick_name);charset(utf8)"`
	Headline      string `json:"headline,omitempty" orm:"column(headline)"`
	Description   string `json:"description,omitempty" orm:"column(description)"`

	Deleted bool `json:"deleted,omitempty" orm:"column(deleted)"`

	Created time.Time `json:"created"  orm:"auto_now_add;type(datetime)"`
	Updated time.Time `json:"updated" orm:"auto_now;type(datetime)"`
}

func NewGroup() *Group {
	return &Group{
		Deleted: false,
	}
}

func (g *Group) Insert() (int64, bool, error) {
	o := orm.NewOrm()
	index, err := o.Insert(g)
	if err != nil {
		return -1, false, errors.New("insert group failed")
	}
	return index, true, nil
}

func (g *Group) Get() (bool, error) {
	o := orm.NewOrm()
	err := o.Read(g, "groupid")
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
	o := orm.NewOrm()
	_, err := o.Update(g, colms...)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (g *Group) Delete() (bool, error) {
	o := orm.NewOrm()
	g.Deleted = true
	_, err := o.Update(g, "groupid", "deleted")
	if err != nil {
		return false, err
	}
	return true, nil
}

func (g *Group) GetGroupByCondition(condition string) (bool, error) {
	o := orm.NewOrm()
	err := o.Read(g, condition, "deleted")
	if err == orm.ErrNoRows {
		return false, errors.New("group not exist")
	}
	return true, nil
}

func GetGroupByCondition(query string) ([]*Group, error) {
	grouplist := new([]*Group)
	o := orm.NewOrm()
	length, err := o.Raw(query).QueryRows(grouplist)
	if err != nil {
		return nil, err
	}
	if length <= 0 {
		return nil, nil
	}
	return *grouplist, nil
}

func DeleteGroupByCondition(query string) (int64, bool, error) {
	o := orm.NewOrm()
	pre, err := o.Raw(query).Prepare()
	if err != nil {
		return 0, false, err
	}
	defer pre.Close()
	result, err := pre.Exec()
	if err != nil {
		return 0, false, err
	}
	num, err := result.RowsAffected()
	if err != nil {
		return num, false, err
	}
	return 0, true, nil
}
