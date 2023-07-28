package reqstorage

import "gorm.io/gorm"

type reqMySql struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *reqMySql {
	return &reqMySql{db: db}
}
