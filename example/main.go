package main

import (
	"os"
	"time"

	"log"

	"github.com/tmazitov/mailman"
	"github.com/tmazitov/mailman/template"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <email> <password>", os.Args[0])
	}

	email := os.Args[1]
	pass := os.Args[2]

	m, err := mailman.NewMailMan(&mailman.MailmanConfig{
		Email: email,
		Pass:  pass,
		Port:  587,
		Host:  "smtp.gmail.com",
	})
	if err != nil {
		panic(err)
	}

	m.Start()
	defer m.Stop()

	err = m.SetupMessageTemplates([]*template.MessageTemplateInfo{
		{
			Name:   "test",
			Path:   "template.html",
			Fields: []string{"firstName", "lastName"},
		},
	})
	if err != nil {
		panic(err)
	}

	// Send simple message
	m.SendMessage(&mailman.MessageInfo{
		Subject:   "Test",
		DistEmail: email,
		Content:   "Hello, world!",
	})

	// Send templated message
	m.SendMessage(&mailman.MessageInfo{
		TemplateName: "test",
		Subject:      "Test",
		DistEmail:    email,
		FieldValues: map[string]string{
			"firstName": "Timur",
			"lastName":  "Mazitov",
		},
	})

	time.Sleep(10 * time.Second)
}
