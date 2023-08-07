package userbiz

import (
	"context"
	"shopbee/common"
	usermodel "shopbee/module/user/model"
)

type DeleteUserStorage interface {
	FindUserByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*usermodel.User, error)

	SoftDelete(
		ctx context.Context,
		id int,
	) error
}

type deleteUserBiz struct {
	store     DeleteUserStorage
	requester common.Requester
}

func NewDeleteUserBiz(store DeleteUserStorage, requester common.Requester) *deleteUserBiz {
	return &deleteUserBiz{store: store, requester: requester}
}

func (biz *deleteUserBiz) DeleteUser(ctx context.Context, id int) error {
	oldData, err := biz.store.FindUserByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrEntityNotFound(usermodel.EntityName, err)
	}

	if biz.requester.GetRole() != "admin" && biz.requester.GetUserId() != oldData.Id {
		return common.ErrNoPermission(nil)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(usermodel.EntityName, nil)
	}

	if err := biz.store.SoftDelete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(usermodel.EntityName, err)
	}

	return nil
}
