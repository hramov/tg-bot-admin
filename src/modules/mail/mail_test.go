package mail

import (
	serviceDto "github.com/hramov/tg-bot-admin/src/app/domains/service/dto"
	"github.com/hramov/tg-bot-admin/src/config"
	"os"
	"testing"
)

func TestSendAddToServiceMail(t *testing.T) {
	err := os.Setenv("MAIL_ACCOUNT", config.TestMailAccount)
	err = os.Setenv("MAIL_PASSWORD", config.TestMailPassword)
	err = SendAddToServiceMail("trykhramov@gmail.com", "GVC OAuth")
	if err != nil {
		t.Error(err)
	}
}

func TestSendAdminRegisterMail(t *testing.T) {
	err := os.Setenv("MAIL_ACCOUNT", config.TestMailAccount)
	err = os.Setenv("MAIL_PASSWORD", config.TestMailPassword)
	err = SendAdminRegisterMail("trykhramov@gmail.com", "admin")
	if err != nil {
		t.Error(err)
	}
}

func TestSendRegisterMail(t *testing.T) {
	err := os.Setenv("MAIL_ACCOUNT", config.TestMailAccount)
	err = os.Setenv("MAIL_PASSWORD", config.TestMailPassword)
	err = SendRegisterMail("trykhramov@gmail.com", "user")
	if err != nil {
		t.Error(err)
	}
}

func TestSendResetPasswordMail(t *testing.T) {
	err := os.Setenv("MAIL_ACCOUNT", config.TestMailAccount)
	err = os.Setenv("MAIL_PASSWORD", config.TestMailPassword)
	err = SendResetPasswordMail("trykhramov@gmail.com", "admin2")
	if err != nil {
		t.Error(err)
	}
}

func TestSendServiceMail(t *testing.T) {
	err := os.Setenv("MAIL_ACCOUNT", config.TestMailAccount)
	err = os.Setenv("MAIL_PASSWORD", config.TestMailPassword)

	data := &serviceDto.RegisterServiceResultDto{
		ClientId:     "client_id",
		ClientSecret: "client_secret",
		JwtSecret:    "jwt_secret",
	}

	err = SendServiceMail("trykhramov@gmail.com", "Дашборд", data)
	if err != nil {
		t.Error(err)
	}
}
