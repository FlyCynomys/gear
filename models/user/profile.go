package user

import (
	"time"

	"github.com/FlyCynomys/gear/models/support"
)

//usr profile modify by user active
//用户的配置只有从属关系，没有删除功能，但是保留删除字段

type Profile struct {
	ID  int64 `json:"id" orm:"column(id)"`
	UID int64 `json:"uid" orm:"column(uid)"`

	ThankFromCount    int `json:"thank_from_count,omitempty" orm:"column(thank_from_count)"`
	ThankToCount      int `json:"thank_to_count,omitempty" orm:"column(thank_to_count)"`
	QuestionCount     int `json:"question_count,omitempty" orm:"column(question_count)"`
	FollowingCount    int `json:"following_count,omitempty" orm:"column(following_count)"`
	VoteStarFromCount int `json:"vote_star_from_count,omitempty" orm:"column(vote_star_from_count)"`
	VoteStarToCount   int `json:"vote_star_to_count,omitempty" orm:"column(vote_star_to_count)"`

	TodoplanCount    int `json:"todoplan_count,omitempty" orm:"column(todoplan_count)"`
	FailedPlanCount  int `json:"failed_plan_count,omitempty" orm:"column(failed_plan_count)"`
	SuccessPlanCount int `json:"success_plan_count,omitempty" orm:"column(success_plan_count)"`

	Loc         []*support.Location `json:"loc,omitempty" orm:"-"`
	Employments []*support.Career   `json:"employments,omitempty"  orm:"-"`

	Deleted bool      `json:"deleted,omitempty" orm:"column(deleted)"`
	Created time.Time `json:"created"  orm:"auto_now_add;type(datetime)"`
	Updated time.Time `json:"updated" orm:"auto_now;type(datetime)"`
}
