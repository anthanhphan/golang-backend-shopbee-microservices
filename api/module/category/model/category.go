package categorymodel

import "shopbee/common"

type Category struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"name" gorm:"column:name;"`
	Image           *common.Image `json:"image" gorm:"column:image;"`
}

func (Category) TableName() string {
	return "categories"
}

func (r *Category) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeProduct)
}
