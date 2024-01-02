package models

type Connector interface {
	Open() error
	Close()
	Ping() error
}
