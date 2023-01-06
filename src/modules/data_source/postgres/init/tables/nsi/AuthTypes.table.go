package nsi

import (
	"github.com/hramov/tg-bot-admin/src/modules/data_source/postgres/init/types"
)

var AuthTypes = types.Table{
	Name: "auth_types",
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
			Name: "is_active",
			Type: "boolean",
		},
	},
	Default: []string{
		`insert into nsi.auth_types (title, is_active)
		 values (
			'Логин/Пароль',
			true
		 )
		`,
	},
}
