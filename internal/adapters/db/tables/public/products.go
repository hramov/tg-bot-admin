package public

import "github.com/hramov/tg-bot-admin/pkg/db/postgres/init/types"

var ProductsTable = types.Table{
	Name: "products",
	Fields: []types.Field{
		{
			Name:   "id",
			Type:   "serial",
			Unique: true,
		},
	},
	Default: []string{},
}
