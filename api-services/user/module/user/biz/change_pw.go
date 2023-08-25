package userbiz

import (
	"context"
	"errors"
	"shopbee/common"
	usermodel "shopbee/module/user/model"
)

type ChangePasswordStorage interface {
	ChangePassword(
		ctx context.Context,
		id int,
		password string,
	) error

	FindUserByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*usermodel.User, error)
}

type changePasswordBiz struct {
	store     ChangePasswordStorage
	requester common.Requester
	hasher    Hasher
}

func NewChangePasswordBiz(
	store ChangePasswordStorage,
	requester common.Requester,
	hasher Hasher,
) *changePasswordBiz {
	return &changePasswordBiz{
		store:     store,
		requester: requester,
		hasher:    hasher,
	}
}

func (biz *changePasswordBiz) ChangePassword(
	ctx context.Context,
	id int,
	oldPass string,
	newPass string,
) error {
	user, err := biz.store.FindUserByCondition(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return err
	}

	passHashed := biz.hasher.Hash(oldPass + user.Salt)

	if user.Password != passHashed {
		return errors.New("your old password not correct")
	}

	newPass = biz.hasher.Hash(newPass + user.Salt)

	if err := biz.store.ChangePassword(ctx, id, newPass); err != nil {
		return err
	}

	return nil
}
