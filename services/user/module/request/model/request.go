package reqmodel

import (
	"time"
)

type UserUpgrade struct {
	Role string `json:"role" gorm:"column:role;"`
}

func (UserUpgrade) TableName() string {
	return "users"
}

type RequestUpgrade struct {
	UserId    int        `json:"user_id" gorm:"column:user_id;"`
	ReqStatus string     `json:"req_status,omitempty" gorm:"column:request_status;default:pending;"`
	Status    int        `json:"status" gorm:"column:status;default:1;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"`
}

func (RequestUpgrade) TableName() string {
	return "request_upgrades"
}

type RequestBanUser struct {
	UserId    int        `json:"user_id" gorm:"column:user_id;"`
	ShopId    int        `json:"shop_id" gorm:"column:shop_id;"`
	Reason    string     `json:"reason,omitempty" gorm:"column:reason;"`
	ReqStatus string     `json:"rp_status,omitempty" gorm:"column:report_status;default:pending;"`
	Status    int        `json:"status" gorm:"column:status;default:1;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"`
}

func (RequestBanUser) TableName() string {
	return "report_accounts"
}
