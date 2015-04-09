package providers

import (
	"testing"
)

func Test_request(t *testing.T) {
	g := NewGitlab("https://git.xiaojukeji.com/api/v3", "KXGLy5yctc7sUkr4oshh")
	u, e := g.CurrentUser()
	if u != "lilinghui" || e != nil {
		t.Error("get user is not lilinghui")
	}
}
