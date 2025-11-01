package mailer

import (
	"bytes"
	"embed"
	"html/template"
)

const (
	FromName            = "GopherSocial"
	maxRetries          = 3
	UserWelcomeTemplate = "user_invitation.tmpl"
)

//go:embed "templates"
var FS embed.FS

type Client interface {
	Send(templateFIle, username, email string, data any, isSandbox bool) (int, error)
}

func buildTemplate(templateFile string, data any) (string, string, error) {
	// Template parsing and building
	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return "", "", err
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return "", "", err
	}

	body := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(body, "body", data)
	if err != nil {
		return "", "", err
	}

	return subject.String(), body.String(), nil
}
