package auth

import (
	"github.com/GearFramework/emarket/internal/app"
	"github.com/GearFramework/emarket/internal/pkg/auth"
	"github.com/GearFramework/emarket/internal/pkg/server"
	"os"
	"strconv"
)

const (
	DefaultPort = 8080
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

func NewServerConfig() *server.Config {
	port, err := strconv.Atoi(os.Getenv("BACKEND_PORT"))
	if err != nil {
		port = DefaultPort
	}
	return &server.Config{
		Addr: os.Getenv("BACKEND_ADDR"),
		Port: port,
	}
}
