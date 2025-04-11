package mailman

type MailmanConfig struct {
	Port  int
	Pass  string
	Email string
	Host  string
}

func (c *MailmanConfig) validate() error {

	if c.Port == 0 || c.Pass == "" || c.Email == "" || c.Host == "" {
		return ErrInvalidMailmanConfig
	}

	return nil
}
