package support

type Location struct {
	ID           int64  `json:"-" orm:"column(id);pk;auto"`
	LID          string `json:"id,omitempty" orm:"column(locid)"`
	Name         string `json:"name,omitempty" orm:"column(name)"`
	URL          string `json:"url,omitempty" orm:"column(url)"`
	AvatarUrl    string `json:"avatar_url,omitempty" orm:"column(avatar_url)"`
	Introduction string `json:"introduction,omitempty" orm:"column(introduction);type(text)"`
	Type         string `json:"type,omitempty" orm:"column(type)"`
	Excerpt      string `json:"excerpt,omitempty" orm:"column(excerpt);type(text)"`
}

type Career struct {
	Company *Company `json:"company,omitempty"`
	Job     *Job     `json:"job,omitempty"`
}

type Company struct {
	ID           int64  `json:"-" orm:"column(id);pk;auto"`
	CompanyID    string `json:"id,omitempty" orm:"column(companyid)"`
	Name         string `json:"name,omitempty" orm:"column(name)"`
	URL          string `json:"url,omitempty" orm:"column(url)"`
	AvatarUrl    string `json:"avatar_url,omitempty" orm:"column(avatar_url)"`
	Introduction string `json:"introduction,omitempty" orm:"column(introduction);type(text)"`
	Type         string `json:"type,omitempty" orm:"column(type)"`
	Excerpt      string `json:"excerpt,omitempty" orm:"column(excerpt);type(text)"`
}

type Job struct {
	ID           int64  `json:"-" orm:"column(id);pk;auto"`
	JobID        string `json:"id,omitempty" orm:"column(jobid)"`
	Name         string `json:"name,omitempty" orm:"column(name)"`
	URL          string `json:"url,omitempty" orm:"column(url)"`
	AvatarUrl    string `json:"avatar_url,omitempty" orm:"column(avatar_url)"`
	Introduction string `json:"introduction,omitempty" orm:"column(introduction);type(text)"`
	Type         string `json:"type,omitempty" orm:"column(type)"`
	Excerpt      string `json:"excerpt,omitempty" orm:"column(excerpt);type(text)"`
}
