package nsi

import (
	"github.com/hramov/tg-bot-admin/src/modules/data_source/postgres/init/types"
)

var Cts = types.Table{
	Name: "cts",
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
