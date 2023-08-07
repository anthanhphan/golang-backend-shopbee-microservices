package productstorage

import (
	"context"
	"shopbee/common"
	productmodel "shopbee/module/product/model"
)

func (s *productMySql) CreateProduct(
	ctx context.Context,
	data *productmodel.Product,
) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
