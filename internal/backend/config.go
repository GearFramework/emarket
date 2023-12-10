package backend

import (
	"github.com/GearFramework/emarket/internal/pkg/server"
	"os"
	"strconv"
)

func NewServerConfig() *server.Config {
	port, _ := strconv.Atoi(os.Getenv("BACKEND_PORT"))
	return &server.Config{
		Addr: os.Getenv("BACKEND_ADDR"),
		Port: port,
	}
}
