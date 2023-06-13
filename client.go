package onedev

import "fmt"

type Token string

func (t Token) ToString() string {
	return fmt.Sprint(t)
}

type BasicAuth struct {
	Username string
	Password string
}

func NewClient(url string, params ...interface{}) *Client {
	c := &Client{Url: trimSuffix(url, "/")}
	for _, param := range params {
		if _, ok := param.(Token); ok {
			c.Token = param.(Token).ToString()
		}
		if _, ok := param.(BasicAuth); ok {
			c.Username = param.(BasicAuth).Username
			c.Password = param.(BasicAuth).Password
		}
	}

	return c
}

func SetToken(token string) Token {
	return Token(token)
}

func SetBasicAuth(username, password string) BasicAuth {
	return BasicAuth{
		Username: username,
		Password: password,
	}
}
