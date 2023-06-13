package onedev

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c Client) GetMilestone(id int) (Milestone, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/milestones/%d", c.Url, id))
	if err != nil {
		return Milestone{}, status, err
	}

	milestone := Milestone{}
	err = json.NewDecoder(body).Decode(&milestone)
	body.Close()

	return milestone, status, err
}

func (c Client) PostMilestone(options Milestone) (int, int, error) {
	payloadbytes, err := json.Marshal(options)
	if err != nil {
		return 0, 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, status, err := c.post(fmt.Sprintf("%s/~api/milestones", c.Url), payload)
	if err != nil {
		return 0, status, err
	}

	id := 0
	err = json.NewDecoder(body).Decode(&id)
	body.Close()

	return id, status, err
}

func (c Client) DeleteMilestone(id int) (int, error) {
	_, status, err := c.delete(fmt.Sprintf("%s/~api/milestones/%d/", c.Url, id))
	if err != nil {
		return status, err
	}
	return status, nil
}
