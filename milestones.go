package onedev

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c Client) GetMilestone(id int) (Milestone, error) {
	body, err := c.get(fmt.Sprintf("%s/~api/milestones/%d", c.Url, id))
	if err != nil {
		return Milestone{}, err
	}

	milestone := Milestone{}
	err = json.NewDecoder(body).Decode(&milestone)
	body.Close()

	return milestone, err
}

func (c Client) PostMilestone(options Milestone) (int, error) {
	payloadbytes, err := json.Marshal(options)
	if err != nil {
		return 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, err := c.post(fmt.Sprintf("%s/~api/milestones", c.Url), payload)
	if err != nil {
		return 0, err
	}

	id := 0
	err = json.NewDecoder(body).Decode(&id)
	body.Close()

	return id, err
}

func (c Client) DeleteMilestone(id int) error {
	_, err := c.delete(fmt.Sprintf("%s/~api/milestones/%d/", c.Url, id))
	if err != nil {
		return err
	}
	return nil
}
