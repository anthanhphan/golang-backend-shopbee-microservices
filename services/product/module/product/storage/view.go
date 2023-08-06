package productstorage

import (
	"context"
	productmodel "shopbee/module/product/model"
)

func (s *productMySql) ViewProduct(
	context context.Context,
	id int,
	moreKey ...string,
) (*productmodel.Product, error) {
	var result productmodel.Product

	db := s.db.Table(productmodel.Product{}.TableName()).Where("status in (1)")
	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}

	if err := db.Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
