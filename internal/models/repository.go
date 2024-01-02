package models

import "sync"

type Repository interface {
	sync.Locker
}
