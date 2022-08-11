package onedev

import (
	"encoding/json"
	"fmt"
)

func (c Client) GetDefaultBranch(id int) (string, error) {
	body, err := c.get(fmt.Sprintf("%s/api/repositories/%d/default-branch", c.Url, id))
	if err != nil {
		return "", err
	}

	defaultbranch := ""
	err = json.NewDecoder(body).Decode(&defaultbranch)
	body.Close()

	return defaultbranch, err
}
