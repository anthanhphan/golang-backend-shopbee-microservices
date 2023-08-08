package wishlistmodel

import (
	"time"
)

type WishList struct {
	UserId    int        `json:"user_id" gorm:"column:user_id;"`
	ProductId int        `json:"product_id" gorm:"column:product_id;"`
	Status    int        `json:"status" gorm:"column:status;default:1;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"`
}

func (WishList) TableName() string {
	return "wish_lists"
}
