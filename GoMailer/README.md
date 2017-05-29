# Sending Complex HTML email with Mailgun API using Go language

In this tutorial, we are going to learn how to send newsletter HTML mail to
recipients using Mailgun API in Go language.

Why Mailgun?

Mailgun provides powerful APIs that enable you to
[send](http://mgprod.wpengine.com/features/sending-email/),
[receive](http://mgprod.wpengine.com/features/inbound-routing/) and
[track](http://mgprod.wpengine.com/features/reports/) email effortlessly. It
provides 10,000 emails per month. Mailgun is very easy to integrate.

In this tutorials, we’ll cover:

* Prerequisites for this tutorial
* How to get MailGun Private and Public Key
* Newsletter template by [limtus.com](https://litmus.com/community/templates/)
* Setting up your application
* Template Parser
* Mailgun Client — Mailer

### Prerequisites for this tutorial

1.  Basic knowledge of [GO ](https://golang.org/)language
1.  IDE — [Gogland](https://www.jetbrains.com/go/) by Jetbrains or [Visual Studio
Code](https://code.visualstudio.com/) by Mircosoft or [Atom](https://atom.io/)

### How to get Mailgun Private and Public Key

1.  [Login ](https://app.mailgun.com/sessions/new)or [Signup
](https://app.mailgun.com/new/signup/)to Mailgun
1.  Scroll down to bottom, you will see Domains and API keys information

![](https://cdn-images-1.medium.com/max/800/1*x3ESh-inKm1l5txIThnrfQ.png)

Copy domain name, Public API Key and Secret API Key. It will be used by Golang
application

Note: — You can add a custom domain, but you need to [verify
domain](https://help.mailgun.com/hc/en-us/articles/202052074-How-do-I-verify-my-domain-).

### Newsletter template by [limtus.com](https://litmus.com/community/templates/)

For this application, we will use
[limtus.com](https://litmus.com/community/templates/) free
[11-pook-newsletter](https://litmus.com/community/templates/11-pook-newsletter)
template link is shared below.

### Setting up your application

It’s time to make our hands dirty. Open your favorite editor (Gogland, VS Code
or Atom). For this article, I will use Gogland editor.

1.  Create folder GoMailer folder inside GOROOT\src folder
1.  Run below command to get mailgun-go package

```go
  go get gopkg.in/mailgun/mailgun-go.v1
```

### Template Parser

[html/template](https://golang.org/pkg/html/template/) implements data-driven
templates for generating HTML output safe against code injection. It provides
the same interface as package text/template and should be used instead of
text/template whenever the output is HTML.

Below code snippet will take template name and data as input and return compiled
template content and error as output. If *templateFilePath *and data required by
the template are correct then it will parse the template and returns content in
string format.

```go
package main

import (
	"bytes"
	"html/template"
)

// ParseTemplate returns parsed template with data in string format
// If there is an error, it will return response with error data
func ParseTemplate(templateFilePath string, data interface{}) (content string, err error) {

	// ParseFiles creates a new Template and parses the template definitions from
	// the named files. The returned template's name will have the (base) name and
	// (parsed) contents of the first file. There must be at least one file.
	// If an error occurs, parsing stops and the returned *Template is nil.
	tmpl, err := template.ParseFiles(templateFilePath)
	if err != nil {
		return "", err
	}

	// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
	// The zero value for Buffer is an empty buffer ready to use.
	buf := new(bytes.Buffer)

	// Execute applies a parsed template to the specified data object,
	// writing the output to wr.
	// If an error occurs executing the template or writing its output,
	// execution stops, but partial results may already have been written to
	// the output writer.
	// A template may be executed safely in parallel.
	if err := tmpl.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
```

Example -

Below snippet show example of newsletter template.
```
Golang Coach Newsletter
This is an opinionated catalog of open source Golang packages.
{{range .}}
{{.Content}}
{{.Url}}
{{end}}
```
This template will receive array of *Article *struct
```go
  type Article struct {
      Name string
      Image   string
      Content string
      Url     string
  }
```
**ParseTemplate **function will receive *file path *of newsletter template and
*array of Article* and returns parsed text content as shown below
```go
package main

import (
	"fmt"
	"path/filepath"
	"os"
	"log"
)

func main() {

	// List of articles
	news := []Article{
		{
			Name: "AuthBoss",
			Image:   "https://camo.githubusercontent.com/6ee0a52cc24492a1be493f87fed66ffd51d12f01/687474703a2f2f692e696d6775722e636f6d2f66504967714c672e6a7067",
			Content: "Authboss is a modular authentication system for the web. It tries to remove as much boilerplate and \"hard things\" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time. ",
			Url:     "https://github.com/go-authboss/authboss",
		},
		{
			Name: "ConnectorDB",
			Image:   "https://raw.githubusercontent.com/connectordb/connectordb/master/screenshot.png",
			Content: "ConnectorDB is an open-source Life Server. You can use it to gather data about your productivity, mood, sync with your android devices, and gather data from a whole host of sensors in real-time. ConnectorDB can be used to manage the smart home, and control your IoT devices, all while gathering data for use in analysis as well as in automated control.",
			Url:     "https://connectordb.io/",
		},
	}

	cwd, _ := os.Getwd()
	var err error

	// Get newsletter text template in string format
	textContent, err := ParseTemplate(filepath.Join(cwd, "templates", "./newsletter.tmpl"), news)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(textContent)
	
}
```
### Mailgun Client — Mailer

The mailer code is as given below:
```go
package main

import (
	"gopkg.in/mailgun/mailgun-go.v1"
)

type MailRequest struct {
	from    string
	title   string
	subject string
	htmlMessage    string
	plainTextMessage string
	to []string
}

// Create new mail request
func NewMailRequest(from string, title string, htmlMessage string, textMessage string, receivers []string) *MailRequest {
	return &MailRequest{
		from:  from,
		title: title,
		htmlMessage:htmlMessage,
		plainTextMessage:textMessage,
		to: receivers,
	}

}

// SendMail is used to send message, it will ask user about title, htmlMessage, textMessage and list of recipient
func (mailRequest *MailRequest) SendMail() (bool, error) {

	// NewMailGun creates a new client instance.
	mg := mailgun.NewMailgun("golang.coach", "MailGun API Key", "MailGun Public API Key")

	// Create message
	message := mg.NewMessage(
		mailRequest.from,
		mailRequest.title,
		mailRequest.plainTextMessage,
		mailRequest.to...,
	)

	// If you're sending a message that isn't already MIME encoded, SetHtml() will arrange to bundle
	// an HTML representation of your message in addition to your plain-text body.
	message.SetHtml(mailRequest.htmlMessage)

	// AddAttachment arranges to send a file from the filesystem along with the e-mail message.
	// The attachment parameter is a filename, which must refer to a file which actually resides
	// in the local filesystem.

	//message.AddAttachment("files/test.jpg")

	// can add cc and bcc user email address
	//message.AddCC("baz@example.com")
	//message.AddBCC("bar@example.com")


	// send message and get result
	if _, _, err :=  mg.Send(message); err != nil {
		return false, err
	}

	return true, nil
}
```
In this code, Mailgun API **mailgun.NewMailgun(“golang.coach”, “MailGun API
Key”, “MailGun Public API Key”)** has been encapsulated in a *struct*
**MailRequest**. It contains basic things like from, title, htmlMessage,
textMessage and receivers list.

The MailRequest has method **SendMail()** which will take MailRequest data and
send mail request using Mailgun API. When mail has been sent successfully, this
method will return true with no error.

You can extend this function or struct to add mail attachment or CC or BCC list
as shown in the comment.

```go
// AddAttachment arranges to send a file from the filesystem along with the e-mail message.
// The attachment parameter is a filename, which must refer to a file which actually resides
// in the local filesystem.

// message.AddAttachment("files/test.jpg")

// can add cc and bcc user email address
// message.AddCC("baz@example.com")
// message.AddBCC("bar@example.com")
```



To use *MailRequest *and send mail, the code is as given below:
```go
mailRequest := NewMailRequest(
   "no-reply@golang.coach",
   "Golang Coach Newsletter",
   htmlContent,
   textContent,
   []string {"durgaprasad.budhwani@gmail.com"},
)

ok, err := mailRequest.SendMail()
if !ok {
       log.Fatal(err)
}
fmt.Println("Mail has been sent")
```
### Get the complete solution

Please have a look at the entire source code at
[GitHub](https://github.com/Golang-Coach/Lessons/tree/master/GoMailer).

All suggestions are welcome.Thanks for your time for reading this article.
