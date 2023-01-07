/**
 * Function InitDb prepares EMPTY database for work with server
 * SQL code is placed inside sql folder
 */

package init

import (
	"fmt"
	"github.com/hramov/tg-bot-admin/internal/adapters/db/tables"
	"github.com/hramov/tg-bot-admin/pkg/client/postgres/init/types"
	"github.com/jmoiron/sqlx"
)

func Start(db *sqlx.DB) error {
	for _, schema := range tables.TableConfigOrder {
		t, ok := tables.TableConfig[schema]
		if !ok {
			continue
		}
		err := createSchema(db, schema)
		if err != nil {
			return err
		}
		for _, table := range t {
			err := createTable(db, schema, table)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func createSchema(db *sqlx.DB, name string) error {
	sql, err := CreateSchemaSQL(name)
	if err != nil {
		return err
	}
	_, err = db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func createTable(db *sqlx.DB, schema string, table types.Table) error {
	sql, err := CreateTableSQL(schema, table)
	if err != nil {
		return err
	}

	_, err = db.Exec(sql)
	if err != nil {
		return err
	}

	if table.Default != nil {
		rows, err := db.Query(fmt.Sprintf("select id from %s.%s limit 1", schema, table.Name))
		if err != nil {
			return err
		}
		defer rows.Close()

		id := ""

		for rows.Next() {
			err := rows.Scan(&id)
			if err != nil {
				return err
			}
		}

		if id == "" {
			for _, query := range table.Default {
				_, err = db.Exec(query)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
