package productbiz

import (
	"context"
	productmodel "shopbee/module/product/model"
)

type ViewProductStorage interface {
	ViewProduct(
		context context.Context,
		id int,
		moreKey ...string,
	) (*productmodel.Product, error)
}

type viewProductBiz struct {
	store ViewProductStorage
}

func NewViewProductBiz(store ViewProductStorage) *viewProductBiz {
	return &viewProductBiz{
		store: store,
	}
}

func (biz *viewProductBiz) ViewProduct(
	ctx context.Context,
	id int,
	moreKey ...string,
) (*productmodel.Product, error) {
	product, err := biz.store.ViewProduct(ctx, id, "Shop", "Category")

	if err != nil {
		return nil, err
	}

	return product, nil
}
