package template

import "errors"

var (
	ErrNotEnoughFields = errors.New("mailman template error : template doesn't follow required fields")
	ErrFieldsNotUnique = errors.New("mailman template error : template fields are not unique")
)
