package onedev

import (
	"encoding/json"
	"fmt"
)

func (c Client) GetGroup(id int) (Group, error) {
	body, err := c.get(fmt.Sprintf("%s/~api/groups/%d", c.Url, id))
	if err != nil {
		return Group{}, err
	}

	group := Group{}
	err = json.NewDecoder(body).Decode(&group)
	body.Close()

	return group, err
}
