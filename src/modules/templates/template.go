package template

import (
	"bytes"
	"html/template"
	"path/filepath"
	"runtime"
)

const (
	Http = "http"
	Mail = "mail"
)

func GetTemplate(templateType string, templateName string) (*template.Template, error) {
	var filePath string
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	switch templateType {
	case Http:
		filePath = filepath.Join(basePath, Http, templateName)
		break
	case Mail:
		filePath = filepath.Join(basePath, Mail, templateName)
		break
	}

	t, err := template.ParseFiles(filePath)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func ParseTemplate(t *template.Template, data interface{}) (string, error) {
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		return "", err
	}
	body := buf.String()
	return body, nil
}
