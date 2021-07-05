package onedev

func NewClient(url, username, password string) *Client {
	return &Client{Username: username, Password: password, Url: trimSuffix(url, "/")}
}
