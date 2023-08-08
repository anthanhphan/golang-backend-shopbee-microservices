package wishliststorage

import (
	"context"
	"shopbee/common"
	wishlistmodel "shopbee/module/wishlist/model"
)

func (s *wishListMySql) AddToWishList(
	ctx context.Context,
	data *wishlistmodel.WishList,
) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
