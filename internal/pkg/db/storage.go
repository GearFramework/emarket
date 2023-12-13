package db

import (
	"context"
	"sync"
)

type Storage struct {
	sync.RWMutex
	connection *StorageConnection
}

func NewStorage(config *StorageConfig) *Storage {
	return &Storage{
		connection: NewConnection(config),
	}
}

func (s *Storage) InitStorage() error {
	if err := s.connection.Open(); err != nil {
		return err
	}
	_, err := s.connection.DB.ExecContext(context.Background(), `
		CREATE SCHEMA IF NOT EXISTS urls
	`)
	if err != nil {
		return err
	}
	_, err = s.connection.DB.ExecContext(context.Background(), `
		CREATE TABLE IF NOT EXISTS urls.shortly (
		    code VARCHAR(8),
			url VARCHAR(1024),
			user_id INT NOT NULL, 
			is_deleted BOOLEAN DEFAULT false,
		    CONSTRAINT code_url PRIMARY KEY (code, url)
		)
	`)
	if err != nil {
		return err
	}
	_, err = s.connection.DB.ExecContext(context.Background(), `
		CREATE INDEX IF NOT EXISTS idx_user_id ON urls.shortly (user_id)
	`)
	return err
}

func (s *Storage) Close() {
	s.connection.Close()
}

func (s *Storage) Ping() error {
	return s.connection.Ping()
}
