package reqbiz

import (
	"context"
	"shopbee/common"
	reqmodel "shopbee/module/request/model"
)

type RequestUpgradeStore interface {
	AcceptRequestUpgrade(
		ctx context.Context,
		data *reqmodel.RequestUpgrade,
	) error

	DenyRequestUpgrade(
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

func (biz *requestUpgradeBiz) AcceptRequestUpgrade(
	ctx context.Context,
	id int,
) error {
	if biz.requester.GetRole() != "admin" {
		return common.ErrNoPermission(nil)
	}

	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{
		"user_id": id,
	})

	if err != nil {
		return common.ErrInternal(err)
	}

	if err := biz.store.AcceptRequestUpgrade(ctx, data); err != nil {
		return common.ErrInternal(nil)
	}

	return nil
}
