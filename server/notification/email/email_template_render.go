package email

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"
)

//go:embed templates/*
var emailTemplates embed.FS

func RanderEmailAuthTemplate(templateName string, data any) (string, error) {
	fmt.Println("templateName: ", templateName)
	tmpl, err := template.ParseFS(emailTemplates, "templates/"+templateName)
	if err != nil {
		return "", fmt.Errorf("parsing template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("executing template: %w", err)
	}
	return buf.String(), nil
}
func RanderWelcomeTemplate(templateName string, data any) (string, error) {
	tmpl, err := template.ParseFS(emailTemplates, "templates/"+templateName)
	if err != nil {
		return "", fmt.Errorf("parsing template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("executing template: %w", err)
	}
	return buf.String(), nil
}
