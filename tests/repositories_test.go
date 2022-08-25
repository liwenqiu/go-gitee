package tests

import (
	"fmt"
	"io"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/liwenqiu/go-gitee"
)

func TestRepositoriesService_List(t *testing.T) {
	hc := retryablehttp.NewClient()
	c, _ := gitee.NewOAuthClient(hc, "replace your access_token")
	vis := gitee.VisibilityValue("all")
	owner := "owner"
	args := &gitee.ListRepositoriesOptions{
		Visibility:  &vis,
		Affiliation: &owner,
	}
	r, s, err := c.Repositories.List(args)
	if err != nil {
		t.Fatalf(err.Error())
	}
	for _, re := range r {
		fullName, humanName := (*re).FullName, (*re).HumanName
		fmt.Println(*fullName, *humanName)
	}
	x, err := io.ReadAll((*s).Body)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(string(x))
}
