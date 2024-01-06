package db

import (
	"sync"
)

type Storage struct {
	sync.RWMutex
	conn *StorageConnection
}

func NewStorage(config *StorageConfig) *Storage {
	return &Storage{
		conn: NewConnection(config),
	}
}

func (s *Storage) InitStorage() error {
	var err error
	if err = s.conn.Open(); err == nil {
		go s.conn.BackgroundCheckConnect()
	}
	return err
}

func (s *Storage) Close() {
	s.conn.Close()
}

func (s *Storage) Ping() error {
	return s.conn.Ping()
}
