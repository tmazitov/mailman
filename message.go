package mailman

type MessageInfo struct {
	TemplateName string            `json:"templateName"`
	Subject      string            `json:"subject"`
	DistEmail    string            `json:"distEmail"`
	FieldValues  map[string]string `json:"fieldValues"`
	Content      string            `json:"content"`
}

func (m *Mailman) SendMessage(info *MessageInfo) error {

	if info == nil || info.DistEmail == "" || (info.TemplateName == "" && info.Content == "") {
		return ErrInvalidMessageInfo
	}

	m.emailChan <- info

	return nil
}
