package models

import (
	"database/sql"
	"github.com/hramov/tg-bot-admin/src/domain/user"
)

type UsersModel struct {
	Id           int
	Name         string
	Phone        string
	Address      string
	GeoLabel     string `db:"geo_label"`
	ChatId       string `db:"chat_id"`
	Email        string
	Password     string
	RegisteredAt sql.NullTime `db:"registered_at"`
	LastLogin    sql.NullTime `db:"last_login"`
}

func (um UsersModel) Map() user.User {
	return user.User{
		Id:       um.Id,
		Name:     um.Name,
		Phone:    um.Phone,
		Address:  um.Address,
		Email:    um.Email,
		Password: um.Password,
		GeoLabel: um.GeoLabel,
		ChatId:   um.ChatId,
	}
}
