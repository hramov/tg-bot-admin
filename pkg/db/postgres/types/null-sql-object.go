package types

import (
	"encoding/json"
)

var NullObject = []uint8{123, 125}

type NullSqlObject[T any] struct {
	Value *T
	Valid bool
}

func (n *NullSqlObject[T]) Scan(value interface{}) error {
	if value == nil || EqualSlice(value.([]uint8), NullObject) {
		n.Value, n.Valid = nil, false
		return nil
	}
	n.Valid = true
	err := json.Unmarshal(value.([]uint8), &n.Value)
	if err != nil {
		return err
	}
	return nil
}
