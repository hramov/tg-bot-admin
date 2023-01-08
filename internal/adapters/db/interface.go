package db

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Mapper[T any] interface {
	Map() T
}

type Connector interface {
	GetConn(ctx context.Context) (*sqlx.Conn, error)
	ReturnConn(ctx context.Context, conn *sqlx.Conn)
}
