package user

import (
	"context"
	"github.com/hramov/tg-bot-admin/internal/adapters/db"
	"github.com/hramov/tg-bot-admin/internal/domain/user"
	"github.com/hramov/tg-bot-admin/pkg/client/postgres"
	"github.com/jmoiron/sqlx"
)

type userStorage struct {
	db db.Connector
}

func NewStorage(db db.Connector) user.Storage {
	return &userStorage{db: db}
}

func (us *userStorage) GetById(ctx context.Context, id int) (*user.User, error) {
	conn, err := us.db.GetConn(ctx)
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

	res, err := postgres.ExecOne[user.User, UsersModel](ctx, conn, sql, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (us *userStorage) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	conn, err := us.db.GetConn(ctx)
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

	res, err := postgres.ExecOne[user.User, UsersModel](ctx, conn, sql, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (us *userStorage) Get(ctx context.Context, limit, offset int) ([]*user.User, error) {
	conn, err := us.db.GetConn(ctx)
	if err != nil {
		return nil, err
	}
	defer func(conn *sqlx.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	sql := `select * from users limit $1 offset $2`
	var params = []interface{}{limit, offset}

	res, err := postgres.Exec[user.User, UsersModel](ctx, conn, sql, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}
