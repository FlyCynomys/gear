package user

import (
	"errors"
	"time"

	"github.com/FlyCynomys/gear/models/support"
	"github.com/astaxie/beego/orm"
)

type User struct {
	ID          int64  `json:"id,omitempty" orm:"column(id);pk;auto"`
	UID         int64  `json:"uid,omitempty" orm:"column(uid)"`
	UrlToken    string `json:"url_token,omitempty" orm:"column(url_token)"`
	NickName    string `json:"nick_name,omitempty" orm:"column(nick_name)"`
	UserType    string `json:"user_type,omitempty" orm:"column(user_type)"`
	Gender      int    `json:"gender,omitempty" orm:"column(gender)"`
	RealName    string `json:"real_name,omitempty" orm:"column(real_name)"`
	AvatarUrl   string `json:"avatar_url,omitempty" orm:"column(avatar_url)"`
	IsOrg       bool   `json:"is_org,omitempty" orm:"column(is_org)"`
	Description string `json:"description,omitempty" orm:"column(description)"`
	CoverUrl    string `json:"cover_url,omitempty" orm:"column(cover_url)"`
	Email       string `json:"email,omitempty" orm:"column(email)"`

	ShowSinaWeibo bool `json:"show_sina_weibo,omitempty" orm:"column(show_sina_weibo)"`
	IsBindSina    bool `json:"is_bind_sina,omitempty" orm:"column(is_bind_sina)"`

	ThankFromCount    int `json:"thank_from_count,omitempty" orm:"column(thank_from_count)"`
	ThankToCount      int `json:"thank_to_count,omitempty" orm:"column(thank_to_count)"`
	QuestionCount     int `json:"question_count,omitempty" orm:"column(question_count)"`
	FollowingCount    int `json:"following_count,omitempty" orm:"column(following_count)"`
	VoteStarFromCount int `json:"vote_star_from_count,omitempty" orm:"column(vote_star_from_count)"`
	VoteStarToCount   int `json:"vote_star_to_count,omitempty" orm:"column(vote_star_to_count)"`

	Headline string `json:"headline,omitempty" orm:"column(headline)"`

	TodoplanCount    int `json:"todoplan_count,omitempty" orm:"column(todoplan_count)"`
	FailedPlanCount  int `json:"failed_plan_count,omitempty" orm:"column(failed_plan_count)"`
	SuccessPlanCount int `json:"success_plan_count,omitempty" orm:"column(success_plan_count)"`

	Loc         []*support.Location `json:"loc,omitempty" orm:"-"`
	Employments []*support.Career   `json:"employments,omitempty"  orm:"-"`

	Deleted bool      `json:"deleted,omitempty" orm:"column(deleted)"`
	Created time.Time `json:"created"  orm:"auto_now_add;type(datetime)"`
	Updated time.Time `json:"updated" orm:"auto_now;type(datetime)"`
}

func NewUser() *User {
	return &User{
		Deleted: false,
	}
}

func (u *User) Insert() (bool, error) {
	if u == nil {
		return false, errors.New("empty object")
	}
	o := orm.NewOrm()
	u.Deleted = false
	err := o.Read(u, "uid", "deleted")
	if err == orm.ErrNoRows {
		_, err := o.Insert(u)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return true, nil
}

func (u *User) Get() (bool, error) {
	o := orm.NewOrm()
	err := o.Read(u, "uid", "deleted")
	if err == orm.ErrNoRows {
		return false, errors.New("user not exist")
	}
	return true, nil
}

func (u *User) Update(colms ...string) (bool, error) {
	if len(colms) <= 0 {
		return true, nil
	}
	o := orm.NewOrm()
	var temp = NewUser()
	temp.UID = u.UID
	err := o.Read(temp, "uid", "deleted")
	if err == orm.ErrNoRows {
		return false, errors.New("user not exist")
	}
	_, err = o.Update(u, colms...)
	if err != nil {
		return false, errors.New("update failed")
	}
	return true, nil
}

func (u *User) Delete() (bool, error) {
	o := orm.NewOrm()
	err := o.Read(u, "uid", "deleted")
	if err == orm.ErrNoRows {
		return true, nil
	}
	_, err = o.Update(u, "deleted")
	if err != nil {
		return false, errors.New("delete user failed")
	}
	return true, nil
}

func (u *User) GetByCondition(condition string) (bool, error) {
	o := orm.NewOrm()
	err := o.Read(u, "condition", "deleted")
	if err == orm.ErrNoRows {
		return false, errors.New("user not exist")
	}
	return true, nil
}

func GetUserByLocation(query string) ([]*User, error) {
	return nil, nil
}

func GetUserByJob(query string) ([]*User, error) {
	return nil, nil
}

func GetUserByCompany(query string) ([]*User, error) {
	return nil, nil
}

func GetUserByNmae(query string) ([]*User, error) {
	return nil, nil
}

func DeleteUserByCondition(query string) (bool, error) {
	return true, nil
}
