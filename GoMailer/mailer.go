package main

import (
	"gopkg.in/mailgun/mailgun-go.v1"
)

// SendMail is used to send message, it will ask user about title, htmlMessage, textMessage and list of recipient
func SendMail(title string, htmlMessage string, textMessage string, receivers ...string, ) (string, error) {

	// NewMailGun creates a new client instance.
	mg := mailgun.NewMailgun("golang.coach", "MailGun API Key", "MailGun Public API Key")

	// Create message
	message := mg.NewMessage(
		"no-reply@golang.coach",
		title,
		textMessage,
		receivers...,
	)

	// If you're sending a message that isn't already MIME encoded, SetHtml() will arrange to bundle
	// an HTML representation of your message in addition to your plain-text body.
	message.SetHtml(htmlMessage)

	// AddAttachment arranges to send a file from the filesystem along with the e-mail message.
	// The attachment parameter is a filename, which must refer to a file which actually resides
	// in the local filesystem.
	//message.AddAttachment("files/test.jpg")

	// can add cc and bcc user email address
	//message.AddCC("baz@example.com")
	//message.AddBCC("bar@example.com")

	// send message and get result
	resp, _, err := mg.Send(message)

	return resp, err
}
