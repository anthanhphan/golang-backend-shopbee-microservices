package productbiz

import (
	"context"
	"shopbee/common"
	productmodel "shopbee/module/product/model"
)

type CreateProductStore interface {
	CreateProduct(ctx context.Context, data *productmodel.Product) error
}

type createProductBiz struct {
	store     CreateProductStore
	requester common.Requester
}

func NewCreateProductBiz(store CreateProductStore, requester common.Requester) *createProductBiz {
	return &createProductBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *createProductBiz) CreateProduct(
	context context.Context,
	data *productmodel.Product,
) error {
	// if err := data.Validate(); err != nil {
	// 	return common.ErrInvalidRequest(err)
	// }

	if biz.requester.GetRole() != "retailer" {
		return common.ErrNoPermission(nil)
	}

	if err := biz.store.CreateProduct(context, data); err != nil {
		return common.ErrCannotCreateEntity(productmodel.EntityName, err)
	}

	return nil
}
