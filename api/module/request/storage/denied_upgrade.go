package reqstorage

import (
	"context"
	"shopbee/common"
	reqmodel "shopbee/module/request/model"
	mailservice "shopbee/module/sendmail"
	usermodel "shopbee/module/user/model"
)

func (s *reqMySql) DenyRequestUpgrade(
	ctx context.Context,
	data *reqmodel.RequestUpgrade,
) error {
	db := s.db

	// Update status of request
	if err := db.Table(reqmodel.RequestUpgrade{}.TableName()).
		Where("user_id = ?", data.UserId).
		Updates(map[string]interface{}{
			"request_status": "denied",
		}).Error; err != nil {

		return common.ErrDB(err)
	}

	var user *usermodel.User
	if err := db.Table("users").
		Where("id = ?", data.UserId).First(&user).Error; err != nil {
		return common.ErrDB(err)
	}

	mailservice.SendMail(user.Email, "Denied upgrade to retailer", mailservice.DeniedUpgrade)

	return nil
}
