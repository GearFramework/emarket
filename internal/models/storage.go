package models

import "sync"

type Storable interface {
	sync.Locker
	Ping() error
	Close()
}
