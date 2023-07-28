package productmodel

import "shopbee/common"

const EntityName = "Product"

type Product struct {
	common.SQLModel `json:",inline"`
	ShopId          int           `json:"-" gorm:"column:shop_id;"`
	CategoryId      int           `json:"category_id" gorm:"column:category_id;"`
	Name            string        `json:"name" gorm:"column:name;"`
	Price           float32       `json:"price" gorm:"column:price;"`
	Description     string        `json:"description" gorm:"column:description;"`
	Quantity        int           `json:"quantity" gorm:"column:quantity;"`
	Condition       string        `json:"condition" gorm:"column:condition;default:new;"`
	Image           *common.Image `json:"imagge" gorm:"column:image;"`
}

func (Product) TableName() string {
	return "products"
}

func (r *Product) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeProduct)
}
