package productbiz

import (
	"context"
	"shopbee/common"
	productmodel "shopbee/module/product/model"
)

type RemoveProductStorage interface {
	RemoveProduct(
		ctx context.Context,
		productId int,
	) error

	FindProduct(
		context context.Context,
		id int,
		moreKey ...string,
	) (*productmodel.Product, error)
}

type removeProductBiz struct {
	store     RemoveProductStorage
	requester common.Requester
}

func NewRemoveProductBiz(
	store RemoveProductStorage,
	requester common.Requester,
) *removeProductBiz {
	return &removeProductBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *removeProductBiz) RemoveProduct(
	ctx context.Context,
	productId int,
) error {
	product, err := biz.store.FindProduct(ctx, productId)

	if err != nil {
		return common.ErrEntityNotFound("product", err)
	}

	if product.ShopId != biz.requester.GetUserId() {
		return common.ErrNoPermission(nil)
	}

	if err := biz.store.RemoveProduct(ctx, productId); err != nil {
		return err
	}

	return nil
}
