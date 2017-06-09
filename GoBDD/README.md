# Golang Best Practices — Behavior-driven development and Continuous Integration

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

### Walk through of Sample Code

It is highly recommended to go through [Playing with Github API with GO-GITHUB
Golang
library](https://medium.com/@durgaprasadbudhwani/playing-with-github-api-with-go-github-golang-library-83e28b2ff093)*
*article. Below Code snippet call Github API and get repository information.

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

2. Github.go

3. Main.go

Folder structure will be as follows:

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

### How to use mocks:

We can create an instance of the mock struct *IRepositoryServices,*

    repositoryServices := new(mocks.IRepositoryServices)

and we can mock function *Get() as shown below:*

    repositoryServices.On("Get", backgroundContext, "golang-coach", "Lessons").Return(repo, nil, nil)

When unit test code request for *Get() *function with mentioned parameters, it
will return* repo *object. The example will be:

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

It will navigate to composer page as show below:

Left side is used to write feature set like :

    TestGithubAPI 
           Should return repository information 
           Should return error when failed to retrieve repository information

and right side, it will generate *_test code snippet as shown below.

Copy code into *github_test.go* and use mocks classes as shown below:

Whenever you make any changes in the source file, GoConvey will run test cases
and the result will be visible in GoConvey site ( *localhost:8080*). You can
also set browser notification.

Let’s have quick look at how we can integrate build pipeline to this source
code. For continuous integration, we will use [Travis-CI](http://travis-ci.org/)

### [Travis-CI](http://travis-ci.org/) Setup

1.  Login to [Travis-CI](http://travis-ci.org/)
1.  Enable Github Repository

3. Create *.travis.yml* and commit this file in Github repository.

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
