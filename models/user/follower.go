package user

import "time"

type FollowShip struct {
	Id        int64 `json:"id,omitempty" orm:"column(id)"`
	UID       int64 `json:"uid,omitempty" orm:"column(uid)"`
	FolloweId int64 `json:"followe_id,omitempty" orm:"column(followe_id)"`

	Deleted bool      `json:"deleted,omitempty" orm:"column(deleted)"`
	Created time.Time `json:"created"  orm:"auto_now_add;type(datetime)"`
	Updated time.Time `json:"updated" orm:"auto_now;type(datetime)"`
}
