package template

import (
	"os"
	"strings"
)

func readFileContent(path string) (string, error) {
	var (
		err  error
		data []byte
	)

	if data, err = os.ReadFile(path); err != nil {
		return "", err
	}

	return string(data), nil
}

func checkFileContentByFields(content string, fields []string) error {
	var (
		field         string
		scannedFields []string = []string{}
	)

	for _, field = range fields {

		for _, scannedField := range scannedFields {
			if scannedField == field {
				return ErrFieldsNotUnique
			}
		}

		if !strings.Contains(content, "{{."+field+"}}") {
			return ErrNotEnoughFields
		}

		scannedFields = append(scannedFields, field)
	}

	return nil
}
