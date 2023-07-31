package productbiz

import (
	"context"
	"shopbee/common"
	productmodel "shopbee/module/product/model"
)

type ListProductStorage interface {
	ListDataWithCondition(
		context context.Context,
		filter *productmodel.Filter,
		paging *common.Paging,
		moreKey ...string,
	) ([]productmodel.Product, error)
}

type listProductBiz struct {
	store ListProductStorage
}

func NewListProductBiz(store ListProductStorage) *listProductBiz {
	return &listProductBiz{
		store: store,
	}
}

func (biz *listProductBiz) ListProduct(
	ctx context.Context,
	filter *productmodel.Filter,
	paging *common.Paging,
	moreKey ...string,
) ([]productmodel.Product, error) {
	result, err := biz.store.ListDataWithCondition(ctx, filter, paging, "Shop")

	if err != nil {
		return nil, err
	}

	return result, nil
}
