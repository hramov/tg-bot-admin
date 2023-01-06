package types

import (
	"github.com/hramov/tg-bot-admin/src/interface/utils"
	"github.com/lib/pq"
)

var NullArray = []uint8{123, 78, 85, 76, 76, 125}

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
