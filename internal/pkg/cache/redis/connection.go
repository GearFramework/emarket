package redis

import (
	"context"
	"fmt"
	"github.com/GearFramework/emarket/internal/pkg/alog"
	"github.com/redis/go-redis/v9"
	"time"
)

var ctx = context.Background()

// Connection Структура для соединения с Redis
type Connection struct {
	DB     *redis.Client
	Config *ConnectionConfig
	logger *alog.Alog
}

// NewConnection Возвращает подключение к Redis
func NewConnection(config *ConnectionConfig) *Connection {
	return &Connection{
		Config: config,
		logger: alog.NewLogger(),
	}
}

// Open Подключение к Redis
func (conn *Connection) Open() error {
	var err error = nil
	var numberConnectionAttempts int8 = 1
	err = conn.tryingToConnect(numberConnectionAttempts)
	if err != nil && conn.hasConnectAttempts() {
		conn.logger.Error(err.Error())
		err = conn.attemptsConnection(numberConnectionAttempts)
	}
	return err
}

// Возвращает true если имеется возможность попытаться снова подключиться
// к Redis после первого неудачного подключения
func (conn *Connection) hasConnectAttempts() bool {
	return conn.Config.ConnectAttempts < 0 || conn.Config.ConnectAttempts > 1
}

// В цикле N-ое кол-во раз или бесконечно, через
// определённые интервалы пытается подключиться к redis
func (conn *Connection) attemptsConnection(numberConnectionAttempts int8) error {
	var err error = nil
	reconnectInterval := time.Duration(conn.Config.ConnectDelay) * time.Second
	for {
		<-time.After(reconnectInterval)
		numberConnectionAttempts++
		if err = conn.tryingToConnect(numberConnectionAttempts); err == nil {
			break
		}
		conn.logger.Error(err.Error())
		conn.logger.Info(fmt.Sprintf("Pause: %d sec", conn.Config.ConnectDelay))
		if conn.Config.ConnectAttempts > 0 &&
			numberConnectionAttempts >= conn.Config.ConnectAttempts {
			break
		}
	}
	return err
}

// N-ная попытка подключения к redis
func (conn *Connection) tryingToConnect(numberConnectionAttempts int8) error {
	conn.logger.Info(fmt.Sprintf(
		"Attempt %d of %d to connect to Redis",
		numberConnectionAttempts,
		conn.Config.ConnectAttempts,
	))
	return conn.connect()
}

// Выполняет несколько попыток подключения к redis в соответствии с конфигом config.
func (conn *Connection) connect() error {
	conn.logger.Info("Try connect to: " + conn.Config.GetDSN())
	conn.DB = redis.NewClient(&redis.Options{
		Addr:     conn.Config.GetDSN(),
		Password: conn.Config.Password,
		DB:       conn.Config.Database,
	})
	conn.logger.Info("Successful redis connected")
	return nil
}

// Ping Проверка подключения к Redis
func (conn *Connection) Ping() error {
	status := conn.DB.Ping(ctx)
	return status.Err()
}

// Close Закрывает открытое соединение с Redis
func (conn *Connection) Close() {
	if conn.DB != nil {
		if err := conn.DB.Close(); err != nil {
			conn.logger.Error(err.Error())
		} else {
			conn.logger.Info("Redis connection closed.")
		}
	}
}

// BackgroundCheckConnect Метод, который через указанное в конфиге время проверяет подключение к Redis
func (conn *Connection) BackgroundCheckConnect() {
	tryoutInterval := time.Duration(conn.Config.BackgroundCheckConnectDelay) * time.Second
	for {
		<-time.After(tryoutInterval)
		if err := conn.Ping(); err != nil {
			fmt.Println(err)
			_ = conn.connect()
		}
	}
}
