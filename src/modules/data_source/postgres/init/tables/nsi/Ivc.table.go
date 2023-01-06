package nsi

import (
	"github.com/hramov/tg-bot-admin/src/modules/data_source/postgres/init/types"
)

var Ivc = types.Table{
	Name: "ivc",
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
	},
}
