package userstorage

import "gorm.io/gorm"

type userMySql struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *userMySql {
	return &userMySql{db: db}
}
