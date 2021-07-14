package onedev

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c Client) GetPullRequest(id int) (PullRequest, error) {
	body, err := c.get(fmt.Sprintf("%s/api/pull-requests/%d", c.Url, id))
	if err != nil {
		return PullRequest{}, err
	}

	pullrequest := PullRequest{}
	err = json.NewDecoder(body).Decode(&pullrequest)
	body.Close()

	return pullrequest, err
}

func (c Client) DeletePullRequest(id int) error {
	_, err := c.delete(fmt.Sprintf("%s/api/pull-requests/%d", c.Url, id))
	if err != nil {
		return err
	}
	return nil
}

func (c Client) GetPullRequests() ([]PullRequest, error) {
	body, err := c.get(fmt.Sprintf("%s/api/pull-requests", c.Url))
	if err != nil {
		return []PullRequest{}, err
	}

	pullrequests := []PullRequest{}
	err = json.NewDecoder(body).Decode(&pullrequests)
	body.Close()

	return pullrequests, err
}

func (c Client) GetPullRequestMergePreview(id int) (MergePreview, error) {
	body, err := c.get(fmt.Sprintf("%s/api/pull-requests/%d/merge-preview", c.Url, id))
	if err != nil {
		return MergePreview{}, err
	}

	mergepreview := MergePreview{}
	err = json.NewDecoder(body).Decode(&mergepreview)
	body.Close()

	return mergepreview, err
}

func (c Client) GetPullRequestAssignments(id int) ([]Assignment, error) {
	body, err := c.get(fmt.Sprintf("%s/api/pull-requests/%d/assignment", c.Url, id))
	if err != nil {
		return []Assignment{}, err
	}

	assignments := []Assignment{}
	err = json.NewDecoder(body).Decode(&assignments)
	body.Close()

	return assignments, err
}

func (c Client) GetPullRequestReviews(id int) ([]Review, error) {
	body, err := c.get(fmt.Sprintf("%s/api/pull-requests/%d/reviews", c.Url, id))
	if err != nil {
		return []Review{}, err
	}

	reviews := []Review{}
	err = json.NewDecoder(body).Decode(&reviews)
	body.Close()

	return reviews, err
}

func (c Client) GetPullRequestComments(id int) ([]Comment, error) {
	body, err := c.get(fmt.Sprintf("%s/api/pull-requests/%d/comments", c.Url, id))
	if err != nil {
		return []Comment{}, err
	}

	comments := []Comment{}
	err = json.NewDecoder(body).Decode(&comments)
	body.Close()

	return comments, err
}

func (c Client) GetPullRequestWatches(id int) ([]Watch, error) {
	body, err := c.get(fmt.Sprintf("%s/api/pull-requests/%d/watches", c.Url, id))
	if err != nil {
		return []Watch{}, err
	}

	watches := []Watch{}
	err = json.NewDecoder(body).Decode(&watches)
	body.Close()

	return watches, err
}

func (c Client) GetPullRequestUpdates(id int) ([]Update, error) {
	body, err := c.get(fmt.Sprintf("%s/api/pull-requests/%d/updates", c.Url, id))
	if err != nil {
		return []Update{}, err
	}

	updates := []Update{}
	err = json.NewDecoder(body).Decode(&updates)
	body.Close()

	return updates, err
}

func (c Client) GetPullRequestCurrentBuild(id int) ([]Build, error) {
	body, err := c.get(fmt.Sprintf("%s/api/pull-requests/%d/current-builds", c.Url, id))
	if err != nil {
		return []Build{}, err
	}

	builds := []Build{}
	err = json.NewDecoder(body).Decode(&builds)
	body.Close()

	return builds, err
}

func (c Client) GetPullRequestChanges(id int) ([]Change, error) {
	body, err := c.get(fmt.Sprintf("%s/api/pull-requests/%d/changes", c.Url, id))
	if err != nil {
		return []Change{}, err
	}

	changes := []Change{}
	err = json.NewDecoder(body).Decode(&changes)
	body.Close()

	return changes, err
}

