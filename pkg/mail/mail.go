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
)

var Instance *Mail

type Mail struct {
	cfg config.MailConfig
}

func New(cfg config.MailConfig) {
	if Instance == nil {
		Instance = &Mail{cfg: cfg}
	}
}

func (m *Mail) sendMail(receiver string, subject string, body string) error {
	host := m.cfg.ServerHostName
	port := m.cfg.ServerPort
	address := host + ":" + port
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	sub := "Subject: " + subject + "\n"
	message := []byte(sub + mime + "\n" + body)
	auth := smtp.PlainAuth("", m.cfg.Account, m.cfg.Password, host)
	err := smtp.SendMail(address, auth, m.cfg.Account, []string{receiver}, message)
	return err
}
