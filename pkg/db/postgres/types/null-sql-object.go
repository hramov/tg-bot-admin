package types

import (
	"encoding/json"
)

type NullSqlObject[T any] struct {
	Value T
	Valid bool
}

func (n *NullSqlObject[T]) Scan(value interface{}) error {
	n.Valid = true
	err := json.Unmarshal(value.([]uint8), &n.Value)
	if err != nil {
		return err
	}
	return nil
}
