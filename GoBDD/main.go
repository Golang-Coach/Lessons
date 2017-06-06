package main

import (
	"fmt"
	"context"
	"golang.org/x/oauth2"
	. "github.com/google/go-github/github"
	. "github/Golang-Coach/Lessons/GoBDD/services"
)
func main() {

	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "22ffe92b14c28bf8ec53e7f0102ed240c1e02633"},
	)
	tokenClient := oauth2.NewClient(context, tokenService)
	client := *NewClient(tokenClient)
	githubApi := NewGithub(client.Repositories, context)
	pack, err := githubApi.GetPackageRepoInfo("Golang-coach", "Lessons")
	fmt.Println(pack)
	fmt.Println(err)
}