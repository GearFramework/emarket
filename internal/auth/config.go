package auth

import (
	"github.com/GearFramework/emarket/internal/app"
	"github.com/GearFramework/emarket/internal/pkg/auth"
	"github.com/GearFramework/emarket/internal/pkg/db"
	"os"
)

func NewAuthConfig() *app.ServiceAuthConfig {
	key := os.Getenv("AUTH_KEY")
	if key == "" {
		key = auth.SecretKey
	}
	return &app.ServiceAuthConfig{
		AuthKey: key,
	}
}

func NewStorageConfig() *db.StorageConfig {
	return &db.StorageConfig{
		//Host:     gear.Getenv("REDIS_HOST", redis.DefaultHost),
		//Port:     gear.AtoUI16(os.Getenv("REDIS_PORT"), redis.DefaultPort),
		//Password: gear.Getenv("REDIS_PASSWORD", redis.DefaultPassword),
		//Database: gear.AtoI(os.Getenv("REDIS_DATABASE"), redis.DefaultDB),
	}
}
