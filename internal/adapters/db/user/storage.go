package user

import (
	"context"
	"fmt"
	"github.com/hramov/tg-bot-admin/internal/adapters/db"
	"github.com/hramov/tg-bot-admin/internal/domain/user"
	"github.com/hramov/tg-bot-admin/pkg/logging"
)

type storage struct {
	db     db.Connector
	logger *logging.Logger
}

func NewStorage(logger *logging.Logger, db db.Connector) user.Storage {
	return &storage{db: db, logger: logger}
}

func (s *storage) GetBy(ctx context.Context, field string, param any) (*user.User, error) {
	sql := fmt.Sprintf("select u.*, r.permissions from users u join roles r on u.role = r.id where u.%s = $1", field)
	var params = []interface{}{param}
	res, err := db.ExecOne[user.User, Model](ctx, s.db, sql, params)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *storage) Get(ctx context.Context) ([]*user.User, error) {
	sql := `select users.*, roles.permissions from users join roles on users.role = roles.id`
	res, err := db.Exec[user.User, Model](ctx, s.db, sql, nil)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return res, nil
}

func (s *storage) Create(ctx context.Context, dto *user.CreateDto) (*int, error) {
	sql := `insert into users (name, phone, address, geo_label, chat_id, email, password, registered_at) values ($1, $2, $3, $4, $5, $6, $7, now()) returning id`
	var params = []interface{}{dto.Name, dto.Phone, dto.Address, dto.GeoLabel, dto.ChatId, dto.Email, dto.Password}
	res, err := db.ExecOne[user.User, Model](ctx, s.db, sql, params)
	if err != nil {
		return nil, err
	}
	return &res.Id, nil
}

func (s *storage) Update(ctx context.Context, dto *user.UpdateDto) (*int, error) {
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
	res, err := db.ExecOne[user.User, Model](ctx, s.db, sql, params)
	if err != nil {
		return nil, err
	}
	return &res.Id, nil
}

func (s *storage) Delete(ctx context.Context, id int) (*int, error) {
	sql := `delete from users where id = $1 returning id`
	var params = []interface{}{id}
	res, err := db.ExecOne[user.User, Model](ctx, s.db, sql, params)
	if err != nil {
		return nil, err
	}
	return &res.Id, nil
}
