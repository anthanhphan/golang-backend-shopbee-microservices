package paymentstorage

import "gorm.io/gorm"

type paymentMySql struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *paymentMySql {
	return &paymentMySql{db: db}
}
