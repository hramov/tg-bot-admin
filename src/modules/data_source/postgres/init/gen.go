package init

import (
	"fmt"
	"github.com/hramov/tg-bot-admin/src/modules/data_source/postgres/init/types"
)

func CreateTableSQL(schema string, table types.Table) (string, error) {
	sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.%s (", schema, table.Name)
	for _, field := range table.Fields {
		if field.Name != "" {
			sql += fmt.Sprintf("%s", field.Name)
		}
		if field.Type != "" {
			sql += fmt.Sprintf(" %s", field.Type)
		}
		if field.Unique {
			sql += fmt.Sprintf(" unique")
		}
		if field.DefaultVal != "" {
			sql += fmt.Sprintf(" default %s", field.DefaultVal)
		}
		if field.References != "" {
			sql += fmt.Sprintf(" references %s", field.References)
		}
		sql += fmt.Sprintf(", ")
	}
	sql = sql[:len(sql)-2]
	sql += ");"
	return sql, nil
}

func CreateSchemaSQL(name string) (string, error) {
	sql := fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s;", name)
	return sql, nil
}
