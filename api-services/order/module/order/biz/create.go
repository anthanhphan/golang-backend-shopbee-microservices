package orderbiz

import (
	"context"
	"shopbee/common"
	ordermodel "shopbee/module/order/model"
)

type CreateOrderStorage interface {
	CreateOder(
		ctx context.Context,
		data *ordermodel.Order,
	) (int, error)

	CreateOderDetail(
		ctx context.Context,
		orderId int,
		data []map[string]interface{},
	) error
}

type createOrderBiz struct {
	store     CreateOrderStorage
	requester common.Requester
}

func NewCreateOrderBiz(
	store CreateOrderStorage,
	requester common.Requester,
) *createOrderBiz {
	return &createOrderBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *createOrderBiz) CreateOrder(
	ctx context.Context,
	data *ordermodel.Order,
	productList []map[string]interface{},
) error {
	id, err := biz.store.CreateOder(ctx, data)
	if err != nil {
		return err
	}

	if err := biz.store.CreateOderDetail(ctx, id, productList); err != nil {
		return err
	}

	return nil
}
