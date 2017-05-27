# Playing with Github API with GO-GITHUB Golang library

G[o-github](https://github.com/google/go-github) library is a client library
that provides an easy way to interact with the [GitHub
API](http://developer.github.com/v3/). This library is being initially developed
for an internal application at Google, so API methods will likely be implemented
in the order that they are needed by that application. We can track the status
of implementation in [this Google
spreadsheet](https://docs.google.com/spreadsheet/ccc?key=0ApoVX4GOiXr-dGNKN1pObFh6ek1DR2FKUjBNZ1FmaEE&usp=sharing).
[API documentation](https://developer.github.com/v3/) is hosted on
[godoc.org](https://godoc.org/github.com/google/go-github/github).

This quick-start demonstrates how to use Go-GitHub library in Golang application
and get GitHub data and display on console or terminal.

We’ll cover:

* Prerequisites for this tutorial
* Github API Rate Limit
* Setting up your application
* Go-Github client with Authentication
* Get specific repository Info
* Github last commit information of specific repository
* Get README.MD Content
* Get Rate Limit Information

### Prerequisites for the Golang tutorial

1.  Basic knowledge of [GO ](https://golang.org/)language
1.  Azure subscription. (If you don’t have an Azure subscription, create a [free
account](https://azure.microsoft.com/free/?WT.mc_id=A261C142F) before you
begin.)
1.  IDE — [Gogland](https://www.jetbrains.com/go/) by Jetbrains or [Visual Studio
Code](https://code.visualstudio.com/) by Mircosoft or [Atom](https://atom.io/)

### Github API Rate Limit

To ensure a high quality of service for all API consumers, github team has
reduced the default rate limit for unauthenticated requests. Unauthenticated
requests will be limited to 60 per hour. To enjoy the default rate limit of
5,000 requests per hour, you’ll need
to[authenticate](https://developer.github.com/v3/#authentication) via Basic Auth
or OAuth.

### How to get Github OAuth2 Access Token

1.  Login to [Github.com](https://github.com/)
1.  Go to user’s settings

<span class="figcaption_hack">User settings link</span>

3. Go to [Personal Access Token](https://github.com/settings/tokens) and click
on Generate new token

<span class="figcaption_hack">Generate new token</span>

4. Fill Token description and check **repo **checkbox.

<span class="figcaption_hack">Token required information and scopes</span>

5. Get/Copy token. This token will be used to Go-Github library

### Setting up your application

It’s time to make our hands dirty. Open your favorite editor (Gogland, VS Code
or Atom). For this article, I will use Gogland editor.

1.  Create folder GoGithub folder inside GOROOT\src folder
1.  Get go-github package

    go get github.com/google/go-github/github

### Go-Github Client

Following code snippet will be used to get Go-Github client

    import "github.com/google/go-github/github"

    client := github.NewClient(nil)

### Go-Github client with Authentication

The go-github library does not directly handle authentication. Instead, when
creating a new client, pass an  that can handle authentication for you. The
easiest and recommended way to do this is using the
OA[uth2](https://github.com/golang/oauth2) library, but you can always use any
other library that provides an . If you have an OAuth2 access token (for
example, a [personal API
token](https://github.com/blog/1509-personal-api-tokens)), you can use it with
the OAuth2 library using:

    import "golang.org/x/oauth2"
    func main() {
      ctx := context.Background()
      ts := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: "... your access token ..."},
      )
      tc := oauth2.NewClient(ctx, ts)

      // get go-github client
      client := github.NewClient(tc)
    }

For API methods that require HTTP Basic Authentication, use the.

### Get specific repository Info

**client.Repositories.Get(ctx context.Context, owner, repo string) (*Repository,
*Response, error) **code will be used to get repository information. Please
refer below code snippet to get repository information. **Get() ***returns
pointer to repository.*

### Github last commit information of specific repository

**client.Repositories.ListCommits(ctx context.Context, owner, repo string, opt
*CommitsListOptions) ([]*RepositoryCommit, *Response, error) **api will provide
you list of commits by descending order. At a time, this api will return maximum
30 commits information. This default behavior can be changed by passing
**CommitsListOptions** arguments.

    // ListOptions specifies the optional parameters to various List methods that
    // support pagination.
    ListOptions 
    {
           // For paginated result sets, page of results to retrieve.
           Page int `url:"page,omitempty"`
           // For paginated result sets, the number of results to include per page.
           PerPage int `url:"per_page,omitempty"`
    }

To get last commit information, take first item from result as shown below:

<span class="figcaption_hack">Get last commit information</span>

### Get README.MD Content

**client.Repositories.GetReadme(ctx context.Context, owner, repo string, opt
*RepositoryContentGetOptions) (*RepositoryContent, *Response, error) ***api will
be used to get README.MD information. Please refer below code snippet:*

### Get Rate Limit Information

**client.RateLimits(ctx context.Context) (*RateLimits, *Response, error) **api
will be used to get rate limit information. Example as follows:

### Get the complete Golang tutorial solution

Please have a look at the entire source code at
[GitHub](https://github.com/Golang-Coach/Lessons/tree/master/GoGithub).

* [Golang](https://medium.com/tag/golang?source=post)
* [Github](https://medium.com/tag/github?source=post)
* [Github Api](https://medium.com/tag/github-api?source=post)
* [Gogithub](https://medium.com/tag/gogithub?source=post)
* [Go](https://medium.com/tag/go?source=post)

### [Durgaprasad Budhwani](https://medium.com/@durgaprasadbudhwani)
