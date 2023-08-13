package cartstorage

import (
	"context"
	"shopbee/common"
	cartmodel "shopbee/module/cart/model"
)

func (s *cartMySql) UpdateCart(
	ctx context.Context,
	data *cartmodel.CartUpdate,
) error {
	db := s.db.Table(cartmodel.CartUpdate{}.TableName())

	if err := db.
		Where("user_id = ? AND product_id = ?", data.UserId, data.ProductId).
		Updates(&data).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
