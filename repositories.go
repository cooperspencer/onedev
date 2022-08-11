package onedev

import (
	"fmt"
	"io"
)

func (c Client) GetDefaultBranch(id int) (string, error) {
	body, err := c.get(fmt.Sprintf("%s/api/repositories/%d/default-branch", c.Url, id))
	if err != nil {
		return "", err
	}

	defaultbranch, err := io.ReadAll(body)
	body.Close()

	return string(defaultbranch), err
}
