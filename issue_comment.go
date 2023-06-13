package onedev

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c Client) GetIssueComment(id int) (Comment, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/issue-comments/%d", c.Url, id))
	if err != nil {
		return Comment{}, status, err
	}

	comment := Comment{}
	err = json.NewDecoder(body).Decode(&comment)
	body.Close()

	return comment, status, err
}

func (c Client) PostIssueComment(options Comment) (int, int, error) {
	payloadbytes, err := json.Marshal(options)
	if err != nil {
		return 0, 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, status, err := c.post(fmt.Sprintf("%s/~api/issue-comments", c.Url), payload)
	if err != nil {
		return 0, status, err
	}

	id := 0
	err = json.NewDecoder(body).Decode(&id)
	body.Close()

	return id, status, err
}

func (c Client) DeleteIssueComment(id int) (int, error) {
	_, status, err := c.delete(fmt.Sprintf("%s/~api/issue-comments/%d/", c.Url, id))
	if err != nil {
		return status, err
	}
	return status, nil
}
