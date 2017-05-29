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
