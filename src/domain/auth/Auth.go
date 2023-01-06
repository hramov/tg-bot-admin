package auth

import (
	"context"
	appError "github.com/hramov/tg-bot-admin/src/interface/error"
	"github.com/hramov/tg-bot-admin/src/modules/jwt"
	"github.com/hramov/tg-bot-admin/src/modules/logger"
	"os"
)

var registry Registry

type Login struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Refresh struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Registry interface {
	GetCandidate(ctx context.Context, email string) (*Login, error)
}

func New(r Registry) {
	if registry == nil {
		registry = r
	}
}

func (l *Login) validate() appError.IAppError {
	if l.Email == "" || l.Password == "" {
		return appError.NoCredentialsError()
	}
	return nil
}

func (l *Login) Login(ctx context.Context) (string, string, appError.IAppError) {
	vErr := l.validate()
	if vErr != nil {
		return "", "", vErr
	}

	candidate, err := registry.GetCandidate(ctx, l.Email)
	if err != nil {
		logger.Instance.Error(err.Error())
		return "", "", appError.DatabaseError(err)
	}

	if candidate.Id == 0 {
		return "", "", appError.LoginOrPasswordIncorrectError()
	}
	valid := jwt.CheckPassword(l.Password, candidate.Password)
	if !valid {
		return "", "", appError.LoginOrPasswordIncorrectError()
	}

	at, rt, err := jwt.CreateToken(candidate.Id, os.Getenv("JWT_SECRET"))
	if err != nil {
		return "", "", appError.LoginOrPasswordIncorrectError()
	}

	return at, rt, nil
}

func (r *Refresh) Refresh(ctx context.Context) (string, string, appError.IAppError) {
	return "", "", nil
}
