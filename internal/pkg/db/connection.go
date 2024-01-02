package db

import (
	"context"
	"github.com/GearFramework/emarket/internal/pkg/alog"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type StorageConnection struct {
	DB        *sqlx.DB
	config    *StorageConfig
	pgxConfig *pgx.ConnConfig
	logger    *alog.Alog
}

func NewConnection(config *StorageConfig) *StorageConnection {
	return &StorageConnection{
		config: config,
		logger: alog.NewLogger(),
	}
}

func (conn *StorageConnection) Open() error {
	var err error = nil
	if conn.pgxConfig, err = conn.getPgxConfig(); err != nil {
		return err
	}
	return conn.openSqlxViaPooler()
}

// openSqlxViaPooler открытие пула соединений
func (conn *StorageConnection) openSqlxViaPooler() error {
	db := stdlib.OpenDB(*conn.pgxConfig)
	conn.DB = sqlx.NewDb(db, "pgx")
	conn.DB.SetMaxOpenConns(conn.config.ConnectMaxOpens)
	return nil
}

func (conn *StorageConnection) getPgxConfig() (*pgx.ConnConfig, error) {
	pgxConfig, err := pgx.ParseConfig(conn.config.ConnectionDSN)
	if err != nil {
		conn.logger.Errorf("Unable to parse DSN: %v\n", err)
		return nil, err
	}
	return pgxConfig, nil
}

func (conn *StorageConnection) Ping() error {
	return conn.DB.PingContext(context.Background())
}

func (conn *StorageConnection) Close() {
	if conn.Ping() == nil {
		conn.logger.Info("Close storage connection")
		if err := conn.DB.Close(); err != nil {
			conn.logger.Error(err.Error())
		}
	}
}
