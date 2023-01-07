package postgres

import (
	"context"
	"testing"
)

const testPgDsn = "postgres://postgres:postgres@localhost:5432/coworking?sslmode=disable"

type TestStruct struct {
	Time string `json:"time" data_source:"time"`
}

func TestConnection(t *testing.T) {
	t.Setenv("PG_DSN", testPgDsn)
	db, _ := Connect()
	conn, err := db.getConnection(context.Background())
	if err != nil {
		t.Errorf("Connection error: %s", err.Error())
	}
	if conn == nil {
		t.Errorf("No conn, %v", conn)
	}
	defer conn.Close()

	var test int32
	err = conn.GetContext(context.Background(), &test, "select 1")
	if err != nil {
		t.Errorf("Query error: %s", err.Error())
	}
	if test != 1 {
		t.Errorf("Wrong result: %d", test)
	}
}
