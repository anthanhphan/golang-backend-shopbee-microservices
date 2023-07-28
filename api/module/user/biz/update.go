package userbiz

import (
	"context"
	"shopbee/common"
	usermodel "shopbee/module/user/model"
)

type UpdateUserStorage interface {
	FindUserByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*usermodel.User, error)

	UpdateData(
		ctx context.Context,
		id int,
		data *usermodel.UserUpdate,
	) error
}

type updateUserBiz struct {
	store     UpdateUserStorage
	requester common.Requester
}

func NewUpdateUserBiz(store UpdateUserStorage, requester common.Requester) *updateUserBiz {
	return &updateUserBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *updateUserBiz) UpdateUser(
	ctx context.Context,
	id int,
	data *usermodel.UserUpdate,
) error {

	result, err := biz.store.FindUserByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if result.Status == 0 {
		return common.ErrEntityDeleted(usermodel.EntityName, nil)
	}

	if biz.requester.GetRole() != data.Role && biz.requester.GetRole() != "admin" {
		return common.ErrNoPermission(nil)
	}

	if biz.requester.GetRole() != "admin" && result.Id != biz.requester.GetUserId() {
		return common.ErrNoPermission(nil)
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return err
	}

	return nil
}
