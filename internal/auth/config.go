package auth

import (
	"github.com/GearFramework/emarket/internal/app"
	"github.com/GearFramework/emarket/internal/pkg/auth"
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
