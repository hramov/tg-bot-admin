package migrations

import (
	"context"
	"github.com/hramov/tg-bot-admin/internal/adapters/db"
	"github.com/hramov/tg-bot-admin/pkg/logging"
)

var database db.Connector
var logger *logging.Logger

type migration interface {
	up() string
	down() string
}

var migrations = map[string]migration{
	"start": &startEmptyMigration{},
}

func Init(d db.Connector, l *logging.Logger) {
	database = d
	logger = l
}

func Start() error {
	ctx := context.TODO()
	conn, err := db.Unwrap(ctx, database)
	for title, migration := range migrations {
		logger.Infof("start migration: %s", title)
		sql := migration.up()
		_, err = conn.ExecContext(ctx, sql)
		if err != nil {
			sql = migration.down()
			_, err = conn.ExecContext(ctx, sql)
			if err != nil {
				logger.Errorf("migration %s down error %s", title, err.Error())
			}
			return err
		}
		logger.Infof("successfully migrate: %s", title)
	}
	return nil
}
