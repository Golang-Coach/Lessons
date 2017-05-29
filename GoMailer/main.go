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
	// Get newsletter html template in string format
	htmlContent, err := ParseTemplate(filepath.Join(cwd, "templates", "./newsletter.html"), news)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(htmlContent)

	// Get newsletter text template in string format
	textContent, err := ParseTemplate(filepath.Join(cwd, "templates", "./newsletter.tmpl"), news)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(textContent)

	mailRequest := NewMailRequest(
		"no-reply@golang.coach",
		"Golang Coach Newsletter",
		htmlContent,
		textContent,
		[]string {"durgaprasad.budhwani@gmail.com"},
	)

	// send mail
	ok, err := mailRequest.SendMail()
	if !ok {
		log.Fatal(err)
	}

	fmt.Println("Mail has been sent")
}
