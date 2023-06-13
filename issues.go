package onedev

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c Client) GetIssue(id int) (Issue, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/issues/%d", c.Url, id))
	if err != nil {
		return Issue{}, status, err
	}

	issue := Issue{}
	err = json.NewDecoder(body).Decode(&issue)
	body.Close()

	return issue, status, err
}

func (c Client) GetIssueFields(id int) ([]IssueField, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/issues/%d/fields", c.Url, id))
	if err != nil {
		return []IssueField{}, status, err
	}

	issuefields := []IssueField{}
	err = json.NewDecoder(body).Decode(&issuefields)
	body.Close()

	return issuefields, status, err
}

func (c Client) GetIssueChanges(id int) ([]IssueChange, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/issues/%d/changes", c.Url, id))
	if err != nil {
		return []IssueChange{}, status, err
	}

	issuechanges := []IssueChange{}
	err = json.NewDecoder(body).Decode(&issuechanges)
	body.Close()

	return issuechanges, status, err
}

func (c Client) GetIssueComments(id int) ([]Comment, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/issues/%d/comments", c.Url, id))
	if err != nil {
		return []Comment{}, status, err
	}

	comments := []Comment{}
	err = json.NewDecoder(body).Decode(&comments)
	body.Close()

	return comments, status, err
}

func (c Client) GetIssueVotes(id int) ([]Vote, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/issues/%d/votes", c.Url, id))
	if err != nil {
		return []Vote{}, status, err
	}

	votes := []Vote{}
	err = json.NewDecoder(body).Decode(&votes)
	body.Close()

	return votes, status, err
}

func (c Client) GetIssueWatches(id int) ([]Watch, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/issues/%d/watches", c.Url, id))
	if err != nil {
		return []Watch{}, status, err
	}

	watches := []Watch{}
	err = json.NewDecoder(body).Decode(&watches)
	body.Close()

	return watches, status, err
}

func (c Client) GetIssuePullRequests(id int) ([]PullRequest, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/issues/%d/pull-requests", c.Url, id))
	if err != nil {
		return []PullRequest{}, status, err
	}

	pullrequests := []PullRequest{}
	err = json.NewDecoder(body).Decode(&pullrequests)
	body.Close()

	return pullrequests, status, err
}

func (c Client) GetIssueCommits(id int) ([]string, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/issues/%d/commits", c.Url, id))
	if err != nil {
		return []string{}, status, err
	}

	commits := []string{}
	err = json.NewDecoder(body).Decode(&commits)
	body.Close()

	return commits, status, err
}

func (c Client) GetIssues(options *IssueQueryOptions) ([]Issue, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/~api/issues", c.Url), nil)

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
		return []Issue{}, status, err
	}

	issues := []Issue{}
	err = json.NewDecoder(body).Decode(&issues)
	body.Close()

	return issues, status, err
}

func (c Client) CreateIssue(options CreateIssueOptions) (int, int, error) {
	payloadbytes, err := json.Marshal(options)
	if err != nil {
		return 0, 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, status, err := c.post(fmt.Sprintf("%s/~api/issues", c.Url), payload)
	if err != nil {
		return 0, status, err
	}

	id := 0
	err = json.NewDecoder(body).Decode(&id)
	body.Close()

	return id, status, err
}

func (c Client) SetIssueTitle(id int, title string) (int, error) {
	payloadbytes, err := json.Marshal(title)
	if err != nil {
		return 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, status, err := c.post(fmt.Sprintf("%s/~api/issues/%d/title", c.Url, id), payload)
	if err != nil {
		return status, err
	}

	body.Close()

	return status, nil
}

func (c Client) SetIssueDescription(id int, title string) (int, error) {
	payloadbytes, err := json.Marshal(title)
	if err != nil {
		return 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, status, err := c.post(fmt.Sprintf("%s/~api/issues/%d/description", c.Url, id), payload)
	if err != nil {
		return status, err
	}

	body.Close()

	return status, nil
}

func (c Client) SetIssueMilestone(id, milestoneId int) (int, error) {
	payloadbytes, err := json.Marshal(milestoneId)
	if err != nil {
		return 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, status, err := c.post(fmt.Sprintf("%s/~api/issues/%d/milestone", c.Url, id), payload)
	if err != nil {
		return status, err
	}

	body.Close()

	return status, nil
}

func (c Client) SetIssueFields(id int, fields map[string]string) (int, error) {
	payloadbytes, err := json.Marshal(fields)
	if err != nil {
		return 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, status, err := c.post(fmt.Sprintf("%s/~api/issues/%d/fields", c.Url, id), payload)
	if err != nil {
		return status, err
	}

	body.Close()

	return status, nil
}

func (c Client) SetStateTransition(id int, options StateTransitionOptions) (int, error) {
	payloadbytes, err := json.Marshal(options)
	if err != nil {
		return 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, status, err := c.post(fmt.Sprintf("%s/~api/issues/%d/state-transitions", c.Url, id), payload)
	if err != nil {
		return status, err
	}

	body.Close()

	return status, nil
}

func (c Client) DeleteIssue(id int) (int, error) {
	_, status, err := c.delete(fmt.Sprintf("%s/~api/issues/%d/", c.Url, id))
	if err != nil {
		return status, err
	}
	return status, nil
}
