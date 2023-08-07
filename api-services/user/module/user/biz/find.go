package userbiz

import (
	"context"
	usermodel "shopbee/module/user/model"
)

type FindUserStorage interface {
	FindUserByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*usermodel.User, error)
}

type findUserBiz struct {
	storeUser LoginStorage
}

func NewFindUserBiz(store LoginStorage) *findUserBiz {
	return &findUserBiz{
		storeUser: store,
	}
}

func (biz *findUserBiz) FindUserById(
	ctx context.Context,
	id int,
) (*usermodel.User, error) {
	user, err := biz.storeUser.FindUserByCondition(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return nil, ErrEmailOrPasswordInvalid
	}

	return user, nil
}
