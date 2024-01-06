package db

import (
	"context"
	"fmt"
	"github.com/GearFramework/emarket/internal/pkg/alog"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"time"
)

type StorageConnection struct {
	DB        *sqlx.DB
	Config    *StorageConfig
	pgxConfig *pgx.ConnConfig
	logger    *alog.Alog
}

func NewConnection(config *StorageConfig) *StorageConnection {
	return &StorageConnection{
		Config: config,
		logger: alog.NewLogger(fmt.Sprintf(
			"PostgreSQL %s@%s:%d/%s",
			config.Username,
			config.Host,
			config.Port,
			config.Database,
		)),
	}
}

func (conn *StorageConnection) Open() error {
	var err error = nil
	var numberConnectionAttempts int8 = 1
	if conn.pgxConfig, err = conn.getPgxConfig(); err != nil {
		return err
	}
	err = conn.tryingToConnect(numberConnectionAttempts)
	if err != nil && conn.hasConnectAttempts() {
		conn.logger.Errorf("failed database pgsqlConnection: %s", err.Error())
		err = conn.attemptsConnection(numberConnectionAttempts)
	}
	return err
	//conn.logger.Info("Open storage connection")
	//var err error = nil
	//if conn.pgxConfig, err = conn.getPgxConfig(); err != nil {
	//	return err
	//}
	//return conn.openSqlxViaPooler()
}

// openSqlxViaPooler открытие пула соединений
func (conn *StorageConnection) openSqlxViaPooler() error {
	db := stdlib.OpenDB(*conn.pgxConfig)
	conn.DB = sqlx.NewDb(db, "pgx")
	conn.DB.SetMaxOpenConns(conn.Config.ConnectMaxOpens)
	return nil
}

func (conn *StorageConnection) getPgxConfig() (*pgx.ConnConfig, error) {
	pgxConfig, err := pgx.ParseConfig(conn.Config.GetDSN())
	if err != nil {
		conn.logger.Errorf("Unable to parse DSN: %s", err.Error())
		return nil, err
	}
	return pgxConfig, nil
}

// Возвращает true если имеется возможность попытаться снова подключиться
// к postgreSQL после первого неудачного подключения
func (conn *StorageConnection) hasConnectAttempts() bool {
	return conn.Config.ConnectAttempts < 0 || conn.Config.ConnectAttempts > 1
}

// В цикле N-ое кол-во раз или бесконечно, через
// определённые интервалы пытается подключиться к postgreSQL
func (conn *StorageConnection) attemptsConnection(numberStorageConnectionAttempts int8) error {
	var err error = nil
	reconnectInterval := time.Duration(conn.Config.ConnectDelay) * time.Second
	for {
		<-time.After(reconnectInterval)
		numberStorageConnectionAttempts++
		if err = conn.tryingToConnect(numberStorageConnectionAttempts); err == nil {
			break
		}
		conn.logger.Errorf("failed database conn: %s", err.Error())
		conn.logger.Errorf("pause: %d sec", conn.Config.ConnectDelay)
		if conn.Config.ConnectAttempts > 0 && numberStorageConnectionAttempts >= conn.Config.ConnectAttempts {
			break
		}
	}
	return err
}

// N-ная попытка подключения к postgreSQL
func (conn *StorageConnection) tryingToConnect(numberConnectionAttempts int8) error {
	conn.logger.Infof(
		"attempt %d of %d to connect to PostgreSQL",
		numberConnectionAttempts,
		conn.Config.ConnectAttempts,
	)
	return conn.open()
}

// Выполняет несколько попыток подключения к БД в соответствии с конфигом config.
func (conn *StorageConnection) open() error {
	conn.logger.Infof("try connect to host: %s:%d", conn.Config.Host, conn.Config.Port)
	if err := conn.openSqlxViaPooler(); err != nil {
		return err
	}
	if err := conn.Ping(); err != nil {
		return err
	}
	conn.logger.Info("successful database connected")
	return nil
}

func (conn *StorageConnection) Ping() error {
	return conn.DB.PingContext(context.Background())
}

func (conn *StorageConnection) Close() {
	if conn.Ping() == nil {
		conn.logger.Info("close storage connection")
		if err := conn.DB.Close(); err != nil {
			conn.logger.Error(err.Error())
		}
	}
}

// BackgroundCheckConnect метод, который через указанное в конфиге время проверяет подключение к PostgreSQL
func (conn *StorageConnection) BackgroundCheckConnect() {
	tryoutInterval := time.Duration(conn.Config.BackgroundCheckConnectDelay) * time.Second
	for {
		<-time.After(tryoutInterval)
		if err := conn.Ping(); err != nil {
			conn.logger.Errorf("background check connection error: %s", err.Error())
			_ = conn.open()
		}
	}
}
