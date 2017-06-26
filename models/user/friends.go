package user

import "time"

type FriendShip struct {
	Id        int64 `json:"id,omitempty" orm:"column(id)"`
	Uid       int64 `json:"uid,omitempty" orm:"column(uid)"`
	FriendUid int64 `json:"frienduid,omitempty" orm:"column(frienduid)"`

	Deleted bool      `json:"deleted,omitempty" orm:"column(deleted)"`
	Created time.Time `json:"created"  orm:"auto_now_add;type(datetime)"`
	Updated time.Time `json:"updated" orm:"auto_now;type(datetime)"`
}
