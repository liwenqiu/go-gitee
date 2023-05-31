package gitee

import (
	"fmt"
	"net/http"
	"strings"
)

type UsersService struct {
	client *Client
}

type GiteeUser struct {
	ID                *int64  `json:"id"`
	Login             *string `json:"login"`
	Name              *string `json:"name"`
	AvatarURL         *string `json:"avatar_url"`
	URL               *string `json:"url"`
	HTMLURL           *string `json:"html_url"`
	Remark            *string `json:"remark"`
	FollowersURL      *string `json:"followers_url"`
	FollowingURL      *string `json:"following_url"`
	GistsURL          *string `json:"gists_url"`
	StarredURL        *string `json:"starred_url"`
	SubscriptionsURL  *string `json:"subscriptions_url"`
	OrganizationsURL  *string `json:"organizations_url"`
	ReposURL          *string `json:"repos_url"`
	EventsURL         *string `json:"events_url"`
	ReceivedEventsURL *string `json:"received_events_url"`
	Type              *string `json:"type"`
}

type User struct {
	GiteeUser
	Weibo       *string `json:"weibo"`
	Bio         *string `json:"bio"`
	Blog        *string `json:"blog"`
	Company     *string `json:"company"`
	Linkedin    *string `json:"linkedin"`
	Email       *string `json:"email"`
	Follers     *int    `json:"followers"`
	Following   *int    `json:"following"`
	MemberRole  *string `json:"member_role"`
	PublicGists *int    `json:"public_gists"`
	PublicRepos *int    `json:"public_repos"`
	Stared      *int    `json:"stared"`
	Watched     *int    `json:"watched"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

// Get fetches a user. Passing the empty string will fetch the authenticated user.
func (s *UsersService) Get(user string) (*User, *Response, error) {
	url := "user"
	if strings.TrimSpace(user) != "" {
		url = fmt.Sprintf("users/%s", strings.TrimSpace(user))
	}
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}

	var p *User
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}
