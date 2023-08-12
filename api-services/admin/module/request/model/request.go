package reqmodel

import (
	"shopbee/common"
	"time"
)

type RequestUpgrade struct {
	UserId    int                `json:"user_id" gorm:"column:user_id;"`
	ReqStatus string             `json:"req_status,omitempty" gorm:"column:request_status;default:pending;"`
	Status    int                `json:"status" gorm:"column:status;default:1;"`
	CreatedAt *time.Time         `json:"created_at,omitempty" gorm:"column:created_at;"`
	UpdatedAt *time.Time         `json:"updated_at,omitempty" gorm:"column:updated_at;"`
	User      *common.SimpleUser `json:"user" gorm:"preload:false"`
}

func (RequestUpgrade) TableName() string {
	return "request_upgrades"
}

func (r *RequestUpgrade) Mask(isAdminOrOwner bool) {
	// r.GenUID(common.DbTypeUser)
	if u := r.User; u != nil {
		u.Mask(isAdminOrOwner)
	}
}

type RequestBanUser struct {
	UserId    int                `json:"user_id" gorm:"column:user_id;"`
	ShopId    int                `json:"shop_id" gorm:"column:shop_id;"`
	Reason    string             `json:"reason,omitempty" gorm:"column:reason;"`
	ReqStatus string             `json:"rp_status,omitempty" gorm:"column:report_status;default:pending;"`
	Status    int                `json:"status" gorm:"column:status;default:1;"`
	CreatedAt *time.Time         `json:"created_at,omitempty" gorm:"column:created_at;"`
	UpdatedAt *time.Time         `json:"updated_at,omitempty" gorm:"column:updated_at;"`
	User      *common.SimpleUser `json:"user" gorm:"preload:false"`
	Shop      *common.SimpleUser `json:"shop" gorm:"preload:false;foreignKey:ShopId;references:Id;"`
}

func (r *RequestBanUser) Mask(isAdminOrOwner bool) {
	// r.GenUID(common.DbTypeUser)
	if u := r.User; u != nil {
		u.Mask(isAdminOrOwner)
	}

	if s := r.Shop; s != nil {
		s.Mask(isAdminOrOwner)
	}
}

func (RequestBanUser) TableName() string {
	return "report_accounts"
}
