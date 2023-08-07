package categorybiz

import (
	"context"
	"shopbee/common"
	categorymodel "shopbee/module/category/model"
)

type ListCategoryStorage interface {
	ListDataWithCondition(
		context context.Context,
		paging *common.Paging,
		moreKey ...string,
	) ([]categorymodel.Category, error)
}

type listCategoryBiz struct {
	store ListCategoryStorage
}

func NewListReportBiz(store ListCategoryStorage) *listCategoryBiz {
	return &listCategoryBiz{
		store: store,
	}
}

func (biz *listCategoryBiz) ListCategory(
	ctx context.Context,
	paging *common.Paging,
	moreKey ...string,
) ([]categorymodel.Category, error) {

	result, err := biz.store.ListDataWithCondition(ctx, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
