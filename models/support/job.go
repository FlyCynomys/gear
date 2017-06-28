package support

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

type Job struct {
	JobID        int64  `json:"id,omitempty" orm:"column(jobid);pk;auto"`
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

func NewJob() *Job {
	return &Job{
		Deleted: false,
	}
}

func (j *Job) Insert() (int64, bool, error) {
	o := orm.NewOrm()
	index, err := o.Insert(j)
	if err != nil {
		return -1, false, errors.New("create Job failed")
	}
	return index, true, nil
}

func (j *Job) Get() (bool, error) {
	o := orm.NewOrm()
	err := o.Read(j, "Jobid")
	if err != nil {
		if err == orm.ErrNoRows {
			return false, errors.New("Job not exist")
		}
		return false, err
	}
	return true, nil
}

func (j *Job) Update(colms ...string) (bool, error) {
	if len(colms) <= 0 {
		return true, nil
	}
	o := orm.NewOrm()
	_, err := o.Update(j, colms...)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (j *Job) Delete() (bool, error) {
	o := orm.NewOrm()
	j.Deleted = true
	_, err := o.Update(j, "Jobid", "deleted")
	if err != nil {
		return false, err
	}
	return true, nil
}
