package orderbiz

import (
	"context"
	"shopbee/common"
)

type ChangeOrderStatusStorage interface {
	ChangeOrderStatus(
		ctx context.Context,
		orderId int,
		status string,
	) error
}

type changeOrderStatusBiz struct {
	store     ChangeOrderStatusStorage
	requester common.Requester
}

func NewChangeOrderStatusBiz(
	store ChangeOrderStatusStorage,
	requester common.Requester,
) *changeOrderStatusBiz {
	return &changeOrderStatusBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *changeOrderStatusBiz) ChangeOrderStatus(
	ctx context.Context,
	orderId int,
	status string,
) error {

	if err := biz.store.ChangeOrderStatus(ctx, orderId, status); err != nil {
		return err
	}

	return nil
}
