package onedev

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c Client) GetPullRequest(id int) (PullRequest, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/pull-requests/%d", c.Url, id))
	if err != nil {
		return PullRequest{}, status, err
	}

	pullrequest := PullRequest{}
	err = json.NewDecoder(body).Decode(&pullrequest)
	body.Close()

	return pullrequest, status, err
}

func (c Client) DeletePullRequest(id int) (int, error) {
	_, status, err := c.delete(fmt.Sprintf("%s/~api/pull-requests/%d", c.Url, id))
	if err != nil {
		return status, err
	}
	return status, nil
}

func (c Client) GetPullRequests() ([]PullRequest, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/pull-requests", c.Url))
	if err != nil {
		return []PullRequest{}, status, err
	}

	pullrequests := []PullRequest{}
	err = json.NewDecoder(body).Decode(&pullrequests)
	body.Close()

	return pullrequests, status, err
}

func (c Client) GetPullRequestMergePreview(id int) (MergePreview, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/pull-requests/%d/merge-preview", c.Url, id))
	if err != nil {
		return MergePreview{}, status, err
	}

	mergepreview := MergePreview{}
	err = json.NewDecoder(body).Decode(&mergepreview)
	body.Close()

	return mergepreview, status, err
}

func (c Client) GetPullRequestAssignments(id int) ([]Assignment, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/pull-requests/%d/assignment", c.Url, id))
	if err != nil {
		return []Assignment{}, status, err
	}

	assignments := []Assignment{}
	err = json.NewDecoder(body).Decode(&assignments)
	body.Close()

	return assignments, status, err
}

func (c Client) GetPullRequestReviews(id int) ([]Review, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/pull-requests/%d/reviews", c.Url, id))
	if err != nil {
		return []Review{}, status, err
	}

	reviews := []Review{}
	err = json.NewDecoder(body).Decode(&reviews)
	body.Close()

	return reviews, status, err
}

func (c Client) GetPullRequestComments(id int) ([]Comment, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/pull-requests/%d/comments", c.Url, id))
	if err != nil {
		return []Comment{}, status, err
	}

	comments := []Comment{}
	err = json.NewDecoder(body).Decode(&comments)
	body.Close()

	return comments, status, err
}

func (c Client) GetPullRequestWatches(id int) ([]Watch, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/pull-requests/%d/watches", c.Url, id))
	if err != nil {
		return []Watch{}, status, err
	}

	watches := []Watch{}
	err = json.NewDecoder(body).Decode(&watches)
	body.Close()

	return watches, status, err
}

func (c Client) GetPullRequestUpdates(id int) ([]Update, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/pull-requests/%d/updates", c.Url, id))
	if err != nil {
		return []Update{}, status, err
	}

	updates := []Update{}
	err = json.NewDecoder(body).Decode(&updates)
	body.Close()

	return updates, status, err
}

func (c Client) GetPullRequestCurrentBuild(id int) ([]Build, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/pull-requests/%d/current-builds", c.Url, id))
	if err != nil {
		return []Build{}, status, err
	}

	builds := []Build{}
	err = json.NewDecoder(body).Decode(&builds)
	body.Close()

	return builds, status, err
}

func (c Client) GetPullRequestChanges(id int) ([]Change, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/pull-requests/%d/changes", c.Url, id))
	if err != nil {
		return []Change{}, 0, err
	}

	changes := []Change{}
	err = json.NewDecoder(body).Decode(&changes)
	body.Close()

	return changes, status, err
}

func (c Client) GetPullRequestFixedIssueNumbers(id int) ([]int, int, error) {
	body, status, err := c.get(fmt.Sprintf("%s/~api/pull-requests/%d/fixed-issue-numbers", c.Url, id))
	if err != nil {
		return []int{}, status, err
	}

	fixed := []int{}
	err = json.NewDecoder(body).Decode(&fixed)
	body.Close()

	return fixed, status, err
}

func (c Client) CreatePullRequest(options PullRequestOptions) (int, int, error) {
	payloadbytes, err := json.Marshal(options)
	if err != nil {
		return 0, 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, status, err := c.post(fmt.Sprintf("%s/~api/pull-requests", c.Url), payload)
	if err != nil {
		return 0, status, err
	}

	id := 0
	err = json.NewDecoder(body).Decode(&id)
	body.Close()

	return id, status, err
}

func (c Client) SetPullRequestTitle(id int, title string) (int, error) {
	payloadbytes, err := json.Marshal(title)
	if err != nil {
		return 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, status, err := c.post(fmt.Sprintf("%s/~api/pull-requests/%d/title", c.Url, id), payload)
	if err != nil {
		return status, err
	}

	body.Close()

	return status, nil
}

func (c Client) SetPullRequestDescription(id int, title string) (int, error) {
	payloadbytes, err := json.Marshal(title)
	if err != nil {
		return 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, status, err := c.post(fmt.Sprintf("%s/~api/pull-requests/%d/description", c.Url, id), payload)
	if err != nil {
		return status, err
	}

	body.Close()

	return status, nil
}

func (c Client) SetPullRequestMergeStrategy(id int, strategy string) (int, error) {
	payloadbytes, err := json.Marshal(strategy)
	if err != nil {
		return 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, status, err := c.post(fmt.Sprintf("%s/~api/pull-requests/%d/merge-strategy", c.Url, id), payload)
	if err != nil {
		return status, err
	}

	body.Close()

	return status, nil
}

func (c Client) ReopenPullRequest(id int, reopen string) (int, error) {
	payloadbytes, err := json.Marshal(reopen)
	if err != nil {
		return 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, status, err := c.post(fmt.Sprintf("%s/~api/pull-requests/%d/reopen", c.Url, id), payload)
	if err != nil {
		return status, err
	}

	body.Close()

	return status, nil
}

func (c Client) DiscardPullRequest(id int, discard string) (int, error) {
	payloadbytes, err := json.Marshal(discard)
	if err != nil {
		return 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, status, err := c.post(fmt.Sprintf("%s/~api/pull-requests/%d/discard", c.Url, id), payload)
	if err != nil {
		return status, err
	}

	body.Close()

	return status, nil
}

func (c Client) MergePullRequest(id int, merge string) (int, error) {
	payloadbytes, err := json.Marshal(merge)
	if err != nil {
		return 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, status, err := c.post(fmt.Sprintf("%s/~api/pull-requests/%d/merge", c.Url, id), payload)
	if err != nil {
		return status, err
	}

	body.Close()

	return status, nil
}

func (c Client) DeletePullRequestSourceBranch(id int, branch string) (int, error) {
	payloadbytes, err := json.Marshal(branch)
	if err != nil {
		return 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, status, err := c.post(fmt.Sprintf("%s/~api/pull-requests/%d/delete-source-branch", c.Url, id), payload)
	if err != nil {
		return status, err
	}

	body.Close()

	return status, nil
}

func (c Client) RestorePullRequestSourceBranch(id int, branch string) (int, error) {
	payloadbytes, err := json.Marshal(branch)
	if err != nil {
		return 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, status, err := c.post(fmt.Sprintf("%s/~api/pull-requests/%d/restore-source-branch", c.Url, id), payload)
	if err != nil {
		return status, err
	}

	body.Close()

	return status, nil
}
