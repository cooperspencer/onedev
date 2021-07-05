package onedev

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (c Client) GetProject(id int) (Project, error) {
	body, err := c.get(fmt.Sprintf("%s/api/projects/%d", c.Url, id))
	if err != nil {
		return Project{}, err
	}

	project := Project{}
	err = json.NewDecoder(body).Decode(&project)

	return project, err
}

func (c Client) CreateProject(options CreateProjectOptions) (int, error) {
	payloadbytes, err := json.Marshal(options)
	if err != nil {
		return 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, err := c.post(fmt.Sprintf("%s/api/projects", c.Url), payload)
	if err != nil {
		return 0, err
	}

	id := 0
	err = json.NewDecoder(body).Decode(&id)

	return id, err
}

func (c Client) GetProjects(options *ProjectQueryOptions) (Project, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/projects", c.Url), nil)

	q := req.URL.Query()
	if options.Query != "" {
		q.Add("query", options.Query)
	}
	q.Add("offset", fmt.Sprintf("%d", options.Offset))
	if options.Count == 0 {
		options.Count = 100
	}
	q.Add("count", fmt.Sprintf("%d", options.Count))

	body, err := c.get(q.Encode())
	if err != nil {
		return Project{}, err
	}

	project := Project{}
	err = json.NewDecoder(body).Decode(&project)

	return project, err
}

func (c Client) GetProjectForks(id int) ([]Project, error) {
	body, err := c.get(fmt.Sprintf("%s/api/projects/%d/forks", c.Url, id))
	if err != nil {
		return []Project{}, err
	}

	forks := []Project{}
	err = json.NewDecoder(body).Decode(&forks)

	return forks, err
}

func (c Client) GetProjectSettings(id int) (ProjectSettings, error) {
	body, err := c.get(fmt.Sprintf("%s/api/projects/%d/setting", c.Url, id))
	if err != nil {
		return ProjectSettings{}, err
	}

	projectsettings := ProjectSettings{}
	err = json.NewDecoder(body).Decode(&projectsettings)

	return projectsettings, err
}

func (c Client) GetGroupAuthorizations(id int) ([]GroupAuthorization, error) {
	body, err := c.get(fmt.Sprintf("%s/api/projects/%d/group-authorizations", c.Url, id))
	if err != nil {
		return []GroupAuthorization{}, err
	}

	groupauthorizations := []GroupAuthorization{}
	err = json.NewDecoder(body).Decode(&groupauthorizations)

	return groupauthorizations, err
}

func (c Client) GetUserAuthorizations(id int) ([]UserAuthorization, error) {
	body, err := c.get(fmt.Sprintf("%s/api/projects/%d/user-authorizations", c.Url, id))
	if err != nil {
		return []UserAuthorization{}, err
	}

	userauthorizations := []UserAuthorization{}
	err = json.NewDecoder(body).Decode(&userauthorizations)

	return userauthorizations, err
}

func (c Client) GetMilestones(id int, options *MilestoneQueryOptions) ([]Milestone, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/projects/%d/milestones", c.Url, id), nil)

	q := req.URL.Query()
	if options.Name != "" {
		q.Add("name", options.Name)
	}
	q.Add("closed", strconv.FormatBool(options.Closed))
	if !options.DueBefore.IsZero() {
		q.Add("dueBefore", options.DueBefore.Format("2006-01-02T15:04:05-0700"))
	}
	if !options.DueAfter.IsZero() {
		q.Add("dueAfter", options.DueAfter.Format("2006-01-02T15:04:05-0700"))
	}

	q.Add("offset", fmt.Sprintf("%d", options.Offset))
	if options.Count == 0 {
		options.Count = 100
	}
	q.Add("count", fmt.Sprintf("%d", options.Count))

	body, err := c.get(q.Encode())
	if err != nil {
		return []Milestone{}, err
	}

	milestones := []Milestone{}
	err = json.NewDecoder(body).Decode(&milestones)

	return milestones, err
}

func (c Client) GetTopContributors(id int, options *ContributorOptions) ([]Contribution, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/projects/%d/milestones", c.Url, id), nil)

	q := req.URL.Query()
	if options.Type != "" {
		q.Add("type", options.Type)
	}
	if !options.SinceDate.IsZero() {
		q.Add("sinceDate", options.SinceDate.Format("2006-01-02"))
	}
	if !options.UntilDate.IsZero() {
		q.Add("untilDate", options.UntilDate.Format("2006-01-02"))
	}

	if options.Count == 0 {
		options.Count = 1
	}
	q.Add("count", fmt.Sprintf("%d", options.Count))

	body, err := c.get(q.Encode())
	if err != nil {
		return []Contribution{}, err
	}

	contributions := []Contribution{}
	err = json.NewDecoder(body).Decode(&contributions)

	return contributions, err
}

func (c Client) UpdateProjectSettings(id int, options ProjectSettings) error {
	payloadbytes, err := json.Marshal(options)
	if err != nil {
		return err
	}
	payload := bytes.NewReader(payloadbytes)
	_, err = c.post(fmt.Sprintf("%s/api/projects/%d/setting", c.Url, id), payload)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) DeleteProject(id int) error {
	_, err := c.delete(fmt.Sprintf("%s/api/projects/%d/", c.Url, id))
	if err != nil {
		return err
	}
	return nil
}
