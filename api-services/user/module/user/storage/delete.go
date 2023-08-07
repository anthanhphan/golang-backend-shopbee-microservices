package userstorage

import (
	"context"
	"shopbee/common"
	usermodel "shopbee/module/user/model"
)

func (s *userMySql) SoftDelete(
	ctx context.Context,
	id int,
) error {
	db := s.db

	if err := db.Table(usermodel.User{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": 0,
		}).Error; err != nil {

		return common.ErrDB(err)
	}

	return nil
}
