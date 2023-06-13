package onedev

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func (c Client) GetMe() (User, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/users/me", c.Url))
	if err != nil {
		return User{}, status, err
	}

	user := User{}
	err = json.NewDecoder(body).Decode(&user)
	body.Close()

	return user, status, err
}

func (c Client) GetUserMemberships(id int) ([]UserMemebership, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/users/%d/memberships", c.Url, id))
	if err != nil {
		return []UserMemebership{}, status, err
	}

	usermemberships := []UserMemebership{}
	err = json.NewDecoder(body).Decode(&usermemberships)
	body.Close()

	return usermemberships, status, err
}

func (c Client) CreateUser(options CreateUserOptions) (int, int, error) {
	payloadbytes, err := json.Marshal(options)
	if err != nil {
		return 0, 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, status, err := c.post(fmt.Sprintf("%s/~api/users", c.Url), payload)
	if err != nil {
		return 0, status, err
	}
	defer body.Close()
	b, err := io.ReadAll(body)
	if err != nil {
		return 0, status, err
	}

	id := 0
	err = json.NewDecoder(bytes.NewReader(b)).Decode(&id)
	if err != nil {
		err = errors.New(string(b))
	}

	return id, status, err
}
