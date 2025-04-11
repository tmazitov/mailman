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

	err = m.SetupMailTemplates([]*template.MessageTemplateInfo{
		{
			Name:     "test",
			FilePath: "template.html",
			Fields:   []string{"firstName", "lastName"},
		},
	})
	if err != nil {
		panic(err)
	}

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
