package productmodel

import (
	"shopbee/common"
	usermodel "shopbee/module/user/model"
)

const EntityName = "Product"

type Product struct {
	common.SQLModel `json:",inline"`
	ShopId          int             `json:"-" gorm:"column:shop_id;"`
	CategoryId      int             `json:"category_id" gorm:"column:category_id;"`
	Name            string          `json:"name" gorm:"column:name;"`
	Price           float32         `json:"price" gorm:"column:price;"`
	Description     string          `json:"description" gorm:"column:description;"`
	Quantity        int             `json:"quantity" gorm:"column:quantity;"`
	Condition       string          `json:"condition" gorm:"column:condition;default:new;"`
	Image           *common.Image   `json:"image" gorm:"column:image;"`
	Shop            *usermodel.User `json:"shop" gorm:"preload:false;foreignKey:ShopId;references:Id;"`
}

func (Product) TableName() string {
	return "products"
}

func (r *Product) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeProduct)

	if s := r.Shop; s != nil {
		s.Mask(isAdminOrOwner)
	}
}
