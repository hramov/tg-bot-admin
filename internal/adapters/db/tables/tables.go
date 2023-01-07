package tables

import "github.com/hramov/tg-bot-admin/pkg/client/postgres/init/types"

var TableConfig = types.Schema{
	"public": []types.Table{},
}

var TableConfigOrder = []string{"public"}
