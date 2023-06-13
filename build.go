package onedev

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c Client) GetBuilds(options *BuildQueryOptions) ([]Build, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/~api/builds", c.Url), nil)

	q := req.URL.Query()
	if options.Query != "" {
		q.Add("query", options.Query)
	}
	q.Add("offset", fmt.Sprintf("%d", options.Offset))
	if options.Count == 0 {
		options.Count = 100
	}
	q.Add("count", fmt.Sprintf("%d", options.Count))

	req.URL.RawQuery = q.Encode()

	body, status, err := c.get(req.URL.String())
	if err != nil {
		return []Build{}, status, err
	}

	builds := []Build{}
	err = json.NewDecoder(body).Decode(&builds)
	body.Close()

	return builds, status, err
}

func (c Client) GetBuild(id int) (Build, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/builds/%d", c.Url, id))
	if err != nil {
		return Build{}, status, err
	}

	build := Build{}
	err = json.NewDecoder(body).Decode(&build)
	body.Close()

	return build, status, err
}

func (c Client) DeleteBuild(id int) (int, error) {
	_, status, err := c.delete(fmt.Sprintf("%s/~api/builds/%d/", c.Url, id))
	if err != nil {
		return status, err
	}
	return status, nil
}
