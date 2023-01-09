package db

import (
	"context"
)

func Exec[T any, V Mapper[T]](ctx context.Context, db Connector, sql string, params []interface{}) ([]*T, error) {
	conn, err := db.GetConn(ctx)
	if err != nil {
		return nil, err
	}
	defer db.ReturnConn(ctx, conn)

	var model []V
	err = conn.SelectContext(ctx, &model, sql, params...)
	if err != nil {
		return nil, err
	}
	var dto []*T
	for _, v := range model {
		d := v.Map()
		dto = append(dto, &d)
	}
	return dto, nil
}

func ExecOne[T any, V Mapper[T]](ctx context.Context, db Connector, sql string, params []interface{}) (*T, error) {
	conn, err := db.GetConn(ctx)
	if err != nil {
		return nil, err
	}
	defer db.ReturnConn(ctx, conn)

	var model []V
	err = conn.SelectContext(ctx, &model, sql, params...)
	if err != nil {
		return nil, err
	}
	if len(model) == 0 {
		return new(T), nil
	}
	dto := model[0].Map()
	return &dto, nil
}
