package onedev

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c Client) GetIssueComment(id int) (Comment, error) {
	body, err := c.get(fmt.Sprintf("%s/api/issue-comments/%d", c.Url, id))
	if err != nil {
		return Comment{}, err
	}

	comment := Comment{}
	err = json.NewDecoder(body).Decode(&comment)
	body.Close()

	return comment, err
}

func (c Client) PostIssueComment(options Comment) (int, error) {
	payloadbytes, err := json.Marshal(options)
	if err != nil {
		return 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, err := c.post(fmt.Sprintf("%s/api/issue-comments", c.Url), payload)
	if err != nil {
		return 0, err
	}

	id := 0
	err = json.NewDecoder(body).Decode(&id)
	body.Close()

	return id, err
}

func (c Client) DeleteIssueComment(id int) error {
	_, err := c.delete(fmt.Sprintf("%s/api/issue-comments/%d/", c.Url, id))
	if err != nil {
		return err
	}
	return nil
}
