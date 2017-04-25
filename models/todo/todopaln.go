package todo

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

var o = orm.NewOrm()

func Init() {
	o = orm.NewOrm()
	o.Using("default")
}

type TodoPlan struct {
	ID          int64  `json:"id,omitempty"`
	TID         int64  `json:"tid,omitempty"`
	OwnerID     int64  `json:"owner_id,omitempty"`
	Headline    string `json:"headline,omitempty"`
	Description string `json:"description,omitempty"`
	PlanType    string `json:"plan_type,omitempty"`

	Deleted bool      `json:"deleted,omitempty"`
	Created time.Time `json:"created"  orm:"auto_now_add;type(datetime)"`
	Updated time.Time `json:"updated" orm:"auto_now;type(datetime)"`
}

func NewTodoPlan() *TodoPlan {
	return &TodoPlan{
		Deleted: false,
	}
}

func (t *TodoPlan) Insert() (bool, error) {
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
