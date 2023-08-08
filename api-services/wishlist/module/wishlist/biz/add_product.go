package wishlistbiz

import (
	"context"
	"shopbee/common"
	wishlistmodel "shopbee/module/wishlist/model"
)

type AddToWishListStorage interface {
	AddToWishList(
		ctx context.Context,
		data *wishlistmodel.WishList,
	) error
}

type addToWishListBiz struct {
	store     AddToWishListStorage
	requester common.Requester
}

func NewAddToWishListBiz(
	store AddToWishListStorage,
	requester common.Requester,
) *addToWishListBiz {
	return &addToWishListBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *addToWishListBiz) AddToWishList(
	ctx context.Context,
	productId int,
) error {
	wishList := wishlistmodel.WishList{
		UserId:    biz.requester.GetUserId(),
		ProductId: productId,
	}

	if err := biz.store.AddToWishList(ctx, &wishList); err != nil {
		return err
	}

	return nil
}
