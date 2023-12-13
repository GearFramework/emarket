package entities

type Customer struct {
	ID             uint64
	FirstName      string
	LastName       string
	PatronymicName string
	PhoneNumber    string
	Email          string
}

type CustomerCart map[uint64]Cart
