package paymentmodel

import (
	"shopbee/common"
)

type Payment struct {
	common.SQLModel `json:",inline"`
	UserId          int     `json:"-" gorm:"column:user_id;"`
	PaymenStatus    string  `json:"payment_status,omitempty" gorm:"column:payment_status;"`
	PaymenMethod    string  `json:"payment_method" gorm:"column:payment_method;"`
	Amount          float64 `json:"amount" gorm:"column:amount;"`
}

func (Payment) TableName() string {
	return "payments"
}
