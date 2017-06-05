package services

import (
	"Lessons/GoBDD/mocks"
	. "github.com/onsi/ginkgo"
	"context"
	. "github.com/google/go-github/github"
	. "github.com/onsi/gomega"
	"errors"
)

var _ = Describe("Github API ", func() {
	var (
		repositoryServices *mocks.IRepositoryServices
		github             Github
		backgroundContext  context.Context
	)

	BeforeEach(func() {
		backgroundContext = context.Background()
		repositoryServices = new(mocks.IRepositoryServices)
		github = NewGithub(repositoryServices, backgroundContext)
	})

	It("should return repository information", func() {
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
		Ω(pack.ForksCount).Should(Equal(starCount))
	})

	It("should return error when failed to retrieve  repository information", func() {
		repositoryServices.On("Get", backgroundContext, "golang-coach", "Lessons").Return(nil, nil, errors.New("Error has been occurred"))
		_, err := github.GetPackageRepoInfo("golang-coach", "Lessons")
		Ω(err).Should(HaveOccurred())
	})

})
