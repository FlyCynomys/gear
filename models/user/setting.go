package user

import (
	"time"

	"github.com/FlyCynomys/gear/models/support"
)

//user setting by user set
//用户的配置只有从属关系，没有删除功能，但是保留删除字段
type UserSetting struct {
	ID  int64 `json:"id" orm:"column(id)"`
	UID int64 `json:"uid" orm:"column(uid)"`

	IsOrg         bool   `json:"is_org,omitempty" orm:"column(is_org)"`
	Description   string `json:"description,omitempty" orm:"column(description)"`
	CoverUrl      string `json:"cover_url,omitempty" orm:"column(cover_url)"`
	Headline      string `json:"headline,omitempty" orm:"column(headline)"`
	ShowSinaWeibo bool   `json:"show_sina_weibo,omitempty" orm:"column(show_sina_weibo)"`
	IsBindSina    bool   `json:"is_bind_sina,omitempty" orm:"column(is_bind_sina)"`

	Loc         *support.Location `json:"location,omitempty" orm:"location"`
	Employments *support.Career   `json:"workin,omitempty"  orm:"workin"`

	Deleted bool      `json:"deleted,omitempty" orm:"column(deleted)"`
	Created time.Time `json:"created"  orm:"auto_now_add;type(datetime)"`
	Updated time.Time `json:"updated" orm:"auto_now;type(datetime)"`
}
