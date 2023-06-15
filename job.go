package onedev

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) TriggerJob(options *TriggerJobQueryOptions) (int, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/~api/trigger-job", c.Url), nil)

	q := req.URL.Query()
	if options.Project != "" {
		q.Add("project", options.Project)
	}
	if options.Branch != "" {
		q.Add("branch", options.Branch)
	}
	if options.Tag != "" {
		q.Add("tag", options.Tag)
	}
	if options.AccessToken != "" {
		q.Add("access-token", options.AccessToken)
	}

	req.URL.RawQuery = q.Encode()

	body, status, err := c.get(req.URL.String())
	if err != nil || status != 200 {
		b, _ := io.ReadAll(body)
		return 0, status, fmt.Errorf(string(b))
	}

	id := 0
	err = json.NewDecoder(body).Decode(&id)
	body.Close()

	return id, status, err
}
