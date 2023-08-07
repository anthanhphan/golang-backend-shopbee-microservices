package reqbiz

import (
	"context"
	"shopbee/common"
	reqmodel "shopbee/module/request/model"
)

type CreateRequestBanUserStore interface {
	CreateRequestBanUser(
		ctx context.Context,
		data *reqmodel.RequestBanUser,
	) error
}

type createRequestBanUserBiz struct {
	store     CreateRequestBanUserStore
	requester common.Requester
}

func NewCreateRequestBanUserBiz(
	store CreateRequestBanUserStore,
	requester common.Requester,
) *createRequestBanUserBiz {
	return &createRequestBanUserBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *createRequestBanUserBiz) CreateRequestBanUser(
	ctx context.Context,
	data *reqmodel.RequestBanUser,
) error {

	if err := biz.store.CreateRequestBanUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity("request_ban_account", nil)
	}

	return nil
}
