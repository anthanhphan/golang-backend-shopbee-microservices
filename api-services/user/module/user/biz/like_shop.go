package userbiz

import (
	"context"
	"shopbee/common"
)

type LikeShopStorage interface {
	LikeShop(
		ctx context.Context,
		userId int,
		shopId int,
	) error

	DislikeShop(
		ctx context.Context,
		userId int,
		shopId int,
	) error

	IsLiked(
		ctx context.Context,
		userId int,
		shopId int,
	) bool
}

type likeShopBiz struct {
	store     LikeShopStorage
	requester common.Requester
}

func NewLikeShopBiz(
	store LikeShopStorage,
	requester common.Requester,

) *likeShopBiz {
	return &likeShopBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *likeShopBiz) LikeShop(
	ctx context.Context,
	userId int,
	shopId int,
) error {
	if err := biz.store.LikeShop(ctx, userId, shopId); err != nil {
		return err
	}

	return nil
}

func (biz *likeShopBiz) DislikeShop(
	ctx context.Context,
	userId int,
	shopId int,
) error {
	if err := biz.store.DislikeShop(ctx, userId, shopId); err != nil {
		return err
	}

	return nil
}

func (biz *likeShopBiz) IsLikedShop(
	ctx context.Context,
	userId int,
	shopId int,
) bool {
	if err := biz.store.IsLiked(ctx, userId, shopId); !err {
		return false
	}

	return true
}
