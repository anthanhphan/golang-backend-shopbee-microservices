package cartmodel

import "time"

type Cart struct {
	UserId          int        `json:"user_id" gorm:"column:user_id;"`
	ProductId       int        `json:"product_id" gorm:"column:product_id;"`
	ProductQuantity int        `json:"quantity" gorm:"column:product_quantity;"`
	Status          int        `json:"status" gorm:"column:status;default:1;"`
	CreatedAt       *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"`
}

func (Cart) TableName() string {
	return "carts"
}

type CartCreate struct {
	UserId          int `json:"user_id" gorm:"column:user_id;"`
	ProductId       int `json:"product_id" gorm:"column:product_id;"`
	ProductQuantity int `json:"quantity" gorm:"column:product_quantity;"`
}

func (CartCreate) TableName() string {
	return "carts"
}
