package mail

import (
	template "github.com/hramov/tg-bot-admin/src/modules/templates"
)

func SendResetPasswordMail(email string, password string) error {
	t, err := template.GetTemplate(template.Mail, resetPasswordMailTemplate)
	if err != nil {
		return err
	}

	data := struct {
		Password string
	}{
		Password: password,
	}

	body, err := template.ParseTemplate(t, data)
	if err != nil {
		return err
	}

	return sendMail(email, "Сброс пароля", body)
}

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

func SendAddToServiceMail(email string, service string) error {
	t, err := template.GetTemplate(template.Mail, addToServiceMailTemplate)
	if err != nil {
		return err
	}

	data := struct {
		ServiceName string
	}{
		ServiceName: service,
	}

	body, err := template.ParseTemplate(t, data)
	if err != nil {
		return err
	}

	return sendMail(email, "Изменения параметров доступа", body)
}

func SendAdminRegisterMail(email string, password string) error {
	t, err := template.GetTemplate(template.Mail, adminRegisterMailTemplate)
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

	return sendMail(email, "Регистрация на сервере аутентификации", body)
}
