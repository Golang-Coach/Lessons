package services

import (
	"context"

	"github.com/Golang-Coach/Lessons/GoBDD/models"
	"github.com/google/go-github/github"
)

// IRepositoryServices : This interface will be used to provide light coupling between github.RepositoryServices and its consumer
type IRepositoryServices interface {
	Get(ctx context.Context, owner, repo string) (*github.Repository, *github.Response, error)
}

// IGithub : This interface will be used to provide light coupling between GithubAPI and its consumer
type IGithub interface {
	GetPackageRepoInfo(owner string, repositoryName string) (*models.Package, error)
}

// Github : This struct will be used to get Github related information
type Github struct {
	repositoryServices IRepositoryServices
	context            context.Context
}

// NewGithub : It will intialized Github class
func NewGithub(context context.Context, repositoryServices IRepositoryServices) Github {
	return Github{
		repositoryServices: repositoryServices,
		context:            context,
	}
}

// GetPackageRepoInfo : This receiver provide Github related repository information
func (service *Github) GetPackageRepoInfo(owner string, repositoryName string) (*models.Package, error) {
	repo, _, err := service.repositoryServices.Get(service.context, owner, repositoryName)
	if err != nil {
		return nil, err
	}
	pack := &models.Package{
		FullName:    *repo.FullName,
		Description: *repo.Description,
		ForksCount:  *repo.ForksCount,
		StarsCount:  *repo.StargazersCount,
	}
	return pack, nil
}
