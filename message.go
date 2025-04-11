package mailman

type MessageInfo struct {
	TemplateName string
	Subject      string
	DistEmail    string
	FieldValues  map[string]string
}

func (m *Mailman) SendMessage(info *MessageInfo) error {

	if info == nil {
		return ErrInvalidMessageInfo
	}

	m.emailChan <- info

	return nil
}