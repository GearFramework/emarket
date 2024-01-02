package redis

import (
	"fmt"
	"github.com/GearFramework/emarket/internal/pkg/gear"
	"os"
)

const (
	DefaultHost                        = "localhost"
	DefaultPort                        = 6379
	DefaultPassword                    = ""
	DefaultDB                          = 0
	defaultConnectAttempts             = 10
	defaultConnectDelay                = 10
	defaultBackgroundCheckConnectDelay = 10
	templateDSN                        = "%s:%d"
)

type ConnectionConfig struct {
	Host                        string // localhost
	Port                        uint16 // 6379
	Password                    string
	Database                    int
	ConnectAttempts             int8  // максимальное количество попыток подключения (по-умолчанию 10)
	ConnectDelay                uint8 // задержка перед повторным подключением в секундах (по-умолчанию 10)
	BackgroundCheckConnectDelay uint8 // по-умолчанию 10
}

func NewRedisConfig() *ConnectionConfig {
	return &ConnectionConfig{
		Host:     gear.Getenv("REDIS_HOST", DefaultHost),
		Port:     gear.AtoUI16(os.Getenv("REDIS_PORT"), DefaultPort),
		Password: gear.Getenv("REDIS_PASSWORD", DefaultPassword),
		Database: gear.AtoI(os.Getenv("REDIS_DATABASE"), DefaultDB),
	}
}

// Возвращает строку подключения к Redis
func (config *ConnectionConfig) GetDSN() string {
	return fmt.Sprintf(
		templateDSN,
		config.Host,
		config.Port,
	)
}
