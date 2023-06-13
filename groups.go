package onedev

import (
	"encoding/json"
	"fmt"
)

func (c Client) GetGroup(id int) (Group, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/groups/%d", c.Url, id))
	if err != nil {
		return Group{}, status, err
	}

	group := Group{}
	err = json.NewDecoder(body).Decode(&group)
	body.Close()

	return group, status, err
}
