package wishlistbiz

import (
	"context"
	"shopbee/common"
	wishlistmodel "shopbee/module/wishlist/model"
)

type ViewMyWishListStorage interface {
	ViewMyWishList(
		context context.Context,
		id int,
	) ([]wishlistmodel.WishList, error)
}

type viewMyWishListBiz struct {
	store     ViewMyWishListStorage
	requester common.Requester
}

func NewViewMyWishListBiz(
	store ViewMyWishListStorage,
	requester common.Requester,
) *viewMyWishListBiz {
	return &viewMyWishListBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *viewMyWishListBiz) ViewMyWishList(
	context context.Context,
	id int,
) ([]wishlistmodel.WishList, error) {

	wishlist, err := biz.store.ViewMyWishList(context, id)

	if err != nil {
		return nil, err
	}

	return wishlist, nil
}
