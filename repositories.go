package onedev

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) GetDefaultBranch(id int) (string, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/repositories/%d/default-branch", c.Url, id))
	if err != nil {
		return "", status, err
	}

	defaultbranch, err := io.ReadAll(body)
	body.Close()

	return string(defaultbranch), status, err
}

func (c Client) GetCommits(id int, options *CommitQueryOptions) ([]string, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/~api/repositories/%d/commits", c.Url, id), nil)

	q := req.URL.Query()

	if options.Count == 0 {
		options.Count = 1
	}
	q.Add("count", fmt.Sprintf("%d", options.Count))
	if options.Query != "" {
		q.Add("query", options.Query)
	}
	req.URL.RawQuery = q.Encode()

	body, status, err := c.get(req.URL.String())
	if err != nil {
		return []string{}, status, err
	}

	commits := []string{}
	err = json.NewDecoder(body).Decode(&commits)
	body.Close()

	return commits, status, err
}

func (c Client) GetCommit(id int, commitHash string) (Commit, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/repositories/%d/commits/%s", c.Url, id, commitHash))
	if err != nil {
		return Commit{}, status, err
	}

	commit := Commit{}
	err = json.NewDecoder(body).Decode(&commit)
	body.Close()

	return commit, status, err
}
