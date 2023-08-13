package cartstorage

import (
	"context"
	"shopbee/common"
	cartmodel "shopbee/module/cart/model"
)

func (s *cartMySql) AddToCart(
	ctx context.Context,
	data *cartmodel.CartCreate,
) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
