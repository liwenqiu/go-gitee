package gitee

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type CommitsService struct {
	client *Client
}

type CommitRef struct {
	SHA string `json:"sha"`
	URL string `json:"url"`
}

type Commit struct {
	URL         string `json:"url"`
	SHA         string `json:"sha"`
	HTMLURL     string `json:"html_url"`
	CommentsURL string `json:"comment_url"`
	Commit      *struct {
		Author *struct {
			Name  string     `json:"name"`
			Date  *time.Time `json:"date"`
			Email string     `json:"email"`
		} `json:"author"`
		Committer *struct {
			Name  string     `json:"name"`
			Date  *time.Time `json:"date"`
			Email string     `json:"email"`
		} `json:"committer"`
		Message string `json:"message"`
		Tree    *struct {
			SHA string `json:"sha"`
			URL string `json:"url"`
		} `json:"tree"`
	} `json:"commit"`
	Author *struct {
		ID                int64  `json:"id"`
		Login             string `json:"login"`
		Name              string `json:"name"`
		AvatarURL         string `json:"avatar_url"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		Remark            string `json:"remark"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
	} `json:"author"`
	Committer *struct {
		ID                int64  `json:"id"`
		Login             string `json:"login"`
		Name              string `json:"name"`
		AvatarURL         string `json:"avatar_url"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		Remark            string `json:"remark"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
	} `json:"committer"`
	Parents []*struct {
		SHA string `json:"sha"`
		URL string `json:"url"`
	} `json:"parents"`
	Stats *struct {
		ID        string `json:"id"`
		Additions int    `json:"additions"`
		Deletions int    `json:"deletions"`
		Total     int    `json:"total"`
	} `json:"stats"`
	Files []*struct {
		SHA        string `json:"sha"`
		Filename   string `json:"filename"`
		Status     string `json:"status"`
		Additions  int    `json:"additions"`
		Deletions  int    `json:"deletions"`
		Changes    int    `json:"changes"`
		BlobURL    string `json:"blob_url"`
		RawURL     string `json:"raw_url"`
		ContentURL string `json:"content_url"`
		Patch      string `json:"patch"`
	} `json:"files"`
}

func (c Commit) String() string {
	return Stringify(c)
}

type ListCommitsOptions struct {
	ListOptions
	SHA    string `json:"sha,omitempty"`
	Path   string `json:"path,omitempty"`
	Author string `json:"author,omitempty"`
}

func (s *CommitsService) ListCommits(owner, repo string, opts *ListCommitsOptions) ([]*Commit, *Response, error) {

	u := fmt.Sprintf("repos/%s/%s/commits", pathEscape(owner), pathEscape(repo))

	req, err := s.client.NewRequest(http.MethodGet, u, opts)
	if err != nil {
		return nil, nil, err
	}

	var c []*Commit
	resp, err := s.client.Do(req, &c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}

// GetCommit 404 {message: Commit not found}
func (s *CommitsService) GetCommit(owner, repo, sha string) (*Commit, *Response, error) {

	if sha == "" {
		return nil, nil, fmt.Errorf("SHA must be a non-empty string")
	}
	u := fmt.Sprintf("repos/%s/%s/commits/%s", pathEscape(owner), pathEscape(repo), url.PathEscape(sha))

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	c := new(Commit)
	resp, err := s.client.Do(req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, err
}
