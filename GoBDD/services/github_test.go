package services

import (
	"Lessons/GoBDD/mocks"
	"context"
	. "github.com/smartystreets/goconvey/convey"
	. "github.com/google/go-github/github"
	"testing"
	"errors"
)

func TestGithubAPI(t *testing.T) {
	Convey("Should return repository information", t, func() {
		backgroundContext := context.Background()
		repositoryServices := new(mocks.IRepositoryServices)
		github := NewGithub(repositoryServices, backgroundContext)

		fullName := "ABC"
		starCount := 10
		repo := &Repository{
			FullName:        &fullName,
			Description:     &fullName,
			ForksCount:      &starCount,
			StargazersCount: &starCount,

		}
		repositoryServices.On("Get", backgroundContext, "golang-coach", "Lessons").Return(repo, nil, nil)
		pack, _ := github.GetPackageRepoInfo("golang-coach", "Lessons")
		So(pack.ForksCount, ShouldEqual, starCount)
	})

	Convey("Should return error when failed to retrieve  repository information", t, func() {
		backgroundContext := context.Background()
		repositoryServices := new(mocks.IRepositoryServices)
		github := NewGithub(repositoryServices, backgroundContext)
		repositoryServices.On("Get", backgroundContext, "golang-coach", "Lessons").Return(nil, nil, errors.New("Error has been occurred"))
		_, err := github.GetPackageRepoInfo("golang-coach", "Lessons")
		So(err, ShouldNotBeEmpty)
	})
}
