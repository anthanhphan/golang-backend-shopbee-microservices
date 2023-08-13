package ordermodel

import "shopbee/common"

type Order struct {
	common.SQLModel `json:",inline"`
	Id              int    `json:"-" gorm:"column:id;"`
	UserId          int    `json:"-" gorm:"column:user_id;"`
	ShopId          int    `json:"shop_id" gorm:"column:shop_id"`
	TotalPrice      int    `json:"total_price" gorm:"column:total_price"`
	PaymentId       int    `json:"payment_id" gorm:"column:payment_id;"`
	ShippingAddr    string `json:"shipping_addr" gorm:"column:shipping_addr;"`
	OrderStatus     string `json:"order_status" gorm:"column:order_status;default:pending"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderDetail struct {
	common.SQLModel `json:",inline"`
	OrderId         int `json:"-" gorm:"column:order_id;"`
	ProductOrigin   any `json:"product_origin" gorm:"column:product_origin;"`
}

func (OrderDetail) TableName() string {
	return "order_details"
}

type OrderCreate struct {
	ShopId       string                   `json:"shop_id"`
	PaymentId    string                   `json:"payment_id"`
	ShippingAddr string                   `json:"shipping_addr"`
	TotalPrice   int                      `json:"total_price" gorm:"column:total_price"`
	ProductList  []map[string]interface{} `json:"product_list"`
}
