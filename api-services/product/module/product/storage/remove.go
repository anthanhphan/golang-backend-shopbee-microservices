package productstorage

import (
	"context"
	"shopbee/common"
	productmodel "shopbee/module/product/model"
)

func (s *productMySql) RemoveProduct(
	ctx context.Context,
	productId int,
) error {
	db := s.db

	// Remove all product are in wishlist
	var wishlists interface{}
	if err := db.
		Table("wishlists").
		Where("product_id = ?", productId).
		Delete(&wishlists).
		Error; err != nil {
		return common.ErrDB(err)
	}

	// Remove all product are in cart
	var cart interface{}
	if err := db.
		Table("carts").
		Where("product_id = ?", productId).
		Delete(&cart).
		Error; err != nil {
		return common.ErrDB(err)
	}

	// Remove product
	if err := db.
		Table(productmodel.Product{}.TableName()).
		Where("id = ?", productId).
		Delete(&productmodel.Product{}).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
