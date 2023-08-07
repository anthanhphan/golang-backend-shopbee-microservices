package categorystorage

import "gorm.io/gorm"

type categoryMySql struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *categoryMySql {
	return &categoryMySql{db: db}
}
