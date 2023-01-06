package mail

import (
	"github.com/hramov/tg-bot-admin/src/config"
	"net/smtp"
	"os"
)

const (
	serviceRegisterMailTemplate = "service_register_mail.html"
	resetPasswordMailTemplate   = "reset_password_mail.html"
	userRegisterMailTemplate    = "register_mail.html"
	addToServiceMailTemplate    = "add_to_service_mail.html"
	adminRegisterMailTemplate   = "admin_register_mail.html"
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
