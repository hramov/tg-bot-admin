package order

import (
	"context"
	"fmt"

	"github.com/hramov/tg-bot-admin/internal/adapters/db"
	"github.com/hramov/tg-bot-admin/internal/domain/order"
	"github.com/hramov/tg-bot-admin/pkg/logging"
)

type storage struct {
	db     db.Connector
	logger *logging.Logger
}

func NewStorage(logger *logging.Logger, db db.Connector) order.Storage {
	return &storage{db: db, logger: logger}
}

func (s storage) Get(ctx context.Context, limit int, lastId int) ([]*order.Order, error) {
	sql := `select orders.* from orders where id > $1 order by id desc limit $2`
	params := []interface{}{lastId, limit}
	res, err := db.Exec[order.Order, Model](ctx, s.db, sql, params)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s storage) GetBy(ctx context.Context, field string, param any) (*order.Order, error) {
	sql := fmt.Sprintf("select * from orders where %s = $1", field)
	var params = []interface{}{param}
	res, err := db.ExecOne[order.Order, Model](ctx, s.db, sql, params)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s storage) Create(ctx context.Context, dto order.Order) (*int, error) {
	sql := `
		insert into orders 
			(user_info, products, total_price, status, created_at, updated_at, sent_at, done_at, deleted_at) 
		values 
			($1, $2, $3, $4, now(), now(), null, null, null)
		returning id`

	var params = []interface{}{dto.UserInfo, dto.Products, dto.TotalPrice, dto.Status}
	res, err := db.ExecOne[order.Order, Model](ctx, s.db, sql, params)
	if err != nil {
		return nil, err
	}
	return &res.Id, nil
}

func (s storage) ChangeStatus(ctx context.Context, orderId, statusId int) (*int, error) {
	sql := `
		update orders 
			set status = $1 
		where id = $2
		returning id;
	`
	var params = []interface{}{orderId, statusId}
	res, err := db.ExecOne[order.Order, Model](ctx, s.db, sql, params)
	if err != nil {
		return nil, err
	}
	return &res.Id, nil
}

func (s storage) GetUserOrders(ctx context.Context, userId int) ([]*order.InfoForUser, error) {
	sql := "select * from orders where user_id = $1"
	var params = []interface{}{userId}
	res, err := db.Exec[order.Order, Model](ctx, s.db, sql, params)
	if err != nil {
		return nil, err
	}
	var result []*order.InfoForUser

	for _, ord := range res {
		result = append(result, &order.InfoForUser{
			OrderId:    ord.Id,
			Products:   []order.ShallowWeedProduct{},
			TotalPrice: int(ord.TotalPrice),
		})
	}
	return result, nil
}
