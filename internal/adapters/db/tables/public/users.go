package public

import "github.com/hramov/tg-bot-admin/pkg/client/postgres/init/types"

var UsersTable = types.Table{
	Name: "users",
	Fields: []types.Field{
		{
			Name:   "id",
			Type:   "serial",
			Unique: true,
		},
	},
	Default: []string{},
}
