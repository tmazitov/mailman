package mailman

import "errors"

var (
	ErrInvalidMailmanConfig = errors.New("mailman error : invalid mailman config")
	ErrUndefinedTemplate   = errors.New("mailman error : undefined template name")
	ErrInvalidMessageInfo  = errors.New("mailman error : invalid message info")
)
