package registry

import (
	"context"
	"github.com/hramov/tg-bot-admin/src/domain/auth"
	"github.com/hramov/tg-bot-admin/src/domain/user"
	"github.com/hramov/tg-bot-admin/src/modules/data_source"
	"github.com/hramov/tg-bot-admin/src/modules/data_source/postgres"
	"github.com/hramov/tg-bot-admin/src/modules/data_source/postgres/models"
	"github.com/jmoiron/sqlx"
)

type AuthRegistry struct {
	dataSource data_source.DataSource
}

func NewAuthRegistry(dataSource data_source.DataSource) AuthRegistry {
	return AuthRegistry{dataSource: dataSource}
}

func (ar AuthRegistry) GetCandidate(ctx context.Context, email string) (*auth.Login, error) {
	conn, err := ar.dataSource.GetConn(ctx)
	if err != nil {
		return nil, err
	}
	defer func(conn *sqlx.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	sql := `select id, email, password from users where email = $1`
	var params = []interface{}{email}

	res, err := postgres.ExecOne[user.User, models.UsersModel](ctx, conn, sql, params)
	if err != nil {
		return nil, err
	}

	return &auth.Login{
		Id:       res.Id,
		Email:    res.Email,
		Password: res.Password,
	}, nil
}
