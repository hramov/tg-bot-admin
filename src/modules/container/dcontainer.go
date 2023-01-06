package domainContainer

import (
	"fmt"
	"sync"
)

type DCStorage struct {
	mu   sync.Mutex
	data map[string]interface{}
}

var domainStore DCStorage

func New() {
	domainStore.data = make(map[string]interface{})
}

func Put[T any](key string, value T) {
	domainStore.mu.Lock()
	domainStore.data[key] = value
	domainStore.mu.Unlock()
}

func Pick[T any](key string) (T, error) {
	domainStore.mu.Lock()
	domain, exists := domainStore.data[key]
	domainStore.mu.Unlock()
	if !exists {
		return *new(T), fmt.Errorf("no data found by key %s", key)
	}
	return domain.(T), nil
}
