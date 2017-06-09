# Golang Best Practices — Behavior-driven development and Continuous Integration [![Build Status](https://travis-ci.org/Golang-Coach/Lessons.svg?branch=master)](https://travis-ci.org/Golang-Coach/Lessons) [![Coverage Status](https://coveralls.io/repos/github/Golang-Coach/Lessons/badge.svg?branch=master)](https://coveralls.io/github/Golang-Coach/Lessons?branch=master)

![](https://cdn-images-1.medium.com/max/1200/0*IEC0BK3MjNcf3O7h.png)

In [software engineering](https://en.wikipedia.org/wiki/Software_engineering),
**behavior-driven development** ( **BDD**) is a [software development
process](https://en.wikipedia.org/wiki/Software_development_process) that
emerged from [test-driven
development](https://en.wikipedia.org/wiki/Test-driven_development) (TDD). The
behavior-driven-development combines the general techniques and principles of
TDD with ideas from [domain-driven
design](https://en.wikipedia.org/wiki/Domain-driven_design) and [object-oriented
analysis and
design](https://en.wikipedia.org/wiki/Object-oriented_analysis_and_design) to
provide software development and management teams with shared tools and a shared
process to collaborate on software development. (
[Wiki](https://en.wikipedia.org/wiki/Behavior-driven_development) link)

The focus of BDD is the language and interactions used in the process of
software development. Behavior-driven developers use their native language in
combination with the ubiquitous language of Domain Driven Design to describe the
purpose and benefit of their code. This allows the developers to focus on why
the code should be created, rather than the technical details and minimizes
translation between the technical language in which the code is written and the
domain language spoken by the business, users, stakeholders, project management
etc.

This article will give you a quick introduction on how to get started with BDD
(Behavior Driven Development) and Continuous integration in Golang.

For BDD, we will use [GoConvey](http://goconvey.co/) and for CI we will use
[Travis-CI](http://travis-ci.org/)

We’ll cover:

* Prerequisites for this tutorial
* Setting up your application
* Walkthrough of Sample Code
* Single Responsibility Principle
* Interface for loose coupling
* [Mockery](https://github.com/vektra/mockery) to generate mocks
* [GoConvey](http://goconvey.co/) for BDD
* [Travis-CI](http://travis-ci.org/) setup
* [Coveralls.io](https://coveralls.io/) setup

### Prerequisites for the Golang tutorial

1.  Basic knowledge of [GO ](https://golang.org/)language
1.  IDE — [Gogland](https://www.jetbrains.com/go/) by Jetbrains or [Visual Studio
Code](https://code.visualstudio.com/) by Microsoft or [Atom](https://atom.io/)
1.  Go through [Playing with Github API with GO-GITHUB Golang
library](https://medium.com/@durgaprasadbudhwani/playing-with-github-api-with-go-github-golang-library-83e28b2ff093)**
**article.

### Setting up your application

It’s time to make our hands dirty. Open your favorite editor (Gogland, VS Code
or Atom). For this article, I will use Gogland editor.

1.  Create folder GoBDD folder inside GOROOT\src folder
1.  Get following Golang packages

```sh
    go get github.com/google/go-github
    go get github.com/stretchr/testify
    go get github.com/smartystreets/goconvey
    go get github.com/onsi/ginkgo/ginkgo
    go get github.com/modocache/gover
    go get github.com/vektra/mockery
```

### Walk through of Sample Code

It is highly recommended to go through [Playing with Github API with GO-GITHUB
Golang
library](https://medium.com/@durgaprasadbudhwani/playing-with-github-api-with-go-github-golang-library-83e28b2ff093)*
*article. Below Code snippet call Github API and get repository information.

```go
package main

import (
	"github.com/google/go-github/github"
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
	
    // Step 1 - create github client
	client := github.NewClient(nil)

    // step 2 - Service call to get repo information
	repo, _, err := client.Repositories.Get(context, "Golang-Coach", "Lessons")

	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}

    // Step 3 - bind result to Package Model
	pack := &Package{
		FullName: *repo.FullName,
		Description: *repo.Description,
		ForksCount: *repo.ForksCount,
		StarsCount: *repo.StargazersCount,
	}

	fmt.Printf("%+v\n", pack)
}
```

To get Github repo information, steps are very simple:

1.  Create Github client — **github.NewClient()**
1.  Service call to get repo information **client.Repositories.Get(context,
“Golang-Coach”, “Lessons”)**
1.  Bind repo result to Package

The above code snippet has few drawbacks:

* The single main function is doing everything, it is calling services and binding
result to Package model. This can be overcome by using [Single Responsibility
Principle](https://en.wikipedia.org/wiki/Single_responsibility_principle)
* Because of tight coupling between *the third party library
***(github.com/google/go-github)** and *main *function, the code is not testable
and would be hard to maintain in long run. This can be overcome by using
[Interface based
segregation](https://en.wikipedia.org/wiki/Interface_segregation_principle)[Dependency
Injection](https://en.wikipedia.org/wiki/Dependency_inversion_principle)

### Single Responsibility Principle (SRP)

In SRP, each class will handle only one responsibility. In above code snippet,
we can move Github API Call and Package Model code into the separate class or
struct as shown below.

1.  Package.go

```go
// File name - github.com/Golang-Coach/Lessons/GoBDD/models/package.go
package models

import "time"

// Package : here you tell us what Salutation is
type Package struct {
    FullName      string
    Description   string
    StarsCount    int
    ForksCount    int
    UpdatedAt     time.Time
    LastUpdatedBy string
    ReadMe        string
    Tags          []string
    Categories    []string
}
```

2. Github.go

```go
// File name - github.com/Golang-Coach/Lessons/GoBDD/services/github.go
package services

import (
    "context"
    "github.com/Golang-Coach/Lessons/GoBDD/models"
    "github.com/google/go-github/github"
)

// Github : This struct will be used to get Github related information
type Github struct {
    repositoryServices *github.RepositoriesService
    context            context.Context
}

// NewGithub : It will intialized Github class
func NewGithub(context context.Context, repositoryServices *github.RepositoriesService) Github {
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
```

3. Main.go

```go
package main

import (
    "github.com/google/go-github/github"
    "github.com/Golang-Coach/Lessons/GoBDD/services"
    "context"
    "fmt"   
)

func main() {
    context := context.Background()    
    client := github.NewClient(nil)   

    // Step 1 - create github api client
    githubAPI := services.NewGithub(context, client.Repositories)

    // Step 1 - Get Repository Package Information
    pack, err := githubAPI.GetPackageRepoInfo("Golang-Coach", "Lessons")
    fmt.Printf("%+v\n", pack)
    fmt.Printf("%+v\n", err)
}
```

Folder structure will be as follows:
![](https://cdn-images-1.medium.com/max/800/1*4vwa9aKkIhei0zKffqqqbQ.png)

Now we have segregated Github Service call and *package *binding logic from the
*main function.*

When we write test cases against above code, it will make actual service call
and will get the result from Github. To avoid the actual service call, we need
to mock Github service.

### Interface for loose coupling

In Go language, the function can be mocked by Interface approach. It is the
nature of the Interfaces to **provide many implementations, thus enable
mocking**. **Instead of actually calling** a dependent system or even a module,
or a **complicated and difficult to instantiate the type**, you can **provide
the simplest interface implementation** that will provide results needed **for
the unit test to complete correctly. **Code
**service.repositoryServices.Get(service.context, owner, repositoryName)**make
service call. To mock this,we need to create interface as shown below:

```go
// IRepositoryServices : This interface will be used to provide loose coupling between github.RepositoryServices and its consumer
type IRepositoryServices interface {
	Get(ctx context.Context, owner, repo string) (*github.Repository, *github.Response, error)
}
```

### [Mockery](https://github.com/vektra/mockery) to generate mocks

Interfaces are naturally super good integration points for tests since the
implementation of an interface can easily be replaced by a mock implementation.
However, writing mocks can be quite tedious and boring. To make life easier
mockery provides the ability to easily generate mocks for Golang interfaces. It
removes the boilerplate coding required to use mocks.

To mock *IRepositoryServices* interface, we need to run below command:

    mockery -name=IRepositoryServices

It will create mocks folder at the root level and also create
*IRepositoryServices.go* file.

```go
// Code generated by mockery v1.0.0
package mocks

import context "context"
import github "github.com/google/go-github/github"
import mock "github.com/stretchr/testify/mock"

// IRepositoryServices is an autogenerated mock type for the IRepositoryServices type
type IRepositoryServices struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, owner, repo
func (_m *IRepositoryServices) Get(ctx context.Context, owner string, repo string) (*github.Repository, *github.Response, error) {
	ret := _m.Called(ctx, owner, repo)

	var r0 *github.Repository
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *github.Repository); ok {
		r0 = rf(ctx, owner, repo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Repository)
		}
	}

	var r1 *github.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, string) *github.Response); ok {
		r1 = rf(ctx, owner, repo)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(ctx, owner, repo)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
```

### How to use mocks:

We can create an instance of the mock struct *IRepositoryServices,*

    repositoryServices := new(mocks.IRepositoryServices)

and we can mock function *Get() as shown below:*

    repositoryServices.On("Get", backgroundContext, "golang-coach", "Lessons").Return(repo, nil, nil)

When unit test code request for *Get() *function with mentioned parameters, it
will return* repo *object. The example will be:

```go
backgroundContext := context.Background()
repositoryServices := new(mocks.IRepositoryServices)
github := NewGithub(backgroundContext, repositoryServices)

fullName := "ABC"
starCount := 10
repo := &Repository{
	FullName:        &fullName,
	Description:     &fullName,
	ForksCount:      &starCount,
	StargazersCount: &starCount,

}
repositoryServices.On("Get", backgroundContext, "golang-coach", "Lessons").Return(repo, nil, nil)
```

GoConvey is an extension of the built-in Go test tool. It facilitates
[Behavior-driven Development
(BDD)](https://en.wikipedia.org/wiki/Behavior-driven_development) in Go, though
this is not the only way to use it. Many people continue to write traditional Go
tests but prefer [GoConvey’s web
UI](https://github.com/smartystreets/goconvey/wiki/Web-UI) for reporting test
results.

### Installation

    go get github.com/smartystreets/goconvey

Start up the GoConvey web server at your project’s path:

    // for linux or mac
    $GOPATH/bin/goconvey 

    // for windows
    %GOPATH%/bin/goconvey

Then watch the test results display in your browser at:

    http:localhost:8080

### GoConvey Composer

The goconvey composer will be useful to generate feature description to Convey
Code Snippet. Click on edit icon on GoConvery browser window as shown below
![](https://cdn-images-1.medium.com/max/800/1*_BbUdWa7Qyr2Tu3ljTuSPg.png)
It will navigate to composer page as show below:
![](https://cdn-images-1.medium.com/max/800/1*UGkMMInqxQnQY3PuRt1jqQ.png)

Left side is used to write feature set like :

    TestGithubAPI 
           Should return repository information 
           Should return error when failed to retrieve repository information

and right side, it will generate *_test code snippet as shown below.
```go
import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGithubAPI(t *testing.T) {

	Convey("Should return repository information", t, nil)

	Convey("Should return error when failed to retrieve  repository information", t, nil)

}
```

Copy code into *github_test.go* and use mocks classes as shown below:
```go
package services

import (
	"github.com/Golang-Coach/Lessons/GoBDD/mocks"
	"context"
	. "github.com/smartystreets/goconvey/convey"
	. "github.com/google/go-github/github"
	"testing"
	"errors"
)

func TestGithubAPI(t *testing.T) {
	Convey("Should return repository information", t, func() {
		backgroundContext := context.Background()
        // create mock of IRepositoryServices interface
		repositoryServices := new(mocks.IRepositoryServices)
        // pass mocked object in NewGithub constructor/func
		github := NewGithub(backgroundContext, repositoryServices)

		fullName := "ABC"
		starCount := 10
		repo := &Repository{
			FullName:        &fullName,
			Description:     &fullName,
			ForksCount:      &starCount,
			StargazersCount: &starCount,

		}
        // when code calls Get method of IRepositoryServices, it will return repo mocked object 
		repositoryServices.On("Get", backgroundContext, "golang-coach", "Lessons").Return(repo, nil, nil)
		pack, _ := github.GetPackageRepoInfo("golang-coach", "Lessons")
        // assert
		So(pack.ForksCount, ShouldEqual, starCount)
	})

	Convey("Should return error when failed to retrieve  repository information", t, func() {
		backgroundContext := context.Background()
		repositoryServices := new(mocks.IRepositoryServices)
		github := NewGithub(backgroundContext, repositoryServices)
		repositoryServices.On("Get", backgroundContext, "golang-coach", "Lessons").Return(nil, nil, errors.New("Error has been occurred"))
		_, err := github.GetPackageRepoInfo("golang-coach", "Lessons")
		So(err, ShouldNotBeEmpty)
	})
}
```

Whenever you make any changes in the source file, GoConvey will run test cases
and the result will be visible in GoConvey site ( *localhost:8080*). You can
also set browser notification.

Let’s have quick look at how we can integrate build pipeline to this source
code. For continuous integration, we will use [Travis-CI](http://travis-ci.org/)

### [Travis-CI](http://travis-ci.org/) Setup

1.  Login to [Travis-CI](http://travis-ci.org/)
1.  Enable Github Repository
![](https://cdn-images-1.medium.com/max/800/1*4cEMZRkhDa7ZiuiKNCJtIg.png)

3. Create *.travis.yml* and commit this file in Github repository.
```yml
language: go

go:
 - tip

before_install:
 - go get golang.org/x/tools/cmd/cover
 - go get github.com/axw/gocov/gocov
 - go get github.com/mattn/goveralls
 - go get github.com/onsi/ginkgo/ginkgo
 - go get github.com/modocache/gover

script:
 - $HOME/gopath/bin/ginkgo -r  -cover -coverpkg=./GoBDD/services/... -trace -race
 - $HOME/gopath/bin/gover
 - $HOME/gopath/bin/goveralls -coverprofile=gover.coverprofile -repotoken "Your Security Token"

```

GoConvey provides you coverage information during developer machine, but it is
also important that during the build, we should also get build status and
coverage information.

*go test -c -coverpkg* is only supported coverage of only one package. For
multiple package coverage, we will use **ginkgo -r -cover**

Below snippet will collect coverage from different packages and generate
*package_name.coverprofile* file.

    $HOME/gopath/bin/ginkgo -r -cover -coverpkg=./GoBDD/services/... -trace -race

To collect coverages from all packages, below code snippet has been used

    $HOME/gopath/bin/gover

### [Coveralls.io](https://coveralls.io/) setup

1.  Login to [coveralls.io](https://coveralls.io/sign-in)
1.  Add repository
![](https://cdn-images-1.medium.com/max/800/1*PvFYAe88MM3wGh_WJ_ehmA.png)

3. Click on details and get *repo_token*

Modify .travis.yml file and put *repo_token, it will push coverage to
coveralls.io site*

    $HOME/gopath/bin/goveralls -coverprofile=gover.coverprofile -repotoken "Your Security Token"

After build is successful, you will see status at travis-ci site <br>
[https://travis-ci.org/Golang-Coach/Lessons](https://travis-ci.org/Golang-Coach/Lessons)
and coverage report at
[https://coveralls.io/repos/github/Golang-Coach/Lessons/](https://coveralls.io/repos/github/Golang-Coach/Lessons/badge.svg?branch=master')

*****

### Get the complete Golang tutorial solution

Please have a look at the entire source code at
[GitHub](https://github.com/Golang-Coach/Lessons/tree/master/GoBDD).

* [Continuous
Integration](https://medium.com/tag/continuous-integration?source=post)

### [Durgaprasad Budhwani](https://medium.com/@durgaprasadbudhwani)
