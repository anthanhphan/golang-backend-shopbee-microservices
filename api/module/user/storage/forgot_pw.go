package userstorage

import (
	"context"
	"shopbee/common"
	usermodel "shopbee/module/user/model"
)

func (s *userMySql) ForgotPassword(
	ctx context.Context,
	email string,
	data *usermodel.User,
) error {
	db := s.db

	if err := db.Where("email = ?", email).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
