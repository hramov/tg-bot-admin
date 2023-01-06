package user

import (
	"context"
	appError "github.com/hramov/tg-bot-admin/src/interface/error"
)

var registry Registry

type Registry interface {
	GetUserById(ctx context.Context, id int) (*User, error)
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	GeoLabel string `json:"geo_label"`
	ChatId   string `json:"chat_id"`
	Password string `json:"password"`
}

func New(r Registry) {
	if registry == nil {
		registry = r
	}
}

func (u *User) Info(ctx context.Context) (*User, appError.IAppError) {
	user, err := registry.GetUserById(ctx, u.Id)
	if err != nil {
		return nil, appError.DatabaseError(err)
	}
	if user.Id == 0 {
		return nil, appError.AdminNotFound()
	}
	return user, nil
}
