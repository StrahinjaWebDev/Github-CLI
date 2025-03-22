package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	url2 "net/url"
)

func main() {
	fmt.Println(githubInfo("StrahinjaWebDev"))
}

// github info returns name number public repos
func githubInfo(login string) (string, int, error) {
	url := `https://api.github.com/users/` + url2.PathEscape(login)
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%#v - %s", url, resp.Status)
	}

	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))

	var r struct {
		Name        string
		PublicRepos int `json:"public_repos"`
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}
	return r.Name, r.PublicRepos, nil

}

type Reply struct {
	Name        string
	PublicRepos int `json:"public_repos"`
}
