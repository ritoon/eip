package cache

import (
	"context"
	"time"
)

// Cacher is an interface to set and get cache
type Cacher interface {
	Set(c context.Context, key string, value interface{}, duration time.Duration) error
	Get(c context.Context, key string) ([]byte, error)
}
