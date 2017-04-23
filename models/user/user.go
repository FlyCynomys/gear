package user

import (
	"errors"

	"github.com/FlyCynomys/gear/models/support"
	"github.com/astaxie/beego/orm"
)

var o = orm.NewOrm()

type User struct {
	ID          int64  `json:"id,omitempty"`
	UID         string `json:"uid,omitempty"`
	UrlToken    string `json:"url_token,omitempty"`
	NickName    string `json:"nick_name,omitempty"`
	UserType    string `json:"user_type,omitempty"`
	Gender      int    `json:"gender,omitempty"`
	RealName    string `json:"real_name,omitempty"`
	AvatarUrl   string `json:"avatar_url,omitempty"`
	IsOrg       bool   `json:"is_org,omitempty"`
	Description string `json:"description,omitempty"`
	CoverUrl    string `json:"cover_url,omitempty"`

	ShowSinaWeibo bool `json:"show_sina_weibo,omitempty"`
	IsBindSina    bool `json:"is_bind_sina,omitempty"`

	ThankFromCount    int `json:"thank_from_count,omitempty"`
	ThankToCount      int `json:"thank_to_count,omitempty"`
	QuestionCount     int `json:"question_count,omitempty"`
	FollowingCount    int `json:"following_count,omitempty"`
	VoteStarFromCount int `json:"vote_star_from_count,omitempty"`
	VoteStarToCount   int `json:"vote_star_to_count,omitempty"`

	Headline string `json:"headline,omitempty"`

	ToDoPlanCount        int `json:"to_do_plan_count,omitempty"`
	FailedToDoPlanCount  int `json:"failed_to_do_plan_count,omitempty"`
	SuccessToDoPlanCount int `json:"success_to_do_plan_count,omitempty"`

	Loc         []*support.Location `json:"loc,omitempty"`
	Employments []*support.Career   `json:"employments,omitempty"`
}

func (u *User) Insert() (bool, error) {
	if u == nil {
		return false, errors.New("empty object")
	}
	err := o.Read(u, "uid")
	if err == orm.ErrNoRows {
		_, err := o.Insert(u)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return true, nil
}
