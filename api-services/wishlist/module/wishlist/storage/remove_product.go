package wishliststorage

import (
	"context"
	"shopbee/common"
	wishlistmodel "shopbee/module/wishlist/model"
)

func (s *wishListMySql) RemoveToWishList(
	ctx context.Context,
	productId int,
	userId int,
) error {
	db := s.db.Table(wishlistmodel.WishList{}.TableName())

	if err := db.
		Where("product_id = ? AND user_id = ?", productId, userId).
		Delete(&wishlistmodel.WishList{}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
