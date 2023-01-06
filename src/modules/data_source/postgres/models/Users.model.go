package models

import (
	"database/sql"
	"github.com/hramov/tg-bot-admin/src/domain/user"
)

type UsersModel struct {
	Id           int
	Name         string
	Phone        sql.NullString
	Address      sql.NullString
	GeoLabel     sql.NullString `db:"geo_label"`
	ChatId       sql.NullString `db:"chat_id"`
	Email        string
	Password     string
	RegisteredAt sql.NullTime `db:"registered_at"`
	LastLogin    sql.NullTime `db:"last_login"`
}

func (um UsersModel) Map() user.User {
	return user.User{
		Id:       um.Id,
		Name:     um.Name,
		Phone:    um.Phone.String,
		Address:  um.Address.String,
		Email:    um.Email,
		Password: um.Password,
		GeoLabel: um.GeoLabel.String,
		ChatId:   um.ChatId.String,
	}
}
