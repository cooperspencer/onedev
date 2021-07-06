package onedev

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c Client) GetIssue(id int) (Issue, error) {
	body, err := c.get(fmt.Sprintf("%s/api/issues/%d", c.Url, id))
	if err != nil {
		return Issue{}, err
	}

	issue := Issue{}
	err = json.NewDecoder(body).Decode(&issue)
	body.Close()

	return issue, err
}

func (c Client) GetIssueFields(id int) ([]IssueField, error) {
	body, err := c.get(fmt.Sprintf("%s/api/issues/%d/fields", c.Url, id))
	if err != nil {
		return []IssueField{}, err
	}

	issuefields := []IssueField{}
	err = json.NewDecoder(body).Decode(&issuefields)
	body.Close()

	return issuefields, err
}

func (c Client) GetIssueChanges(id int) ([]IssueChange, error) {
	body, err := c.get(fmt.Sprintf("%s/api/issues/%d/changes", c.Url, id))
	if err != nil {
		return []IssueChange{}, err
	}

	issuechanges := []IssueChange{}
	err = json.NewDecoder(body).Decode(&issuechanges)
	body.Close()

	return issuechanges, err
}

func (c Client) GetIssueComments(id int) ([]Comment, error) {
	body, err := c.get(fmt.Sprintf("%s/api/issues/%d/comments", c.Url, id))
	if err != nil {
		return []Comment{}, err
	}

	comments := []Comment{}
	err = json.NewDecoder(body).Decode(&comments)
	body.Close()

	return comments, err
}

func (c Client) GetIssueVotes(id int) ([]Vote, error) {
	body, err := c.get(fmt.Sprintf("%s/api/issues/%d/votes", c.Url, id))
	if err != nil {
		return []Vote{}, err
	}

	votes := []Vote{}
	err = json.NewDecoder(body).Decode(&votes)
	body.Close()

	return votes, err
}

func (c Client) GetIssueWatches(id int) ([]Watch, error) {
	body, err := c.get(fmt.Sprintf("%s/api/issues/%d/watches", c.Url, id))
	if err != nil {
		return []Watch{}, err
	}

	watches := []Watch{}
	err = json.NewDecoder(body).Decode(&watches)
	body.Close()

	return watches, err
}

func (c Client) GetIssuePullRequests(id int) ([]PullRequest, error) {
	body, err := c.get(fmt.Sprintf("%s/api/issues/%d/pull-requests", c.Url, id))
	if err != nil {
		return []PullRequest{}, err
	}

	pullrequests := []PullRequest{}
	err = json.NewDecoder(body).Decode(&pullrequests)
	body.Close()

	return pullrequests, err
}

func (c Client) GetIssueCommits(id int) ([]string, error) {
	body, err := c.get(fmt.Sprintf("%s/api/issues/%d/commits", c.Url, id))
	if err != nil {
		return []string{}, err
	}

	commits := []string{}
	err = json.NewDecoder(body).Decode(&commits)
	body.Close()

	return commits, err
}

func (c Client) GetIssues(options *IssueQueryOptions) ([]Issue, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/issues", c.Url), nil)

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

	body, err := c.get(req.URL.String())
	if err != nil {
		return []Issue{}, err
	}

	issues := []Issue{}
	err = json.NewDecoder(body).Decode(&issues)
	body.Close()

	return issues, err
}

func (c Client) CreateIssue(options CreateIssueOptions) (int, error) {
	payloadbytes, err := json.Marshal(options)
	if err != nil {
		return 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, err := c.post(fmt.Sprintf("%s/api/issues", c.Url), payload)
	if err != nil {
		return 0, err
	}

	id := 0
	err = json.NewDecoder(body).Decode(&id)
	body.Close()

	return id, err
}

func (c Client) SetIssueTitle(id int, title string) error {
	payloadbytes, err := json.Marshal(title)
	if err != nil {
		return err
	}
	payload := bytes.NewReader(payloadbytes)
	body, err := c.post(fmt.Sprintf("%s/api/issues/%d/title", c.Url, id), payload)
	if err != nil {
		return err
	}

	body.Close()

	return nil
}

func (c Client) SetIssueDescription(id int, title string) error {
	payloadbytes, err := json.Marshal(title)
	if err != nil {
		return err
	}
	payload := bytes.NewReader(payloadbytes)
	body, err := c.post(fmt.Sprintf("%s/api/issues/%d/description", c.Url, id), payload)
	if err != nil {
		return err
	}

	body.Close()

	return nil
}

func (c Client) SetIssueMilestone(id, milestoneId int) error {
	payloadbytes, err := json.Marshal(milestoneId)
	if err != nil {
		return err
	}
	payload := bytes.NewReader(payloadbytes)
	body, err := c.post(fmt.Sprintf("%s/api/issues/%d/milestone", c.Url, id), payload)
	if err != nil {
		return err
	}

	body.Close()

	return nil
}

func (c Client) SetIssueFields(id int, fields map[string]string) error {
	payloadbytes, err := json.Marshal(fields)
	if err != nil {
		return err
	}
	payload := bytes.NewReader(payloadbytes)
	body, err := c.post(fmt.Sprintf("%s/api/issues/%d/fields", c.Url, id), payload)
	if err != nil {
		return err
	}

	body.Close()

	return nil
}

func (c Client) SetStateTransition(id int, options StateTransitionOptions) error {
	payloadbytes, err := json.Marshal(options)
	if err != nil {
		return err
	}
	payload := bytes.NewReader(payloadbytes)
	body, err := c.post(fmt.Sprintf("%s/api/issues/%d/state-transitions", c.Url, id), payload)
	if err != nil {
		return err
	}

	body.Close()

	return nil
}

func (c Client) DeleteIssue(id int) error {
	_, err := c.delete(fmt.Sprintf("%s/api/issues/%d/", c.Url, id))
	if err != nil {
		return err
	}
	return nil
}
