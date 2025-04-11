package mailman

type MessageInfo struct {
	TemplateName string
	Subject      string
	DistEmail    string
	FieldValues  map[string]string
	Content      string
}

func (m *Mailman) SendMessage(info *MessageInfo) error {

	if info == nil || info.DistEmail == "" || (info.TemplateName == "" && info.Content == "") {
		return ErrInvalidMessageInfo
	}

	m.emailChan <- info

	return nil
}
