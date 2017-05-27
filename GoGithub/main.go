package main

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"context"
	"fmt"
	"os"
)

// Model
type Package struct {
	FullName      string
	Description   string
	StarsCount    int
	ForksCount    int
	LastUpdatedBy string
}

func main() {
	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "fbffb6bc6c30c864db49a9b7b869a35a48d28ec0"},
	)
	tokenClient := oauth2.NewClient(context, tokenService)

	client := github.NewClient(tokenClient)

	repo, _, err := client.Repositories.Get(context, "Golang-Coach", "Lessons")

	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}

	pack := &Package{
		FullName: *repo.FullName,
		Description: *repo.Description,
		ForksCount: *repo.ForksCount,
		StarsCount: *repo.StargazersCount,
	}

	fmt.Printf("%+v\n", pack)

	commitInfo, _, err := client.Repositories.ListCommits(context, "Golang-Coach", "Lessons", nil)

	if err != nil {
		fmt.Printf("Problem in commit information %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", commitInfo[0]) // Last commit information

	// repository readme information
	readme, _, err := client.Repositories.GetReadme(context, "facebook", "react-native", nil)
	if err != nil {
		fmt.Printf("Problem in getting readme information %v\n", err)
		return
	}

	// get content
	content, err := readme.GetContent()
	if err != nil {
		fmt.Printf("Problem in getting readme content %v\n", err)
		return
	}

	fmt.Println(content)

	// Get Rate limit information

	rateLimit, _, err := client.RateLimits(context)
	if err != nil {
		fmt.Printf("Problem in getting rate limit information %v\n", err)
		return
	}

	fmt.Printf("Limit: %d \nRemaining %d \n", rateLimit.Core.Limit, rateLimit.Core.Remaining ) // Last commit information

}
