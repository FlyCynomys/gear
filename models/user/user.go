package user

import (
	"errors"
	"time"

	"github.com/FlyCynomys/gear/models/support"
	"github.com/astaxie/beego/orm"
)

var o = orm.NewOrm()

func Init() {
	o = orm.NewOrm()
	o.Using("default")
}

type User struct {
	ID          int64  `json:"id,omitempty"`
	UID         int64  `json:"uid,omitempty"`
	UrlToken    string `json:"url_token,omitempty"`
	NickName    string `json:"nick_name,omitempty"`
	UserType    string `json:"user_type,omitempty"`
	Gender      int    `json:"gender,omitempty"`
	RealName    string `json:"real_name,omitempty"`
	AvatarUrl   string `json:"avatar_url,omitempty"`
	IsOrg       bool   `json:"is_org,omitempty"`
	Description string `json:"description,omitempty"`
	CoverUrl    string `json:"cover_url,omitempty"`
	Email       string `json:"email,omitempty"`

	ShowSinaWeibo bool `json:"show_sina_weibo,omitempty"`
	IsBindSina    bool `json:"is_bind_sina,omitempty"`

	ThankFromCount    int `json:"thank_from_count,omitempty"`
	ThankToCount      int `json:"thank_to_count,omitempty"`
	QuestionCount     int `json:"question_count,omitempty"`
	FollowingCount    int `json:"following_count,omitempty"`
	VoteStarFromCount int `json:"vote_star_from_count,omitempty"`
	VoteStarToCount   int `json:"vote_star_to_count,omitempty"`

	Headline string `json:"headline,omitempty"`

	TodoplanCount    int `json:"todoplan_count,omitempty"`
	FailedPlanCount  int `json:"failed_plan_count,omitempty"`
	SuccessPlanCount int `json:"success_plan_count,omitempty"`

	Loc         []*support.Location `json:"loc,omitempty"`
	Employments []*support.Career   `json:"employments,omitempty"`

	Deleted bool      `json:"deleted,omitempty"`
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
