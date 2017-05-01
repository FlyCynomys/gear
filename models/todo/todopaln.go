package todo

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

type TodoPlan struct {
	ID          int64  `json:"id,omitempty" orm:"column(id);pk;auto"`
	TID         int64  `json:"tid,omitempty" orm:"column(tid)"`
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

func (t *TodoPlan) Insert() (bool, error) {
	o := orm.NewOrm()
	err := o.Read(t, "tid", "deleted")
	if err == orm.ErrNoRows {
		_, err = o.Insert(t)
		if err != nil {
			return false, errors.New("create plan failed")
		}
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (t *TodoPlan) Get() (bool, error) {
	o := orm.NewOrm()
	err := o.Read(t, "tid", "deleted")
	if err == orm.ErrNoRows {
		return false, errors.New("plan not exist")
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (t *TodoPlan) Update(colms ...string) (bool, error) {

	if len(colms) <= 0 {
		return true, nil
	}
	o := orm.NewOrm()

	temp := NewTodoPlan()
	temp.TID = t.TID
	err := o.Read(temp, "tid")
	if err == orm.ErrNoRows {
		return false, errors.New("plan not exist")
	}
	_, err = o.Update(t, colms...)
	if err != nil {
		return false, errors.New("update plan failed")
	}
	return true, nil
}

func (t *TodoPlan) Delete() (bool, error) {
	o := orm.NewOrm()
	err := o.Read(t, "tid", "deleted")
	if err == orm.ErrNoRows {
		return true, nil
	}
	t.Deleted = true
	_, err = o.Update(t, "deleted")
	if err != nil {
		return false, errors.New("delete plan failed")
	}
	return true, nil
}

func (t *TodoPlan) GetTodoPlanByCondition(condition string) (bool, error) {
	o := orm.NewOrm()
	err := o.Read(t, condition, "deleted")
	if err == orm.ErrNoRows {
		return false, errors.New("plan not exist")
	}
	return true, nil
}

func GetTodoPlanByCondition() ([]*TodoPlan, error) {
	return nil, nil
}

func DeleteTodoPlanByCondition(query string) (bool, error) {
	return true, nil
}
