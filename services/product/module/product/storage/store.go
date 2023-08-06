package productstorage

import "gorm.io/gorm"

type productMySql struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *productMySql {
	return &productMySql{db: db}
}
