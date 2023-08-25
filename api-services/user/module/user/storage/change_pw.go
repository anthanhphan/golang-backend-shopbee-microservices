package userstorage

import (
	"context"
	usermodel "shopbee/module/user/model"
)

func (s *userMySql) ChangePassword(
	ctx context.Context,
	id int,
	password string,
) error {
	db := s.db

	if err := db.Table(usermodel.User{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"password": password,
		}).Error; err != nil {

		return err
	}

	return nil
}
