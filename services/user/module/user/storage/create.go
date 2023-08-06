package userstorage

import (
	"context"
	"shopbee/common"
	usermodel "shopbee/module/user/model"
)

func (store *userMySql) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	db := store.db.Begin()

	if err := db.Table(data.TableName()).Create(&data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
