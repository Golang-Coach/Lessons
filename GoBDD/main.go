package main

import (
	"context"
	"fmt"

	"github.com/Golang-Coach/Lessons/GoBDD/services"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {

	backgroundContext := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "22ffe92b14c28bf8ec53e7f0102ed240c1e02633"},
	)
	tokenClient := oauth2.NewClient(backgroundContext, tokenService)
	client := *github.NewClient(tokenClient)
	githubAPI := services.NewGithub(backgroundContext, client.Repositories)
	pack, err := githubAPI.GetPackageRepoInfo("Golang-coach", "Lessons")
	fmt.Println(pack)
	fmt.Println(err)
}
