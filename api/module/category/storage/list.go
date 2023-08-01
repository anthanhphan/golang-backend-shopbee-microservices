package categorystorage

import (
	"context"
	"shopbee/common"
	categorymodel "shopbee/module/category/model"
)

func (s *categoryMySql) ListDataWithCondition(
	context context.Context,
	paging *common.Paging,
	moreKey ...string,
) ([]categorymodel.Category, error) {
	var result []categorymodel.Category

	db := s.db.Table(categorymodel.Category{}.TableName()).Where("status in (1)")

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
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(result) > 0 {
		last := result[len(result)-1]
		last.Mask(false)

		paging.NextCursor = last.FakeId.String()
	}

	return result, nil
}
