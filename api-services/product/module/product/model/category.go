package productmodel

import "shopbee/common"

type Category struct {
	common.SQLModel `json:",inline"`
	RId             int           `json:"rid,omitempty" gorm:"-"`
	Name            string        `json:"name" gorm:"column:name;"`
	Image           *common.Image `json:"image" gorm:"column:image;"`
}

func (Category) TableName() string {
	return "categories"
}

func (r *Category) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeProduct)
}
