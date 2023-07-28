package reqbiz

import (
	"context"
	"shopbee/common"
	reqmodel "shopbee/module/request/model"
)

type ListReportStorage interface {
	ListReportWithCondition(
		context context.Context,
		filter *reqmodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]reqmodel.RequestBanUser, error)
}

type listReportBiz struct {
	store     ListReportStorage
	requester common.Requester
}

func NewListReportBiz(store ListReportStorage, requester common.Requester) *listReportBiz {
	return &listReportBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *listReportBiz) ListReport(
	ctx context.Context,
	filter *reqmodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]reqmodel.RequestBanUser, error) {
	if biz.requester.GetRole() != "admin" {
		return nil, common.ErrNoPermission(nil)
	}

	result, err := biz.store.ListReportWithCondition(ctx, filter, paging, "User", "Shop")

	if err != nil {
		return nil, err
	}

	return result, nil
}
