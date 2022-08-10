package onedev

import (
	"encoding/json"
	"fmt"
)

func (c Client) GetMe() (User, error) {
	body, err := c.get(fmt.Sprintf("%s/api/users/me", c.Url))
	if err != nil {
		return User{}, err
	}

	user := User{}
	err = json.NewDecoder(body).Decode(&user)
	body.Close()

	return user, err
}
