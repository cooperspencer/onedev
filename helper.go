package onedev

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func trimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func (c Client) get(url string) (io.ReadCloser, int, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 0, err
	}
	if len(c.Token) > 0 {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	} else {
		req.SetBasicAuth(c.Username, c.Password)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return resp.Body, resp.StatusCode, err
}

func (c Client) delete(url string) (io.ReadCloser, int, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, 0, err
	}
	if len(c.Token) > 0 {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	} else {
		req.SetBasicAuth(c.Username, c.Password)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return resp.Body, resp.StatusCode, err
}

func (c Client) post(url string, payload io.Reader) (io.ReadCloser, int, error) {
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, 0, err
	}
	if len(c.Token) > 0 {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	} else {
		req.SetBasicAuth(c.Username, c.Password)
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return resp.Body, resp.StatusCode, err
}
