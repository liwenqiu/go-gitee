package gitee

import (
	"fmt"
	"net/http"
	"time"
)

type TagsService struct {
	client *Client
}

type Tag struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Commit  struct {
		SHA  string    `json:"sha"`
		Date time.Time `json:"date"`
	} `json:"commit"`
}

func (t Tag) String() string {
	return Stringify(t)
}

func (s *TagsService) ListTags(owner, repo string) ([]*Tag, *Response, error) {

	u := fmt.Sprintf("repos/%s/%s/tags", pathEscape(owner), pathEscape(repo))

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var t []*Tag
	resp, err := s.client.Do(req, &t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, err
}
