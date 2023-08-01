package userbiz

import (
	"context"
	"shopbee/common"
	usermodel "shopbee/module/user/model"
)

type ListUserStorage interface {
	ListDataWithCondition(
		context context.Context,
		filter *usermodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]usermodel.User, error)
}

type listUserBiz struct {
	store     ListUserStorage
	requester common.Requester
}

func NewListUserBiz(store ListUserStorage, requester common.Requester) *listUserBiz {
	return &listUserBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *listUserBiz) ListUser(
	ctx context.Context,
	filter *usermodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]usermodel.User, error) {

	// if biz.requester.GetRole() != "admin" {
	// 	return nil, common.ErrNoPermission(nil)
	// }

	result, err := biz.store.ListDataWithCondition(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
