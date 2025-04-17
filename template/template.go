package template

import (
	"log"
	"strings"
)

type MessageTemplateInfo struct {
	Name    string   `json:"name"`
	Path    string   `json:"path"`
	Fields  []string `json:"fields"`
	Subject string   `json:"subject"`
}

type MessageTemplate struct {
	Name            string
	templateContent string
	fields          []string
	Subject         string
}

func NewMessageTemplate(info *MessageTemplateInfo) (*MessageTemplate, error) {

	var (
		templateContent string
		err             error
	)

	if templateContent, err = readFileContent(info.Path); err != nil {
		return nil, err
	}
	if err = checkFileContentByFields(templateContent, info.Fields); err != nil {
		return nil, err
	}

	return &MessageTemplate{
		Name:            info.Name,
		templateContent: templateContent,
		fields:          info.Fields,
		Subject:         info.Subject,
	}, nil
}

func (m *MessageTemplate) checkField(field string) bool {

	for _, f := range m.fields {
		if f == field {
			return true
		}
	}

	return false
}

func (m *MessageTemplate) PrepareMessageContent(fieldValues map[string]string) string {

	var (
		content string = m.templateContent
	)

	for field, value := range fieldValues {

		if !m.checkField(field) {
			log.Printf("mailman template warning : field %s doesn't exist in template %s\n", field, m.Name)
			continue
		}

		content = strings.Replace(content, "{{."+field+"}}", value, 1)
	}

	return content
}
