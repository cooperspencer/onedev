package onedev

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c Client) GetIssueWatch(id int) (Watch, error) {
	body, err := c.get(fmt.Sprintf("%s/~api/issue-watches/%d", c.Url, id))
	if err != nil {
		return Watch{}, err
	}

	watch := Watch{}
	err = json.NewDecoder(body).Decode(&watch)
	body.Close()

	return watch, err
}

func (c Client) PostIssueWatch(options Watch) (int, error) {
	payloadbytes, err := json.Marshal(options)
	if err != nil {
		return 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, err := c.post(fmt.Sprintf("%s/~api/issue-watches", c.Url), payload)
	if err != nil {
		return 0, err
	}

	id := 0
	err = json.NewDecoder(body).Decode(&id)
	body.Close()

	return id, err
}

func (c Client) DeleteIssueWatch(id int) error {
	_, err := c.delete(fmt.Sprintf("%s/~api/issue-watches/%d/", c.Url, id))
	if err != nil {
		return err
	}
	return nil
}
