package reqstorage

import (
	"context"
	"shopbee/common"
	reqmodel "shopbee/module/request/model"

	"gorm.io/gorm"
)

func (s *reqMySql) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string,
) (*reqmodel.RequestUpgrade, error) {
	db := s.db.Table(reqmodel.RequestUpgrade{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var data reqmodel.RequestUpgrade

	if err := db.Where(conditions).First(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &data, nil
}
