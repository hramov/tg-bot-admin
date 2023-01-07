package mail

import (
	"github.com/hramov/tg-bot-admin/internal/config"
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
