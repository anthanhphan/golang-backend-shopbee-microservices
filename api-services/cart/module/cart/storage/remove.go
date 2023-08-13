package cartstorage

import (
	"context"
	"shopbee/common"
	cartmodel "shopbee/module/cart/model"
)

func (s *cartMySql) RemoveProductFromCart(
	ctx context.Context,
	productId int,
	userId int,
) error {
	db := s.db.Table(cartmodel.Cart{}.TableName())

	if err := db.
		Where("product_id = ? AND user_id = ?", productId, userId).
		Delete(&cartmodel.Cart{}).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
