package cartstorage

import "gorm.io/gorm"

type cartMySql struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *cartMySql {
	return &cartMySql{db: db}
}
