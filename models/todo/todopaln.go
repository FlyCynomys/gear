package todo

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

type TodoPlan struct {
	TodoPlanID  int64  `json:"todoplanid,omitempty" orm:"column(todoplanid);pk;auto"`
	OwnerID     int64  `json:"owner_id,omitempty" orm:"column(owner_id)"`
	Headline    string `json:"headline,omitempty" orm:"column(headlines)"`
	Description string `json:"description,omitempty" orm:"column(description)"`
	PlanType    string `json:"plan_type,omitempty" orm:"column(plan_type)"`

	Deleted bool      `json:"deleted,omitempty" orm:"column(deleted)"`
	Created time.Time `json:"created"  orm:"auto_now_add;type(datetime)"`
	Updated time.Time `json:"updated" orm:"auto_now;type(datetime)"`
}

func NewTodoPlan() *TodoPlan {
	return &TodoPlan{
		Deleted: false,
	}
}

func (c *TodoPlan) Insert() (int64, bool, error) {
	o := orm.NewOrm()
	index, err := o.Insert(c)
	if err != nil {
		return -1, false, errors.New("create TodoPlan failed")
	}
	return index, true, nil
}

func (c *TodoPlan) Get() (bool, error) {
	o := orm.NewOrm()
	err := o.Read(c, "todoplanid")
	if err != nil {
		if err == orm.ErrNoRows {
			return false, errors.New("TodoPlan not exist")
		}
		return false, err
	}
	return true, nil
}

func (c *TodoPlan) Update(colms ...string) (bool, error) {
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

func (c *TodoPlan) Delete() (bool, error) {
	o := orm.NewOrm()
	c.Deleted = true
	_, err := o.Update(c, "todoplanid", "deleted")
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetTodoPlanByCondition() ([]*TodoPlan, error) {
	return nil, nil
}

func DeleteTodoPlanByCondition(query string) (bool, error) {
	return true, nil
}
