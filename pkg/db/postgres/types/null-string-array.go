package types

import (
	"github.com/lib/pq"
)

var NullArray = []uint8{123, 78, 85, 76, 76, 125}

type NullStringArray struct {
	String []string
	Valid  bool
}

func (n *NullStringArray) Scan(value interface{}) error {
	if value == nil || EqualSlice(value.([]uint8), NullArray) {
		n.String, n.Valid = []string{}, false
		return nil
	}
	n.Valid = true
	return pq.Array(&n.String).Scan(value)
}
