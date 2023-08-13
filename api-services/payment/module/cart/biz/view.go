package cartbiz

import (
	"context"
	"shopbee/common"
)

type ViewMyCartStorage interface {
	ViewMyCart(
		ctx context.Context,
		userId int,
	) ([]map[string]interface{}, error)
}

type viewMyCartBiz struct {
	store     ViewMyCartStorage
	requester common.Requester
}

func NewViewMyCartBiz(
	store ViewMyCartStorage,
	requester common.Requester,
) *viewMyCartBiz {
	return &viewMyCartBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *viewMyCartBiz) ViewMyCart(
	ctx context.Context,
	userId int,
) ([]map[string]interface{}, error) {

	cart, err := biz.store.ViewMyCart(ctx, userId)
	if err != nil {
		return nil, err
	}

	return cart, nil
}
