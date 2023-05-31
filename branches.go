package gitee

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type BranchesService struct {
	client *Client
}

type BranchRef struct {
	Name          string     `json:"name"`
	Commit        *CommitRef `json:"commit"`
	Protected     bool       `json:"protected"`
	ProtectionURL string     `json:"protection_url"`
}

func (br BranchRef) String() string {
	return Stringify(br)
}

func (s *BranchesService) ListBranches(owner, repo string) ([]*BranchRef, *Response, error) {

	u := fmt.Sprintf("repos/%s/%s/branches", pathEscape(owner), pathEscape(repo))

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var b []*BranchRef
	resp, err := s.client.Do(req, &b)
	if err != nil {
		return nil, resp, err
	}

	return b, resp, err
}

type Branch struct {
	Name   string `json:"name"`
	Commit *struct {
		SHA    string `json:"sha"`
		URL    string `json:"url"`
		Commit *struct {
			Author *struct {
				Name  string     `json:"name"`
				Date  *time.Time `json:"date"`
				Email string     `json:"email"`
			} `json:"author"`
			URL     string `json:"url"`
			Message string `json:"message"`
			Tree    *struct {
				SHA string `json:"sha"`
				URL string `json:"url"`
			} `json:"tree"`
			Committer *struct {
				Name  string     `json:"name"`
				Date  *time.Time `json:"date"`
				Email string     `json:"email"`
			} `json:"committer"`
		} `json:"commit"`
		Author *struct {
			AvatarURL string `json:"avatar_url"`
			URL       string `json:"url"`
			ID        int64  `json:"id"`
			Login     string `json:"login"`
		} `json:"author"`
		Parents []*struct {
			SHA string `json:"sha"`
			URL string `json:"url"`
		} `json:"parents"`
		Committer *struct {
			AvatarURL string `json:"avatar_url"`
			URL       string `json:"url"`
			ID        int64  `json:"id"`
			Login     string `json:"login"`
		} `json:"committer"`
	} `json:"commit"`
	Protected     bool   `json:"protected"`
	ProtectionURL string `json:"protection_url"`
}

func (b Branch) String() string {
	return Stringify(b)
}

func (s *BranchesService) GetBranch(owner, repo, branch string) (*Branch, *Response, error) {

	u := fmt.Sprintf("repos/%s/%s/branches/%s", pathEscape(owner), pathEscape(repo), url.PathEscape(branch))

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	b := new(Branch)
	resp, err := s.client.Do(req, b)
	if err != nil {
		return nil, resp, err
	}

	return b, resp, err
}

func (s *BranchesService) CreateBranch(owner, repo, refs, branch_name string) (*Branch, *Response, error) {
	u := fmt.Sprintf("repos/%s/%s/branches", pathEscape(owner), pathEscape(repo))
	opt := struct {
		Refs       string `json:"refs"`
		BranchName string `json:"branch_name"`
	}{
		Refs:       refs,
		BranchName: branch_name,
	}
	req, err := s.client.NewRequest(http.MethodPost, u, &opt)
	if err != nil {
		return nil, nil, err
	}
	b := new(Branch)
	resp, err := s.client.Do(req, b)
	if err != nil {
		return nil, resp, err
	}
	return b, resp, err
}
