// Package mail
/**
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
*/

package mail

import (
	"github.com/hramov/tg-bot-admin/internal/config"
	"net/smtp"
	"os"
)

func sendMail(receiver string, subject string, body string) error {
	host := config.ServerHostName
	port := config.ServerPort
	address := host + ":" + port
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	sub := "Subject: " + subject + "\n"
	message := []byte(sub + mime + "\n" + body)
	auth := smtp.PlainAuth("", os.Getenv("MAIL_ACCOUNT"), os.Getenv("MAIL_PASSWORD"), host)
	err := smtp.SendMail(address, auth, os.Getenv("MAIL_ACCOUNT"), []string{receiver}, message)
	return err
}
