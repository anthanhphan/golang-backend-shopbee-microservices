package wishliststorage

import "gorm.io/gorm"

type wishListMySql struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *wishListMySql {
	return &wishListMySql{db: db}
}
