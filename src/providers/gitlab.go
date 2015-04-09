package providers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Gitlab struct {
	url   string
	c     *http.Client
	token string
}

func NewGitlab(remote, token string) (g *Gitlab) {
	g = new(Gitlab)
	g.url = remote
	g.token = token
	g.c = &http.Client{}
	return g
}

func (g *Gitlab) request(url, cmd string) (results map[string]interface{}, err error) {
	real_url := g.url + url
	req, _ := http.NewRequest(cmd, real_url, nil)
	//req.Header.Add("PRIVATE-TOKEN", "KXGLy5yctc7sUkr4oshh")
	req.Header.Add("PRIVATE-TOKEN", g.token)
	var r *http.Response
	r, err = g.c.Do(req)
	fmt.Println("=============")
	fmt.Println(r)
	fmt.Println("=============")
	decoder := json.NewDecoder(r.Body)
	var v interface{}
	if err = decoder.Decode(&v); err == nil {
		var ok bool
		if results, ok = v.(map[string]interface{}); ok {
			return
		}
	}
	return
}

func (g *Gitlab) CurrentUser() (u string, err error) {
	var r map[string]interface{}
	if r, err = g.request("/user", "GET"); err == nil {
		ui := r["username"]
		var ok bool
		if u, ok = ui.(string); !ok {
			err = errors.New("result is not string type")
		}
	}
	return
}
