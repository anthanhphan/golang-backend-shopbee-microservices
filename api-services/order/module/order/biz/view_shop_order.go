package orderbiz

import (
	"context"
	"shopbee/common"
	ordermodel "shopbee/module/order/model"
)

type ViewShopOrderStorage interface {
	ViewShopOrder(
		ctx context.Context,
		shopId int,
	) ([]ordermodel.Order, error)

	ViewShopOrderDetail(
		ctx context.Context,
		orderId int,
	) ([]ordermodel.OrderDetail, error)
}

type viewShopOrderBiz struct {
	store     ViewShopOrderStorage
	requester common.Requester
}

func NewViewShopOrderBiz(
	store ViewShopOrderStorage,
	requester common.Requester,
) *viewShopOrderBiz {
	return &viewShopOrderBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *viewShopOrderBiz) ViewShopOrder(
	ctx context.Context,
	shopId int,
) ([]ordermodel.Order, error) {
	result, err := biz.store.ViewShopOrder(ctx, shopId)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (biz *viewShopOrderBiz) ViewShopOrderDetail(
	ctx context.Context,
	orderId int,
) ([]ordermodel.OrderDetail, error) {
	result, err := biz.store.ViewShopOrderDetail(ctx, orderId)

	if err != nil {
		return nil, err
	}

	return result, nil
}
