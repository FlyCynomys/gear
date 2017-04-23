package todo

import "time"

type TodoPlan struct {
	Deleted bool      `json:"deleted,omitempty"`
	Created time.Time `json:"created"  orm:"auto_now_add;type(datetime)"`
	Updated time.Time `json:"updated" orm:"auto_now;type(datetime)"`
}
