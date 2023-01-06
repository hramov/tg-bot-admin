package users

import (
	"github.com/hramov/tg-bot-admin/src/modules/data_source/postgres/init/types"
)

var Users = types.Table{
	Name: "users",
	Fields: []types.Field{
		{
			Name:   "id",
			Type:   "uuid",
			Unique: true,
		},
		{
			Name: "fio",
			Type: "varchar",
		},
		{
			Name:   "ad",
			Type:   "varchar",
			Unique: true,
		},
		{
			Name:   "ip",
			Type:   "varchar",
			Unique: true,
		},
		{
			Name: "password",
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
		{
			Name:       "ivc",
			Type:       "integer",
			References: "nsi.ivc(id)",
		},
	},
}
