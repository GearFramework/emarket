package redis

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedis(t *testing.T) {
	cache := NewCache(NewRedisConfig())
	assert.NotNil(t, cache)
	err := cache.InitCache()
	assert.NoError(t, err)
	err = cache.Set(context.Background(), "test", "test_val", 0)
	assert.NoError(t, err)
	ex := cache.Exists(context.Background(), "test")
	assert.Equal(t, true, ex)
	val, err := cache.Get(context.Background(), "test")
	assert.NoError(t, err)
	assert.Equal(t, "test_val", val)
	ex = cache.Exists(context.Background(), "test1")
	assert.Equal(t, false, ex)
	cache.Close()
}
