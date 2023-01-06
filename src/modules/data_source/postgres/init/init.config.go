package init

import (
	"github.com/hramov/tg-bot-admin/src/modules/data_source/postgres/init/types"
)

var TableConfig = types.Schema{
	"public": []types.Table{},
}

var TableConfigOrder = []string{"public"}
