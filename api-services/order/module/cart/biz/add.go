package cartbiz

import (
	"context"
	"shopbee/common"
	cartmodel "shopbee/module/cart/model"
)

type AddToCartStorage interface {
	AddToCart(
		ctx context.Context,
		data *cartmodel.CartCreate,
	) error
}

type addToCartBiz struct {
	store     AddToCartStorage
	requester common.Requester
}

func NewAddToCartBiz(
	store AddToCartStorage,
	requester common.Requester,
) *addToCartBiz {
	return &addToCartBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *addToCartBiz) AddToCart(
	ctx context.Context,
	data *cartmodel.CartCreate,
) error {
	if err := biz.store.AddToCart(ctx, data); err != nil {
		return common.ErrCannotCreateEntity("cart", err)
	}

	return nil
}
