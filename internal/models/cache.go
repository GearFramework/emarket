package models

import (
	"context"
	"time"
)

type Cachable interface {
	Set(ctx context.Context, key string, value interface{}, exp time.Duration) error
	Get(ctx context.Context, key string) (interface{}, error)
	Exists(ctx context.Context, key string) bool
	ExistsMulti(ctx context.Context, keys ...string) (bool, int64)
}
