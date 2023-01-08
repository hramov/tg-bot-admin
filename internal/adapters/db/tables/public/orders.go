package public

import "github.com/hramov/tg-bot-admin/pkg/db/postgres/init/types"

var OrdersTable = types.Table{
	Name: "orders",
	Fields: []types.Field{
		{
			Name:   "id",
			Type:   "serial",
			Unique: true,
		},
	},
	Default: []string{},
}
