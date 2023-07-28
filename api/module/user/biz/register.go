package userbiz

import (
	"context"
	"shopbee/common"
	usermodel "shopbee/module/user/model"
)

type RegisterStorage interface {
	FindUserByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*usermodel.User, error)

	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBiz struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBiz(registerStorage RegisterStorage, hasher Hasher) *registerBiz {
	return &registerBiz{
		registerStorage: registerStorage,
		hasher:          hasher,
	}
}

func (biz *registerBiz) Register(
	ctx context.Context,
	data *usermodel.UserCreate,
) error {
	// Check email register existed status
	user, err := biz.registerStorage.FindUserByCondition(ctx, map[string]interface{}{
		"email": data.Email,
	})

	if user != nil {
		return common.ErrEntityExisted(usermodel.EntityName, err)
	}

	salt := common.GenSalt(50)
	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt

	if err := biz.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
