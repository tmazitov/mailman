# Mailman

Mailman is a lightweight Go library designed to simplify email handling in your applications.

## Features

- Easy-to-use API for sending emails.
- Supports multiple email providers.
- Lightweight and dependency-free.

## Installation

```bash
go get github.com/tmazitov/mailman
```

## Usage

1. Create an Mailman instance :
```go
	mail, err := mailman.NewMailMan(&mailman.MailmanConfig{
		Email: "example@mail.com",
		Pass:  "password",
		Port:  587,
		Host:  "smtp.gmail.com",
	})
```

2. Start it :
```go
	mail.Start()
```

3. Setup message templates :
```go
    mail.SendMessage(&mailman.MessageInfo{
		Subject:      "Test",
		DistEmail:    "example@mail.com",
        Content:      "Hello World!"
	})
```

4. Do not forget to stop it :
```go
    defer mail.Stop()
```

## Templating

1. Setup templates :

```go
err = m.SetupMessageTemplates([]*template.MessageTemplateInfo{
    {
        Name:     "test",               // Template name
        FilePath: "template.html",      // Path to html file
        Fields:   []string{"example"},  // Fields that you have inside the file with this format {{.example}}
    },
})
```

2. Send message using template ( Content field will be ignored ) :
```go
    mail.SendMessage(&mailman.MessageInfo{
        TemplateName: "test",
		Subject:      "Test",
		DistEmail:    "example@mail.com",
        Content:      "Hello World!"
        FieldValues:  map[string]string{
			"example": "exampleValue",
		}
	})
```


## License

This project is licensed under the MIT License.