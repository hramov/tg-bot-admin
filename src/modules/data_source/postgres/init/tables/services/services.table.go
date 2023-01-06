package services

import (
	"github.com/hramov/tg-bot-admin/src/modules/data_source/postgres/init/types"
)

var Services = types.Table{
	Name: "services",
	Fields: []types.Field{
		{
			Name:   "id",
			Type:   "serial",
			Unique: true,
		},
		{
			Name:   "title",
			Type:   "varchar",
			Unique: true,
		},
		{
			Name: "description",
			Type: "text",
		},
		{
			Name:       "auth_type",
			Type:       "integer",
			References: "nsi.auth_types(id)",
		},
		{
			Name:   "url",
			Type:   "text",
			Unique: true,
		},
		{
			Name:   "callback_url",
			Type:   "text",
			Unique: true,
		},
		{
			Name:   "client_id",
			Type:   "varchar",
			Unique: true,
		},
		{
			Name:   "client_secret",
			Type:   "varchar",
			Unique: true,
		},
		{
			Name:   "jwt_secret",
			Type:   "varchar",
			Unique: true,
		},
		{
			Name: "admin_email",
			Type: "varchar",
		},
		{
			Name:       "date_registered",
			Type:       "timestamp",
			DefaultVal: "now()",
		},
		{
			Name: "date_deleted",
			Type: "timestamp",
		},
	},
}
