package mailman

import (
	"log"

	"github.com/tmazitov/mailman/template"
	"gopkg.in/gomail.v2"
)

func (m *Mailman) worker(messageInfoChan chan *MessageInfo) {

	var (
		info    *MessageInfo
		message *gomail.Message
		err     error
	)

	for info = range messageInfoChan {
		if message, err = m.makeMessage(info); err != nil {
			log.Println(err)
			continue
		}
		if err = m.send(message); err != nil {
			log.Println(err)
		}
	}
}

func (m *Mailman) send(message *gomail.Message) error {
	return gomail.
		NewDialer(m.config.Host, m.config.Port, m.config.Email, m.config.Pass).
		DialAndSend(message)
}

func (m *Mailman) makeMessage(info *MessageInfo) (*gomail.Message, error) {

	var (
		message  *gomail.Message
		template *template.MessageTemplate
		content  string
	)

	template = m.templates[info.TemplateName]
	if template == nil {
		return nil, ErrUndefinedTemplate
	}

	content = template.PrepareMessageContent(info.FieldValues)

	message = gomail.NewMessage()
	message.SetHeader("From", m.config.Email)
	message.SetHeader("To", info.DistEmail)
	message.SetHeader("Subject", info.Subject)
	message.SetBody("text/html", content)
	return message, nil
}
