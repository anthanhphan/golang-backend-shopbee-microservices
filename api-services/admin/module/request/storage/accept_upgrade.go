package reqstorage

import (
	"context"
	"shopbee/common"
	reqmodel "shopbee/module/request/model"
)

func (s *reqMySql) AcceptRequestUpgrade(
	ctx context.Context,
	data *reqmodel.RequestUpgrade,
) error {
	db := s.db

	// Update role for user
	if err := db.Table("users").
		Where("id = ?", data.UserId).
		Updates(map[string]interface{}{
			"role": "retailer",
		}).Error; err != nil {

		return common.ErrDB(err)
	}

	var user *common.SimpleUser
	if err := db.Table("users").
		Where("id = ?", data.UserId).First(&user).Error; err != nil {
		return common.ErrDB(err)
	}

	// Update status of request
	if err := db.Table(reqmodel.RequestUpgrade{}.TableName()).
		Where("user_id = ?", data.UserId).
		Updates(map[string]interface{}{
			"request_status": "accepted",
		}).Error; err != nil {

		return common.ErrDB(err)
	}

	return nil
}
