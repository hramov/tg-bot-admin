package user

import (
	"database/sql"
	"github.com/hramov/tg-bot-admin/internal/domain/user"
	"github.com/hramov/tg-bot-admin/pkg/db/postgres/types"
	"github.com/hramov/tg-bot-admin/pkg/jwt"
)

type Model struct {
	Id           int
	Role         int
	Permissions  types.NullSqlObject[jwt.Permissions] `db:"permissions"`
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

func (um Model) Map() user.User {
	return user.User{
		Id:           um.Id,
		Role:         um.Role,
		Permissions:  um.Permissions.Value,
		Name:         um.Name,
		Phone:        um.Phone.String,
		Address:      um.Address.String,
		Email:        um.Email,
		Password:     um.Password,
		GeoLabel:     um.GeoLabel.String,
		ChatId:       um.ChatId.String,
		RegisteredAt: um.RegisteredAt.Time,
		LastLogin:    um.LastLogin.Time,
	}
}
