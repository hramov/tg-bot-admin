package init

import (
	"github.com/hramov/tg-bot-admin/src/modules/data_source/postgres/init/tables/nsi"
	"github.com/hramov/tg-bot-admin/src/modules/data_source/postgres/init/tables/services"
	"github.com/hramov/tg-bot-admin/src/modules/data_source/postgres/init/tables/users"
	"github.com/hramov/tg-bot-admin/src/modules/data_source/postgres/init/types"
)

var TableConfig = types.Schema{
	"nsi": []types.Table{
		nsi.Ivc,
		nsi.Cts,
		nsi.AuthTypes,
	},
	"users": []types.Table{
		users.Users,
		users.Admins,
	},
	"services": []types.Table{
		services.Services,
		services.UserService,
	},
}

var TableConfigOrder = []string{"nsi", "users", "services"}
