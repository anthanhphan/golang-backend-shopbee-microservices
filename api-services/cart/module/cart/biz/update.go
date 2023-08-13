package cartbiz

import (
	"context"
	"shopbee/common"
	cartmodel "shopbee/module/cart/model"
)

type UpdateCartStorage interface {
	UpdateCart(
		ctx context.Context,
		data *cartmodel.CartUpdate,
	) error
}

type updateCartBiz struct {
	store     UpdateCartStorage
	requester common.Requester
}

func NewUpdateCartBiz(
	store UpdateCartStorage,
	requester common.Requester,
) *updateCartBiz {
	return &updateCartBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *updateCartBiz) UpdateCart(
	ctx context.Context,
	data *cartmodel.CartUpdate,
) error {
	if err := biz.store.UpdateCart(ctx, data); err != nil {
		return err
	}

	return nil
}
