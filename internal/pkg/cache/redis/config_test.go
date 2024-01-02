package redis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type RedisTestConfig struct {
	name   string
	config *ConnectionConfig
}

func (test *RedisTestConfig) tfunc(t *testing.T) {
	c := NewCache(test.config)
	assert.NotNil(t, c)
	err := c.InitCache()
	assert.NoError(t, err)
	c.Close()
}

var tests []RedisTestConfig = []RedisTestConfig{
	{
		name: "Default settings",
		config: &ConnectionConfig{
			Host:     DefaultHost,
			Port:     DefaultPort,
			Password: DefaultPassword,
			Database: DefaultDB,
		},
	},
	{
		name:   "Settings from .env file",
		config: NewRedisConfig(),
	},
}

func TestRedisConfig(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.tfunc(t)
		})
	}
}
