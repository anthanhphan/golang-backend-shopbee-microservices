package orderbiz

import (
	"context"
	"shopbee/common"
	ordermodel "shopbee/module/order/model"
)

type ViewOrderStorage interface {
	ViewOrder(
		ctx context.Context,
		userId int,
	) ([]ordermodel.Order, error)

	ViewOrderDetail(
		ctx context.Context,
		orderId int,
	) ([]ordermodel.OrderDetail, error)
}

type viewOrderBiz struct {
	store     ViewOrderStorage
	requester common.Requester
}

func NewViewOrderBiz(
	store ViewOrderStorage,
	requester common.Requester,
) *viewOrderBiz {
	return &viewOrderBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *viewOrderBiz) ViewOrder(
	ctx context.Context,
	userId int,
) ([]ordermodel.Order, error) {
	result, err := biz.store.ViewOrder(ctx, userId)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (biz *viewOrderBiz) ViewOrderDetail(
	ctx context.Context,
	orderId int,
) ([]ordermodel.OrderDetail, error) {
	result, err := biz.store.ViewOrderDetail(ctx, orderId)

	if err != nil {
		return nil, err
	}

	return result, nil
}
