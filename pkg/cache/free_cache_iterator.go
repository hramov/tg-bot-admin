package cache

import (
	"github.com/coocood/freecache"
)

type iterator struct {
	iter *freecache.Iterator
}

func (i *iterator) Next() *Entry {
	entry := i.iter.Next()
	if entry == nil {
		return nil
	}

	return &Entry{
		Key:   entry.Key,
		Value: entry.Value,
	}
}
