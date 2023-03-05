package main

import (
	"fmt"

	"github.com/laudo395/golang-expert/git"
)

func main() {
	tag := "v1.26.2"
	donoDoRepositorio := "kubernetes"
	repo := "kubernetes"

	b := git.BuscadorGit{
		DonoDoRepositorio: donoDoRepositorio,
		Repo:              repo,
	}
	b.BuscaGitTag(tag)

	fmt.Println("Sai do programa")
}
