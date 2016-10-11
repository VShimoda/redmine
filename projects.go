package redmine

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type RequestProjectInclude struct {
	Trackers        bool
	IssueCategories bool
	EnabledModules  bool
}

type Project struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	Identifier      string          `json:"identifier"`
	Status          int             `json:"status"`
	Trackers        []Tracker       `json:"trackers"`
	IssueCategories []IssueCategory `json:"issue_categories"`
	EnabledModules  []EnabledModule `json:"enabled_modules"`
	CreatedOn       string          `json:"created_on"`
	UpdatedOn       string          `json:"updated_on"`
	Errors          []Error         `json:"errors"`
}

type Tracker struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type IssueCategory struct {
	Id   int    `json:"id"`
	Name string `json:"test"`
}

type EnabledModule struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (r Redmine) ProjectsList(args *RequestProjectInclude) error {
	client := &http.Client{}
	endpoint := r.BaseURL + "/projects.json"
	req, _ := http.NewRequest("GET", endpoint, nil)

	q := req.URL.Query()

	q.Add("key", r.Key)
	var includes []string
	if args.Trackers {
		includes = append(includes, "trackers")
	}
	if args.IssueCategories {
		includes = append(includes, "issue_categories")
	}
	if args.EnabledModules {
		includes = append(includes, "enabled_modules")
	}
	if len(includes) != 0 {
		q.Add("includes", strings.Join(includes, ","))
	}
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Status)
	fmt.Println(string(body))
	return nil
}
