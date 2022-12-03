package onedev

import (
	"encoding/json"
	"fmt"
)

func (c Client) GetMe() (User, error) {
	body, err := c.get(fmt.Sprintf("%s/~api/users/me", c.Url))
	if err != nil {
		return User{}, err
	}

	user := User{}
	err = json.NewDecoder(body).Decode(&user)
	body.Close()

	return user, err
}

func (c Client) GetUserMemberships(id int) ([]UserMemebership, error) {
	body, err := c.get(fmt.Sprintf("%s/~api/users/%d/memberships", c.Url, id))
	if err != nil {
		return []UserMemebership{}, err
	}

	usermemberships := []UserMemebership{}
	err = json.NewDecoder(body).Decode(&usermemberships)
	body.Close()

	return usermemberships, err
}
