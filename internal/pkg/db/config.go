package db

import (
	"fmt"
	"github.com/GearFramework/emarket/internal/pkg/gear"
	"os"
)

const (
	DefaultHost                        = "localhost"
	DefaultPort                        = 5432
	DefaultConnectAttempts             = 10
	DefaultConnectDelay                = 10 // seconds
	DefaultConnectMaxOpens             = 10
	DefaultBackgroundCheckConnectDelay = 10 // seconds
	templateDSN                        = "postgres://%s:%s@%s:%d/%s"
)

type StorageConfig struct {
	Host                        string // localhost
	Port                        uint16 // 5432
	Username                    string
	Password                    string
	Database                    string
	ConnectAttempts             int8   // максимальное количество попыток подключения (по-умолчанию 10)
	ConnectDelay                uint64 // задержка перед повторным подключением в секундах (по-умолчанию 10)
	ConnectMaxOpens             int    // максимальное количество открытых соединений (по-умолчанию 10)
	BackgroundCheckConnectDelay uint64 // по-умолчанию 10
}

func NewStorageConfig() *StorageConfig {
	return &StorageConfig{
		Host:                        gear.Getenv("PG_HOST", DefaultHost),
		Port:                        gear.AtoUI16(os.Getenv("PG_PORT"), DefaultPort),
		Username:                    gear.Getenv("PG_USER", ""),
		Password:                    gear.Getenv("PG_PASSWORD", ""),
		Database:                    gear.Getenv("PG_DATABASE", ""),
		ConnectMaxOpens:             gear.AtoI(os.Getenv("PG_MAX_OPENS"), DefaultConnectMaxOpens),
		ConnectAttempts:             gear.AtoI8(os.Getenv("PG_CONN_ATTEMPTS"), DefaultConnectAttempts),
		ConnectDelay:                gear.AtoUI64(os.Getenv("PG_CONN_DELAY"), DefaultConnectDelay),
		BackgroundCheckConnectDelay: gear.AtoUI64(os.Getenv("PG_BG_CONN_DELAY"), DefaultBackgroundCheckConnectDelay),
	}
}

func (config *StorageConfig) GetDSN() string {
	return fmt.Sprintf(
		templateDSN,
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
}
