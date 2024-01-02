package customer

import (
	"github.com/GearFramework/emarket/internal/models"
	"github.com/GearFramework/emarket/internal/pkg/cache/redis"
)

type RepositoryCustomers struct {
	models.Repository
	Cache *redis.CacheRedis
}
