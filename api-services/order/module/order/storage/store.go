package orderstorage

import "gorm.io/gorm"

type orderMySql struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *orderMySql {
	return &orderMySql{db: db}
}
