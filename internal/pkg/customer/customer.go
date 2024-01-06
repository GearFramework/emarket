package customer

import (
	"context"
	"errors"
	"fmt"
	"github.com/GearFramework/emarket/internal/entities"
	"github.com/GearFramework/emarket/internal/models"
	"github.com/GearFramework/emarket/internal/pkg/cache/redis"
)

var ErrAlreadyRegistered error = errors.New("customer already registered")
var keys []string = []string{"email", "phone"}

type RepositoryCustomers struct {
	models.Repository
	Storage  *models.Storable
	Cache    *redis.CacheRedis
	wrappers []interface{}
}

type WrapperStorage struct {
}

func NewRepository(Storage *models.Storable, Cache *redis.CacheRedis) *RepositoryCustomers {
	return &RepositoryCustomers{}
}

func (repo *RepositoryCustomers) New(ctx context.Context, props map[string]interface{}) (*entities.Customer, error) {
	repo.Lock()
	defer repo.Unlock()
	pk := repo.GetCacheKeys(props)
	if repo.Exists(ctx, pk...) {
		return nil, ErrAlreadyRegistered
	}
	return &entities.Customer{}, nil
}

func (repo *RepositoryCustomers) GetCacheKeys(props map[string]interface{}) []string {
	var pk []string
	for _, f := range keys {
		if v, ok := props[f]; ok {
			pk = append(pk, fmt.Sprintf("customer:%s:%v", f, v))
		}
	}
	return pk
}

func (repo *RepositoryCustomers) Exists(ctx context.Context, keys ...string) bool {
	exists, count := repo.Cache.ExistsMulti(ctx, keys...)
	return exists || (exists == false && count > 0)
}
