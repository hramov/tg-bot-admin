package public

import "github.com/hramov/tg-bot-admin/pkg/db/postgres/init/types"

var UsersTable = types.Table{
	Name: "users",
	Fields: []types.Field{
		{
			Name:   "id",
			Type:   "serial",
			Unique: true,
		},
		{
			Name:       "role",
			Type:       "integer",
			References: "roles(id)",
		},
		{
			Name: "name",
			Type: "varchar",
		},
		{
			Name:   "phone",
			Type:   "varchar",
			Unique: true,
		},
		{
			Name: "address",
			Type: "varchar",
		},
		{
			Name: "geo_label",
			Type: "varchar",
		},
		{
			Name:   "chat_id",
			Type:   "varchar",
			Unique: true,
		},
		{
			Name:   "email",
			Type:   "varchar",
			Unique: true,
		},
		{
			Name: "password",
			Type: "varchar",
		},
		{
			Name:       "registered_at",
			Type:       "timestamp",
			DefaultVal: "now()",
		},
		{
			Name: "last_login",
			Type: "timestamp",
		},
	},
	Default: []string{
		`insert into users (role, name, phone, address, geo_label, chat_id, email, password)
		values (1, 'Sergey', '000-000-00-00', '', '', '', 'trykhramov@gmail.com', 'admin')`,
	},
}
