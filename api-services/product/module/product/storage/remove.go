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

	db.Exec("DELETE FROM wishlists WHERE product_id = ?", productId)
	db.Exec("DELETE FROM carts WHERE product_id = ?", productId)

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
