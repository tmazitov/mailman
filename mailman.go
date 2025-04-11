package mailman

import (
	"fmt"

	"github.com/tmazitov/mailman/template"
)

type Mailman struct {
	isStarted bool
	emailChan chan *MessageInfo
	templates map[string]*template.MessageTemplate
	config    *MailmanConfig
}

func NewMailMan(config *MailmanConfig) (*Mailman, error) {

	if config == nil {
		return nil, ErrInvalidMailmanConfig
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	return &Mailman{
		config:    config,
		isStarted: false,
		emailChan: make(chan *MessageInfo),
		templates: map[string]*template.MessageTemplate{},
	}, nil
}

func (m *Mailman) Start() {

	if m.isStarted {
		return
	}

	go func() {
		m.worker(m.emailChan)
	}()

	m.isStarted = true
}

func (m *Mailman) Stop() {
	if !m.isStarted {
		return
	}

	m.isStarted = false
	close(m.emailChan)
}

func (m *Mailman) SetupMailTemplates(info []*template.MessageTemplateInfo) error {

	var (
		newTemplate *template.MessageTemplate
		err         error
	)

	for _, templateInfo := range info {
		if m.templates[templateInfo.Name] != nil {
			fmt.Printf("mailman waning : template %s already exists\n", templateInfo.Name)
			continue
		}
		newTemplate, err = template.NewMessageTemplate(templateInfo)
		if err != nil {
			return err
		}
		m.templates[templateInfo.Name] = newTemplate
	}

	return nil
}
