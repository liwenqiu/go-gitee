package gitee

import (
	"fmt"
	"net/http"
)

type RepositoriesService struct {
	client *Client
}

type Compare struct {
	BaseCommit      *Commit
	MergeBaseCommit *Commit
	Commits         []*Commit
	Files           []*struct {
		SHA        string `json:"sha,omitempty"`
		FileName   string `json:"filename,omitempty"`
		Status     string `json:"status,omitempty"`
		Additions  int    `json:"additions,omitempty"`
		Deletions  int    `json:"deletions,omitempty"`
		Changes    int    `json:"changes,omitempty"`
		BlobURL    string `json:"blob_url,omitempty"`
		RawURL     string `json:"raw_url,omitempty"`
		ContentURL string `json:"content_url,omitempty"`
		Patch      string `json:"patch,omitempty"`
	}
}

func (s *RepositoriesService) Compare(owner, repo, base, head string) (*Compare, *Response, error) {

	u := fmt.Sprintf("repos/%s/%s/compare/%s...%s", pathEscape(owner), pathEscape(repo), pathEscape(base), pathEscape(head))

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	c := new(Compare)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, nil
}
