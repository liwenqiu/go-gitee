package gitee

import (
	"net/http"
	"strconv"
)

type PullRequestsService struct {
	client *Client
}

type Diff struct {
	Sha       string `json:"sha"`
	Filename  string `json:"filename"`
	Status    string `json:"status"`
	Additions string `json:"additions"`
	Deletions string `json:"deletions"`
	BlobURL   string `json:"blob_url"`
	RawURL    string `json:"raw_url"`
	Patch     struct {
		Diff        string `json:"diff"`
		NewPath     string `json:"new_path"`
		OldPath     string `json:"old_path"`
		AMode       string `json:"a_mode"`
		BMode       string `json:"b_mode"`
		NewFile     bool   `json:"new_file"`
		RenamedFile bool   `json:"renamed_file"`
		DeletedFile bool   `json:"deleted_file"`
		TooLarge    bool   `json:"too_large"`
	} `json:"patch"`
}

type ListPullRequestDiffsOptions struct {
	Owner  string
	Repo   string
	Number int64
}

// ListPullRequestDiffs https://gitee.com/api/v5/swagger#/getV5ReposOwnerRepoPullsNumberFiles
func (s *PullRequestsService) ListPullRequestDiffs(opts *ListPullRequestDiffsOptions) ([]*Diff, *Response, error) {
	url := "repos/" + opts.Owner + "/" + opts.Repo + "/pulls/" + strconv.Itoa(int(opts.Number)) + "/files"
	req, err := s.client.NewRequest(http.MethodGet, url, opts)
	if err != nil {
		return nil, nil, err
	}

	var f []*Diff
	resp, err := s.client.Do(req, &f)
	if err != nil {
		return nil, resp, err
	}

	return f, resp, nil
}

// TODO: CreatePullRequest
// CreatePullRequest https://gitee.com/api/v5/swagger#/postV5ReposOwnerRepoPulls
func (s *PullRequestsService) CreatePullRequest() (*Response, error) {
	return nil, nil
}
