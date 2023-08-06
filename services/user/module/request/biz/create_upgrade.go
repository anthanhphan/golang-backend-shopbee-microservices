package reqbiz

import (
	"context"
	"shopbee/common"
	reqmodel "shopbee/module/request/model"
)

type RequestUpgradeStore interface {
	CreateRequestUpgrade(
		ctx context.Context,
		data *reqmodel.RequestUpgrade,
	) error

	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*reqmodel.RequestUpgrade, error)
}

type requestUpgradeBiz struct {
	store     RequestUpgradeStore
	requester common.Requester
}

func NewRequestUpgradeBiz(
	store RequestUpgradeStore,
	requester common.Requester,
) *requestUpgradeBiz {
	return &requestUpgradeBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *requestUpgradeBiz) CreateRequestUpgrade(
	ctx context.Context,
	data *reqmodel.RequestUpgrade,
) error {
	req, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{
		"user_id": data.UserId,
	})

	if req != nil {
		return common.ErrEntityExisted("request", err)
	}

	if err := biz.store.CreateRequestUpgrade(ctx, data); err != nil {
		return common.ErrCannotCreateEntity("request_upgrade", nil)
	}

	return nil
}
