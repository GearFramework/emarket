package db

import (
	"fmt"
	"github.com/GearFramework/emarket/internal/entities"
	"strings"
)

const (
	QUERY_CUSTOMER_BY_ID = `
		SELECT *
          FROM auth.customers
		 WHERE id = ?
	`
	QUERY_BY = "SELECT * FROM auth.%s WHERE %s"
)

func (s *Storage) Get(table string, pk map[string]any) (interface{}, error) {
	var condition []string
	vals := []interface{}{}
	for fieldName, pkValue := range pk {
		condition = append(condition, fieldName+" = ?")
		vals = append(vals, pkValue)
	}
	var customer entities.Customer
	err := s.conn.DB.Get(
		&customer,
		fmt.Sprintf(QUERY_BY, table, strings.Join(condition, " AND ")),
		vals...,
	)
	return customer, err
}

func (s *Storage) Insert() {

}
