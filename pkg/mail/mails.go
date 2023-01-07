package mail

import (
	template "github.com/hramov/tg-bot-admin/pkg/templates"
)

const (
	userRegisterMailTemplate = ""
)

func SendRegisterMail(email string, password string) error {
	t, err := template.GetTemplate(template.Mail, userRegisterMailTemplate)
	if err != nil {
		return err
	}

	data := struct {
		Email    string
		Password string
	}{
		Email:    email,
		Password: password,
	}

	body, err := template.ParseTemplate(t, data)
	if err != nil {
		return err
	}

	return sendMail(email, "Регистрация в системе аутентификации", body)
}
