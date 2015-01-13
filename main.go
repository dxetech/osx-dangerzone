package main

import (
	"flag"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"log"
	"strings"
)

var username = flag.String("username", "", "Sendgrid username")
var apiKey = flag.String("sendgrid_key", "", "Sendgrid api key")
var to = flag.String("to", "", "Email address to field. Use semicolon for multiple addresses")
var from = flag.String("from", "", "Email address from field")
var additionalText = flag.String("text", "", "Additional message text")

func main() {
	flag.Parse()
	if len(*username) == 0 {
		log.Fatalln("Username required")
	}
	if len(*apiKey) == 0 {
		log.Fatalln("Api key required")
	}
	if len(*to) == 0 {
		log.Fatalln("To field required")
	}
	if len(*from) == 0 {
		log.Fatalln("From field required")
	}
	sg := sendgrid.NewSendGridClient(*username, *apiKey)
	message := sendgrid.NewMail()
	for _, t := range strings.Split(*to, ";") {
		message.AddTo(t)
		message.AddToName(t)
	}
	message.SetSubject("OS X Server Error Occurred")
	text := ""
	if len(*additionalText) > 0 {
		text = fmt.Sprintf("An error occured: %s", *additionalText)
	} else {
		text = "An error occured"
	}
	message.SetText(text)
	message.SetFrom(*from)
	if err := sg.Send(message); err == nil {
		log.Println("Email sent!")
	} else {
		fmt.Println(err)
	}
}
