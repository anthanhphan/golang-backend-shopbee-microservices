package model

import "time"

type User struct {
	Id          string     `json:"id" gorm:"column:ID;"`
	FullName    string     `json:"fullname" gorm:"column:FULL_NAME;"`
	Password    string     `json:"password" gorm:"column:PASSWORD;"`
	PhoneNumber string     `json:"phonenumber" gorm:"column:PHONE_NUMBER;"`
	Email       string     `json:"email" gorm:"column:EMAIL;"`
	Address     string     `json:"address" gorm:"column:ADDRESS;"`
	Sex         string     `json:"sex" gorm:"column:SEX;"`
	Date        *time.Time `json:"date" gorm:"column:DATE;"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:CREATED_AT;"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"column:UPDATED_AT;"`
	Token       string
}

type UserCreation struct {
	Id          string `json:"id" gorm:"column:ID;"`
	FullName    string `json:"fullname" gorm:"column:FULL_NAME;"`
	Password    string `json:"password" gorm:"column:PASSWORD;"`
	PhoneNumber string `json:"phonenumber" gorm:"column:PHONE_NUMBER;"`
	Email       string `json:"email" gorm:"column:EMAIL;"`
	Sex         string `json:"sex" gorm:"column:SEX;"`
}
