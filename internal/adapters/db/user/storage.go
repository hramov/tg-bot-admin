package user

import (
	"context"
	"fmt"
	"github.com/hramov/tg-bot-admin/internal/adapters/db"
	"github.com/hramov/tg-bot-admin/internal/domain/user"
	"github.com/hramov/tg-bot-admin/pkg/db/postgres"
	"github.com/hramov/tg-bot-admin/pkg/logging"
)

type userStorage struct {
	db     db.Connector
	logger *logging.Logger
}

func NewStorage(logger *logging.Logger, db db.Connector) user.IStorage {
	return &userStorage{db: db, logger: logger}
}

func (us *userStorage) GetBy(ctx context.Context, field string, param any) (*user.User, error) {
	conn, err := us.db.GetConn(ctx)
	if err != nil {
		return nil, err
	}
	defer us.db.ReturnConn(ctx, conn)

	sql := fmt.Sprintf("select u.*, r.permissions from users u join roles r on u.role = r.id where u.%s = $1", field)
	var params = []interface{}{param}
	res, err := postgres.ExecOne[user.User, Model](ctx, conn, sql, params)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (us *userStorage) GetById(ctx context.Context, id int) (*user.User, error) {
	conn, err := us.db.GetConn(ctx)
	if err != nil {
		return nil, err
	}
	defer us.db.ReturnConn(ctx, conn)

	sql := `select * from users where id = $1`
	var params = []interface{}{id}

	res, err := postgres.ExecOne[user.User, Model](ctx, conn, sql, params)
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
	defer us.db.ReturnConn(ctx, conn)

	sql := `select id, email, password from users where email = $1`
	var params = []interface{}{email}

	res, err := postgres.ExecOne[user.User, Model](ctx, conn, sql, params)
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
	defer us.db.ReturnConn(ctx, conn)

	sql := `select * from users limit $1 offset $2`
	var params = []interface{}{limit, offset}

	res, err := postgres.Exec[user.User, Model](ctx, conn, sql, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (us *userStorage) Create(ctx context.Context, dto *user.CreateDto) (*int, error) {
	conn, err := us.db.GetConn(ctx)
	if err != nil {
		return nil, err
	}
	defer us.db.ReturnConn(ctx, conn)

	sql := `insert into users (name, phone, address, geo_label, chat_id, email, password, registered_at) values ($1, $2, $3, $4, $5, $6, $7, now()) returning id`
	var params = []interface{}{dto.Name, dto.Phone, dto.Address, dto.GeoLabel, dto.ChatId, dto.Email, dto.Password}

	res, err := postgres.ExecOne[user.User, Model](ctx, conn, sql, params)
	if err != nil {
		return nil, err
	}

	return &res.Id, nil
}

func (us *userStorage) Update(ctx context.Context, dto *user.UpdateDto) (*int, error) {
	conn, err := us.db.GetConn(ctx)
	defer us.db.ReturnConn(ctx, conn)

	sql := `
		update users set
			name = $1,
			phone = $2,
			address = $3,
			geo_label = $4,
			chat_id = $5,
			email = $6
			password = &7
		where id = $8
		returning id
	`
	var params = []interface{}{dto.Name, dto.Phone, dto.Address, dto.GeoLabel, dto.ChatId, dto.Email, dto.Password, dto.Id}

	res, err := postgres.ExecOne[user.User, Model](ctx, conn, sql, params)
	if err != nil {
		return nil, err
	}

	return &res.Id, nil
}

func (us *userStorage) Delete(ctx context.Context, id int) (*int, error) {
	conn, err := us.db.GetConn(ctx)
	if err != nil {
		return nil, err
	}
	defer us.db.ReturnConn(ctx, conn)

	sql := `delete from users where id = $1 returning id`
	var params = []interface{}{id}

	res, err := postgres.ExecOne[user.User, Model](ctx, conn, sql, params)
	if err != nil {
		return nil, err
	}

	return &res.Id, nil
}
