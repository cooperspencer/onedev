package onedev

func NewClient(url, username, password string) *Client {
	return &Client{Username: username, Password: password, Url: trimSuffix(url, "/")}
}

func NewClientWithToken(url, token string) *Client {
	return &Client{Token: token, Url: trimSuffix(url, "/")}
}
