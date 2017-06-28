package support

import "time"

type Career struct {
	CareerID  int64 `json:"careerid" orm:"column(careerid);pk;auto"`
	CompanyId int64 `json:"companyid,omitempty" orm:"column(companyid)"`
	JobId     int64 `json:"jobid,omitempty" orm:"column(companyid)"`

	Deleted bool      `json:"deleted,omitempty" orm:"column(deleted)"`
	Created time.Time `json:"created"  orm:"auto_now_add;type(datetime)"`
	Updated time.Time `json:"updated" orm:"auto_now;type(datetime)"`
}
