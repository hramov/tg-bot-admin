package authTest

import (
	"context"
	"github.com/hramov/tg-bot-admin/src/domain/auth"
	appError "github.com/hramov/tg-bot-admin/src/interface/error"
	"github.com/hramov/tg-bot-admin/src/modules/logger"
	"testing"
)

func TestLogin_NoCredentials(t *testing.T) {
	auth.New(NewAuthRegistryMock())

	login := &auth.Login{}

	_, _, err := login.Login(context.Background())
	if err == nil {
		t.Error("should be an error")
	}

	if err.Error() != appError.NoCredentialsError().Error() {
		t.Errorf("error mismatch: %v", err.Error())
	}
}

func TestLogin_WrongEmail(t *testing.T) {
	auth.New(NewAuthRegistryMock())

	login := &auth.Login{
		Email:    "wrongEmail",
		Password: "admin",
	}

	_, _, err := login.Login(context.Background())
	if err == nil {
		t.Error("should be an error")
	}

	if err.Error() != appError.LoginOrPasswordIncorrectError().Error() {
		t.Errorf("error mismatch: %v", err.Error())
	}
}

func TestLogin_WrongPassword(t *testing.T) {
	auth.New(NewAuthRegistryMock())

	login := &auth.Login{
		Email:    "admin@admin.ru",
		Password: "admin1",
	}

	_, _, err := login.Login(context.Background())
	if err == nil {
		t.Error("should be an error")
	}

	if err.Error() != appError.LoginOrPasswordIncorrectError().Error() {
		t.Errorf("error mismatch: %v", err.Error())
	}
}

func TestLogin_DatabaseError(t *testing.T) {
	err := logger.NewTest()
	if err != nil {
		t.Error(err.Error())
	}
	auth.New(NewAuthRegistryMock())
	login := &auth.Login{
		Email:    "admin@admin.com",
		Password: "admin1",
	}
	_, _, err = login.Login(context.Background())
	if err == nil {
		t.Error("should be an error")
	}
}

func TestLogin_Success(t *testing.T) {
	auth.New(NewAuthRegistryMock())

	login := &auth.Login{
		Email:    "admin@admin.ru",
		Password: "admin",
	}

	_, _, err := login.Login(context.Background())
	if err != nil {
		t.Error(err.Error())
	}
}
