package types

import "encoding/json"

// NullObjectArray TODO implement
var NullObjectArray = []uint8{123, 125}

type NullSqlObjectArray[T any] struct {
	Value *T
	Valid bool
}

func (n *NullSqlObjectArray[T]) Scan(value interface{}) error {
	if value == nil || EqualSlice(value.([]uint8), NullObjectArray) {
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
