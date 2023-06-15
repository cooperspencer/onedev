package onedev

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c Client) GetIssueVote(id int) (Vote, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/issue-votes/%d", c.Url, id))
	if err != nil {
		return Vote{}, status, err
	}

	vote := Vote{}
	err = json.NewDecoder(body).Decode(&vote)
	body.Close()

	return vote, status, err
}

func (c Client) PostIssueVote(options *Vote) (int, int, error) {
	payloadbytes, err := json.Marshal(options)
	if err != nil {
		return 0, 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, status, err := c.post(fmt.Sprintf("%s/~api/issue-votes", c.Url), payload)
	if err != nil {
		return 0, status, err
	}

	id := 0
	err = json.NewDecoder(body).Decode(&id)
	body.Close()

	return id, status, err
}

func (c Client) DeleteIssueVote(id int) (int, error) {
	_, status, err := c.delete(fmt.Sprintf("%s/~api/issue-votes/%d/", c.Url, id))
	if err != nil {
		return status, err
	}
	return status, nil
}
