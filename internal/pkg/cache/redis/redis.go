package redis

import (
	"context"
	"sync"
	"time"
)

type CacheRedis struct {
	sync.RWMutex
	conn *Connection
}

func NewCache(config *ConnectionConfig) *CacheRedis {
	return &CacheRedis{
		conn: NewConnection(config),
	}
}

func (cache *CacheRedis) InitCache() error {
	return cache.conn.Open()
}

func (cache *CacheRedis) Close() {
	cache.conn.Close()
}

func (cache *CacheRedis) Ping() error {
	return cache.conn.Ping()
}

func (cache *CacheRedis) Set(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	cmd := cache.conn.DB.Set(ctx, key, value, exp)
	return cmd.Err()
}

func (cache *CacheRedis) Get(ctx context.Context, key string) (interface{}, error) {
	cmd := cache.conn.DB.Get(ctx, key)
	return cmd.Val(), cmd.Err()
}

func (cache *CacheRedis) GetMap(ctx context.Context, key, col string) (interface{}, error) {
	return struct{}{}, nil
}

func (cache *CacheRedis) Exists(ctx context.Context, key string) bool {
	val, err := cache.conn.DB.Exists(ctx, key).Result()
	return val > 0 && err == nil
}

func (cache *CacheRedis) ExistsMulti(ctx context.Context, keys ...string) (bool, int64) {
	val, err := cache.conn.DB.Exists(ctx, keys...).Result()
	return val == int64(len(keys)) && err == nil, val
}

func (cache *CacheRedis) ExistsMap(ctx context.Context, key, col string) bool {
	val, err := cache.conn.DB.HExists(ctx, key, col).Result()
	return val && err == nil
}
