package git

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Buscador interface {
	BuscaGitTag(repo string)
}

type BuscadorGit struct {
	Repo              string
	DonoDoRepositorio string
}

type JSONDoGithub struct {
	FullName    string   `json:"full_name"`
	HTMLUrl     string   `json:"html_url"`
	Description string   `json:"description"`
	Topics      []string `json:"topics"`
}

// Declarando uma função para buscar os commits
func (b *BuscadorGit) BuscaGitTag(tag string) {
	fmt.Println("Buscando commits pertecentes a git tag...")
	fmt.Printf("%s %s %s", tag, b.DonoDoRepositorio, b.Repo+"\n")

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", b.DonoDoRepositorio, b.Repo)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode == 200 {
		b, _ := io.ReadAll(resp.Body)

		j := new(JSONDoGithub)
		json.Unmarshal(b, j)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(j.FullName)
	}
}
