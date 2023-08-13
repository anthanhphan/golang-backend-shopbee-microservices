package cartbiz

import (
	"context"
	"shopbee/common"
)

type RemoveProductStorage interface {
	RemoveProductFromCart(
		ctx context.Context,
		productId int,
		userId int,
	) error
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
	userId int,
) error {

	if err := biz.store.RemoveProductFromCart(ctx, productId, userId); err != nil {
		return common.ErrCannotDeleteEntity("product", err)
	}

	return nil
}
