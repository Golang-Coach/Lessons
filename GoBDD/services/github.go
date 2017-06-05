package services

import (
	. "github.com/google/go-github/github"
	"context"
	. "Lessons/GoBDD/models"
)

type IRepositoryServices interface {
	Get(ctx context.Context, owner, repo string) (*Repository, *Response, error)
}

type IGithub interface {
	GetPackageRepoInfo(owner string, repositoryName string) (*Package, error)
}

type Github struct {
	repositoryServices IRepositoryServices
	context context.Context
}

func NewGithub(repositoryServices IRepositoryServices, context context.Context) Github {
	return Github{
		repositoryServices: repositoryServices,
		context: context,
	}
}

func (service *Github) GetPackageRepoInfo(owner string, repositoryName string) (*Package, error) {
	repo, _, err := service.repositoryServices.Get(service.context, owner, repositoryName)
	if err != nil {
		return nil, err
	}
	pack := &Package{
		FullName:    *repo.FullName,
		Description: *repo.Description,
		ForksCount:   *repo.ForksCount,
		StarsCount:  *repo.StargazersCount,
	}
	return pack, nil
}

