package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/liwenqiu/go-gitee"
)

func TestListBranches_List(t *testing.T) {
	hc := retryablehttp.NewClient()
	c, _ := gitee.NewOAuthClient(hc, "replace your access_token")
	r, s, err := c.Branches.ListBranches("ombak", "go-algo")
	if err != nil {
		t.Fatalf(err.Error())
	}
	for _, b := range r {
		fmt.Println(b)
	}
	fmt.Println(s.StatusCode)
}

func TestGetBranch_List(t *testing.T) {
	hc := retryablehttp.NewClient()
	c, _ := gitee.NewOAuthClient(hc, "replace your access_token")
	r, s, err := c.Branches.GetBranch("ombak", "go-algo", "master")
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(r.Commit.Commit.Author.Date, r.ProtectionURL)
	fmt.Println(s.StatusCode)
}

func TestCreateBranch_List(t *testing.T) {
	hc := retryablehttp.NewClient()
	c, _ := gitee.NewOAuthClient(hc, "replace your access_token")
	r, s, err := c.Branches.CreateBranch("ombak", "go-algo", "master", "api-test1")
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(r.Commit.Commit.Author.Date, r.ProtectionURL)
	fmt.Println(s.StatusCode)
}
