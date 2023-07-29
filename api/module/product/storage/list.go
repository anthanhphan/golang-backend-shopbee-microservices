package productstorage

import (
	"context"
	"shopbee/common"
	productmodel "shopbee/module/product/model"
)

func (s *productMySql) ListDataWithCondition(
	context context.Context,
	filter *productmodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]productmodel.Product, error) {
	var result []productmodel.Product

	db := s.db.Table(productmodel.Product{}.TableName()).Where("status in (1)")

	if f := filter; f != nil {
		if f.ShopId > 0 {
			db = db.Where("shop_id = ?", f.ShopId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
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
