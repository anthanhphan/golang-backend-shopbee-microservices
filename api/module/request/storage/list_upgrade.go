package reqstorage

import (
	"context"
	"shopbee/common"
	reqmodel "shopbee/module/request/model"
)

func (s *reqMySql) ListUpgradeWithCondition(
	context context.Context,
	filter *reqmodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]reqmodel.RequestUpgrade, error) {
	var result []reqmodel.RequestUpgrade

	db := s.db.Table(reqmodel.RequestUpgrade{}.TableName()).Where("status in (1)")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(int(offset))
	}

	if err := db.
		Limit(int(paging.Limit)).
		Order("user_id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
