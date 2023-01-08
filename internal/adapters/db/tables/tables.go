package tables

import "github.com/hramov/tg-bot-admin/pkg/db/postgres/init/types"

var TableConfig = types.Schema{
	"public": []types.Table{},
}

var TableConfigOrder = []string{"public"}
