package registry

import (
	"context"
	"github.com/hramov/tg-bot-admin/src/domain/user"
	"github.com/hramov/tg-bot-admin/src/modules/data_source"
	"github.com/hramov/tg-bot-admin/src/modules/data_source/postgres"
	"github.com/hramov/tg-bot-admin/src/modules/data_source/postgres/models"
	"github.com/jmoiron/sqlx"
)

type UserRegistry struct {
	dataSource data_source.DataSource
}

func NewUserRegistry(dataSource data_source.DataSource) UserRegistry {
	return UserRegistry{dataSource: dataSource}
}

func (ur UserRegistry) GetUserById(ctx context.Context, id int) (*user.User, error) {
	conn, err := ur.dataSource.GetConn(ctx)
	if err != nil {
		return nil, err
	}
	defer func(conn *sqlx.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	sql := `select * from users where id = $1`
	var params = []interface{}{id}

	res, err := postgres.ExecOne[user.User, models.UsersModel](ctx, conn, sql, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}
