package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Mapper[T any] interface {
	Map() T
}

func Exec[T any, V Mapper[T]](ctx context.Context, conn *sqlx.Conn, sql string, params []interface{}) (*[]T, error) {
	var model []V

	err := conn.SelectContext(ctx, &model, sql, params...)
	if err != nil {
		return nil, err
	}

	var dto []T
	for _, v := range model {
		d := v.Map()
		dto = append(dto, d)
	}

	return &dto, nil
}

func ExecOne[T any, V Mapper[T]](ctx context.Context, conn *sqlx.Conn, sql string, params []interface{}) (*T, error) {
	var model []V

	err := conn.SelectContext(ctx, &model, sql, params...)
	if err != nil {
		return nil, err
	}

	if len(model) == 0 {
		return new(T), nil
	}

	dto := model[0].Map()
	return &dto, nil
}
