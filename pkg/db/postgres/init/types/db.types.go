package types

import (
	"encoding/json"
	"github.com/hramov/tg-bot-admin/pkg/utils"
	"github.com/lib/pq"
)

type Field struct {
	Name       string
	Type       string
	Unique     bool
	DefaultVal string
	References string
}

type Table struct {
	Name    string
	Fields  []Field
	Default []string
}

type Schema = map[string][]Table

type CreateResult struct {
	Id int `json:"id"`
}

var NullArray = []uint8{123, 78, 85, 76, 76, 125}
var NullObject = []uint8{123, 125}

type NullStringArray struct {
	String []string
	Valid  bool
}

func (n *NullStringArray) Scan(value interface{}) error {
	if value == nil || utils.EqualSlice(value.([]uint8), NullArray) {
		n.String, n.Valid = []string{}, false
		return nil
	}
	n.Valid = true
	return pq.Array(&n.String).Scan(value)
}

type NullSqlObject[T any] struct {
	Value *T
	Valid bool
}

func (n *NullSqlObject[T]) Scan(value interface{}) error {
	if value == nil || utils.EqualSlice(value.([]uint8), NullObject) {
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
