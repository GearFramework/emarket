package models

import "context"

type Cachable interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (interface{}, error)
	Exists(ctx context.Context, key string) bool
}