func (c Client) GetPullRequestFixedIssueNumbers(id int) ([]int, error) {
	body, err := c.get(fmt.Sprintf("%s/api/pull-requests/%d/fixed-issue-numbers", c.Url, id))
	if err != nil {
		return []int{}, err
	}

	fixed := []int{}
	err = json.NewDecoder(body).Decode(&fixed)
	body.Close()

	return fixed, err
}

func (c Client) CreatePullRequest(options PullRequestOptions) (int, error) {
	payloadbytes, err := json.Marshal(options)
	if err != nil {
		return 0, err
	}
	payload := bytes.NewReader(payloadbytes)
	body, err := c.post(fmt.Sprintf("%s/api/pull-requests", c.Url), payload)
	if err != nil {
		return 0, err
	}

	id := 0
	err = json.NewDecoder(body).Decode(&id)
	body.Close()

	return id, err
}

func (c Client) SetPullRequestTitle(id int, title string) error {
	payloadbytes, err := json.Marshal(title)
	if err != nil {
		return err
	}
	payload := bytes.NewReader(payloadbytes)
	body, err := c.post(fmt.Sprintf("%s/api/pull-requests/%d/title", c.Url, id), payload)
	if err != nil {
		return err
	}

	body.Close()

	return nil
}

func (c Client) SetPullRequestDescription(id int, title string) error {
	payloadbytes, err := json.Marshal(title)
	if err != nil {
		return err
	}
	payload := bytes.NewReader(payloadbytes)
	body, err := c.post(fmt.Sprintf("%s/api/pull-requests/%d/description", c.Url, id), payload)
	if err != nil {
		return err
	}

	body.Close()

	return nil
}

func (c Client) SetPullRequestMergeStrategy(id int, strategy string) error {
	payloadbytes, err := json.Marshal(strategy)
	if err != nil {
		return err
	}
	payload := bytes.NewReader(payloadbytes)
	body, err := c.post(fmt.Sprintf("%s/api/pull-requests/%d/merge-strategy", c.Url, id), payload)
	if err != nil {
		return err
	}

	body.Close()

	return nil
}

func (c Client) ReopenPullRequest(id int, reopen string) error {
	payloadbytes, err := json.Marshal(reopen)
	if err != nil {
		return err
	}
	payload := bytes.NewReader(payloadbytes)
	body, err := c.post(fmt.Sprintf("%s/api/pull-requests/%d/reopen", c.Url, id), payload)
	if err != nil {
		return err
	}

	body.Close()

	return nil
}

func (c Client) DiscardPullRequest(id int, discard string) error {
	payloadbytes, err := json.Marshal(discard)
	if err != nil {
		return err
	}
	payload := bytes.NewReader(payloadbytes)
	body, err := c.post(fmt.Sprintf("%s/api/pull-requests/%d/discard", c.Url, id), payload)
	if err != nil {
		return err
	}

	body.Close()

	return nil
}

func (c Client) MergePullRequest(id int, merge string) error {
	payloadbytes, err := json.Marshal(merge)
	if err != nil {
		return err
	}
	payload := bytes.NewReader(payloadbytes)
	body, err := c.post(fmt.Sprintf("%s/api/pull-requests/%d/merge", c.Url, id), payload)
	if err != nil {
		return err
	}

	body.Close()

	return nil
}

func (c Client) DeletePullRequestSourceBranch(id int, branch string) error {
	payloadbytes, err := json.Marshal(branch)
	if err != nil {
		return err
	}
	payload := bytes.NewReader(payloadbytes)
	body, err := c.post(fmt.Sprintf("%s/api/pull-requests/%d/delete-source-branch", c.Url, id), payload)
	if err != nil {
		return err
	}

	body.Close()

	return nil
}

func (c Client) RestorePullRequestSourceBranch(id int, branch string) error {
	payloadbytes, err := json.Marshal(branch)
	if err != nil {
		return err
	}
	payload := bytes.NewReader(payloadbytes)
	body, err := c.post(fmt.Sprintf("%s/api/pull-requests/%d/restore-source-branch", c.Url, id), payload)
	if err != nil {
		return err
	}

	body.Close()

	return nil
}
