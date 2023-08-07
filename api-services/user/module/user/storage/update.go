package userstorage

import (
	"context"
	"shopbee/common"
	usermodel "shopbee/module/user/model"
)

func (s *userMySql) UpdateData(
	ctx context.Context,
	id int,
	data *usermodel.UserUpdate,
) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
