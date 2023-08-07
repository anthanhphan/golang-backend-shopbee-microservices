package reqbiz

import (
	"context"
	"shopbee/common"
	reqmodel "shopbee/module/request/model"
)

type ListRequetUpgradeStorage interface {
	ListUpgradeWithCondition(
		context context.Context,
		filter *reqmodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]reqmodel.RequestUpgrade, error)
}

type listRequetUpgradeBiz struct {
	store     ListRequetUpgradeStorage
	requester common.Requester
}

func NewListRequetUpgradeBiz(store ListRequetUpgradeStorage, requester common.Requester) *listRequetUpgradeBiz {
	return &listRequetUpgradeBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *listRequetUpgradeBiz) ListRequetUpgrade(
	ctx context.Context,
	filter *reqmodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]reqmodel.RequestUpgrade, error) {

	if biz.requester.GetRole() != "admin" {
		return nil, common.ErrNoPermission(nil)
	}

	result, err := biz.store.ListUpgradeWithCondition(ctx, filter, paging, "User")

	if err != nil {
		return nil, err
	}

	return result, nil
}
