package onedev

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c Client) GetIssueWatch(id int) (Watch, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/issue-watches/%d", c.Url, id))
	if err != nil {
		return Watch{}, status, err
	}

	watch := Watch{}
	err = json.NewDecoder(body).Decode(&watch)
	body.Close()

	return watch, status, err
}

func (c Client) PostIssueWatch(options Watch) (int, int, error) {
	payloadbytes, err := json.Marshal(options)
	if err != nil {
		return 0, 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, status, err := c.post(fmt.Sprintf("%s/~api/issue-watches", c.Url), payload)
	if err != nil {
		return 0, status, err
	}

	id := 0
	err = json.NewDecoder(body).Decode(&id)
	body.Close()

	return id, status, err
}

func (c Client) DeleteIssueWatch(id int) (int, error) {
	_, status, err := c.delete(fmt.Sprintf("%s/~api/issue-watches/%d/", c.Url, id))
	if err != nil {
		return status, err
	}
	return status, nil
}
