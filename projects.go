package onedev

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func (c Client) GetProject(id int) (Project, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/projects/%d", c.Url, id))
	if err != nil {
		return Project{}, status, err
	}

	project := Project{}
	err = json.NewDecoder(body).Decode(&project)
	body.Close()

	return project, status, err
}

func (c Client) CreateProject(options CreateProjectOptions) (int, int, error) {
	payloadbytes, err := json.Marshal(options)
	if err != nil {
		return 0, 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, status, err := c.post(fmt.Sprintf("%s/~api/projects", c.Url), payload)
	if err != nil {
		return 0, status, err
	}
	defer body.Close()
	b, err := io.ReadAll(body)
	if err != nil {
		return 0, status, err
	}

	id := 0
	err = json.NewDecoder(bytes.NewReader(b)).Decode(&id)
	if err != nil {
		err = errors.New(string(b))
	}

	return id, status, err
}

func (c Client) GetProjects(options *ProjectQueryOptions) ([]Project, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/~api/projects", c.Url), nil)

	q := req.URL.Query()
	if options.Query != "" {
		q.Add("query", options.Query)
	}
	q.Add("offset", fmt.Sprintf("%d", options.Offset))
	if options.Count == 0 {
		options.Count = 100
	}
	q.Add("count", fmt.Sprintf("%d", options.Count))

	req.URL.RawQuery = q.Encode()

	body, status, err := c.get(req.URL.String())
	if err != nil {
		return []Project{}, status, err
	}

	projects := []Project{}
	err = json.NewDecoder(body).Decode(&projects)
	body.Close()

	return projects, status, err
}

func (c Client) GetProjectForks(id int) ([]Project, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/projects/%d/forks", c.Url, id))
	if err != nil {
		return []Project{}, 0, err
	}

	forks := []Project{}
	err = json.NewDecoder(body).Decode(&forks)
	body.Close()

	return forks, status, err
}

func (c Client) GetProjectSettings(id int) (ProjectSettings, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/projects/%d/setting", c.Url, id))
	if err != nil {
		return ProjectSettings{}, status, err
	}

	projectsettings := ProjectSettings{}
	err = json.NewDecoder(body).Decode(&projectsettings)
	body.Close()

	return projectsettings, status, err
}

func (c Client) GetGroupAuthorizations(id int) ([]GroupAuthorization, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/projects/%d/group-authorizations", c.Url, id))
	if err != nil {
		return []GroupAuthorization{}, status, err
	}

	groupauthorizations := []GroupAuthorization{}
	err = json.NewDecoder(body).Decode(&groupauthorizations)
	body.Close()

	return groupauthorizations, status, err
}

func (c Client) GetUserAuthorizations(id int) ([]UserAuthorization, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/projects/%d/user-authorizations", c.Url, id))
	if err != nil {
		return []UserAuthorization{}, status, err
	}

	userauthorizations := []UserAuthorization{}
	err = json.NewDecoder(body).Decode(&userauthorizations)
	body.Close()

	return userauthorizations, status, err
}

func (c Client) GetMilestones(id int, options *MilestoneQueryOptions) ([]Milestone, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/~api/projects/%d/milestones", c.Url, id), nil)

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

	req.URL.RawQuery = q.Encode()

	body, status, err := c.get(req.URL.String())
	if err != nil {
		return []Milestone{}, status, err
	}

	milestones := []Milestone{}
	err = json.NewDecoder(body).Decode(&milestones)
	body.Close()

	return milestones, status, err
}

func (c Client) GetTopContributors(id int, options *ContributorOptions) ([]Contribution, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/~api/projects/%d/milestones", c.Url, id), nil)

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

	req.URL.RawQuery = q.Encode()

	body, status, err := c.get(req.URL.String())
	if err != nil {
		return []Contribution{}, status, err
	}

	contributions := []Contribution{}
	err = json.NewDecoder(body).Decode(&contributions)
	body.Close()

	return contributions, status, err
}

func (c Client) UpdateProjectSettings(id int, options ProjectSettings) (int, error) {
	payloadbytes, err := json.Marshal(options)
	if err != nil {
		return 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	_, status, err := c.post(fmt.Sprintf("%s/~api/projects/%d/setting", c.Url, id), payload)
	if err != nil {
		return status, err
	}
	return status, nil
}

func (c Client) DeleteProject(id int) (int, error) {
	_, status, err := c.delete(fmt.Sprintf("%s/~api/projects/%d/", c.Url, id))
	if err != nil {
		return status, err
	}
	return status, nil
}

func (c Client) GetCloneUrl(id int) (CloneUrl, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/projects/%d/clone-url", c.Url, id))
	if err != nil {
		return CloneUrl{}, status, err
	}

	cloneurls := CloneUrl{}
	err = json.NewDecoder(body).Decode(&cloneurls)
	body.Close()

	return cloneurls, status, err
}
