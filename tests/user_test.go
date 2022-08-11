package test

import (
	"fmt"
	"io"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/liwenqiu/go-gitee"
)

func TestUsersService_List(t *testing.T) {
	hc := retryablehttp.NewClient()
	c, _ := gitee.NewOAuthClient(hc, "replace your access_token")

	r, s, err := c.Users.Get("abc")
	if err != nil {
		t.Fatalf(err.Error())
	}
	avatar_url, id, html_url, name := (*r).AvatarURL, (*r).ID, (*r).HTMLURL, (*r).Name

	fmt.Println(*avatar_url, *id, *html_url, *name)
	fmt.Println("-------------------------------")
	x, err := io.ReadAll((*s).Body)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(string(x))
}
