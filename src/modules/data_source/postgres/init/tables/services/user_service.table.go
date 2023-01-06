package services

import (
	"github.com/hramov/tg-bot-admin/src/modules/data_source/postgres/init/types"
)

var UserService = types.Table{
	Name: "user_service",
	Fields: []types.Field{
		{
			Name:   "id",
			Type:   "serial",
			Unique: true,
		},
		{
			Name:       "user_id",
			Type:       "uuid",
			References: "users.users(id)",
		},
		{
			Name:       "service_id",
			Type:       "integer",
			References: "services.services(id)",
		},
		{
			Name:       "date_registered",
			Type:       "timestamp",
			DefaultVal: "now()",
		},
	},
}
