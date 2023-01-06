package users

import (
	"fmt"
	"github.com/hramov/tg-bot-admin/src/modules/data_source/postgres/init/types"
	"github.com/hramov/tg-bot-admin/src/modules/jwt"
)

func getPass(pass string) string {
	hashed, _ := jwt.CreateHashedPassword(pass)
	return hashed
}

func getSql() string {
	return fmt.Sprintf(`insert into users.admins (id, fio, ad, email, ip, username, password, date_registered)
		 values (
			(select gen_random_uuid()),
			'Главный администратор',
			'khramovsi@gvc.oao.rzd',
			'khramovsi@gvc.rzd',
			'10.200.20.115',
			'admin',
			'%s',
			now()
		 )`, getPass("admin"))
}

var Admins = types.Table{
	Name: "admins",
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
			Name:   "email",
			Type:   "varchar",
			Unique: true,
		},
		{
			Name:   "ip",
			Type:   "varchar",
			Unique: true,
		},
		{
			Name:   "ad",
			Type:   "varchar",
			Unique: true,
		},
		{
			Name:   "username",
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
	},
	Default: []string{
		getSql(),
	},
}
