package redmine

type Redmine struct {
	BaseURL string
	Key     string
}

type Error struct {
	Error string `json:"errors"`
}

func (r Redmine) GetIssues() error {
	//endpoint := r.BaseURL + "issues/"
	return nil
}
