package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

func BeginTx(ctx context.Context, db Connector) (*sqlx.Conn, *sqlx.Tx, error) {
	conn, connErr := db.GetConn(ctx)
	if connErr != nil {
		return nil, nil, connErr
	}
	tx, err := conn.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	})
	return conn, tx, err
}

func CommitTx(ctx context.Context, db Connector, conn *sqlx.Conn, tx *sqlx.Tx) error {
	defer db.ReturnConn(ctx, conn)
	err := tx.Commit()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	return nil
}

func RollbackTx(ctx context.Context, db Connector, conn *sqlx.Conn, tx *sqlx.Tx) error {
	defer db.ReturnConn(ctx, conn)
	return tx.Rollback()
}

// ExecTx TODO think about connection closing
func ExecTx[T any, V Mapper[T]](ctx context.Context, tx *sqlx.Tx, s string, params []interface{}) ([]*T, error) {
	valid := validateFilters(ctx, new(T))
	if !valid {
		return nil, fmt.Errorf("filters not valid")
	}
	s, filterParams, err := formatSqlFilters(s, "users", 1, ctx)
	for _, v := range filterParams {
		params = append(params, v)
	}

	var model []V
	err = tx.SelectContext(ctx, &model, s, params...)
	if err != nil {
		rErr := tx.Rollback()
		if rErr != nil {
			log.Println(rErr.Error())
		}
		return nil, err
	}
	var dto []*T
	for _, v := range model {
		d := v.Map()
		dto = append(dto, &d)
	}
	return dto, nil
}

func ExecOneTx[T any, V Mapper[T]](ctx context.Context, tx *sqlx.Tx, s string, params []interface{}) (*T, error) {
	var model []V
	err := tx.SelectContext(ctx, &model, s, params...)
	if err != nil {
		return nil, err
	}
	if len(model) == 0 {
		return new(T), nil
	}
	dto := model[0].Map()
	return &dto, nil
}

func Exec[T any, V Mapper[T]](ctx context.Context, db Connector, s string, params []interface{}) ([]*T, error) {
	valid := validateFilters(ctx, new(T))
	if !valid {
		return nil, fmt.Errorf("filters not valid")
	}
	s, filterParams, err := formatSqlFilters(s, "users", 1, ctx)
	for _, v := range filterParams {
		params = append(params, v)
	}

	conn, err := db.GetConn(ctx)
	if err != nil {
		return nil, err
	}
	defer db.ReturnConn(ctx, conn)
	var model []V
	err = conn.SelectContext(ctx, &model, s, params...)
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

func ExecOne[T any, V Mapper[T]](ctx context.Context, db Connector, s string, params []interface{}) (*T, error) {
	conn, err := db.GetConn(ctx)
	if err != nil {
		return nil, err
	}
	defer db.ReturnConn(ctx, conn)
	var model []V
	err = conn.SelectContext(ctx, &model, s, params...)
	if err != nil {
		return nil, err
	}
	if len(model) == 0 {
		return new(T), nil
	}
	dto := model[0].Map()
	return &dto, nil
}
