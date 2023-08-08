package wishliststorage

import (
	"context"
	"shopbee/common"
	wishlistmodel "shopbee/module/wishlist/model"
)

func (s *wishListMySql) ViewMyWishList(
	context context.Context,
	id int,
) ([]wishlistmodel.WishList, error) {
	var resullt []wishlistmodel.WishList

	db := s.db.Table(wishlistmodel.WishList{}.TableName()).Where("status in (1)")
	db = db.Select("product_id").Where("user_id = ?", id)

	if err := db.Find(&resullt).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return resullt, nil
}
