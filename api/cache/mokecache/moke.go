package mokecache

import (
	"context"
	"errors"
	"time"
)

type Moke struct {
	cache map[string]interface{}
}

func New() *Moke {
	return &Moke{
		cache: make(map[string]interface{}),
	}
}

func (m *Moke) Set(c context.Context, key string, value interface{}, duration time.Duration) error {
	m.cache[key] = value
	return nil
}

func (m *Moke) Get(c context.Context, key string) ([]byte, error) {
	v, ok := m.cache[key]
	if !ok {
		return nil, errors.New("cache: Get key: not found:" + key)
	}
	return v.([]byte), nil
}
