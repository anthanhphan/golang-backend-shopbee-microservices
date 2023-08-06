package reqstorage

import (
	"context"
	"shopbee/common"
	reqmodel "shopbee/module/request/model"
)

func (s *reqMySql) CreateRequestUpgrade(
	ctx context.Context,
	data *reqmodel.RequestUpgrade,
) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *reqMySql) CreateRequestBanUser(
	ctx context.Context,
	data *reqmodel.RequestBanUser,
) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
