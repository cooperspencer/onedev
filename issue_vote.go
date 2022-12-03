package onedev

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c Client) GetIssueVote(id int) (Vote, error) {
	body, err := c.get(fmt.Sprintf("%s/~api/issue-votes/%d", c.Url, id))
	if err != nil {
		return Vote{}, err
	}

	vote := Vote{}
	err = json.NewDecoder(body).Decode(&vote)
	body.Close()

	return vote, err
}

func (c Client) PostIssueVote(options Vote) (int, error) {
	payloadbytes, err := json.Marshal(options)
	if err != nil {
		return 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, err := c.post(fmt.Sprintf("%s/~api/issue-votes", c.Url), payload)
	if err != nil {
		return 0, err
	}

	id := 0
	err = json.NewDecoder(body).Decode(&id)
	body.Close()

	return id, err
}

func (c Client) DeleteIssueVote(id int) error {
	_, err := c.delete(fmt.Sprintf("%s/~api/issue-votes/%d/", c.Url, id))
	if err != nil {
		return err
	}
	return nil
}
