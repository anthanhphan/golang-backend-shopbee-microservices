package userbiz

import (
	"context"
	"errors"
	"shopbee/common"
	"shopbee/component/tokenprovider"
	usermodel "shopbee/module/user/model"
)

var (
	ErrEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrEmailOrPasswordInvalid",
	)
)

type LoginStorage interface {
	FindUserByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*usermodel.User, error)
}

type loginBiz struct {
	// appctx        appctx.AppContext
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBiz(
	storeUser LoginStorage,
	tokenProvider tokenprovider.Provider,
	hasher Hasher,
	expiry int,
) *loginBiz {
	return &loginBiz{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (biz *loginBiz) Login(
	ctx context.Context,
	data *usermodel.UserLogin,
) (*tokenprovider.Token, error) {
	user, err := biz.storeUser.FindUserByCondition(ctx, map[string]interface{}{
		"email": data.Email,
	})

	if err != nil {
		return nil, ErrEmailOrPasswordInvalid
	}

	passHashed := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, ErrEmailOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
	// refreshToken, err := biz.tokenProvider.Generate(payload, biz.expiry)

	// if err != nil {
	// 	return nil, common.ErrInternal(err)
	// }

	// account := usermodel.NewAccount(accessToken, refreshToken)

	// return account, nil
}

func (biz *loginBiz) LoginAdmin(
	ctx context.Context,
	data *usermodel.UserLogin,
) (*tokenprovider.Token, error) {
	user, err := biz.storeUser.FindUserByCondition(ctx, map[string]interface{}{
		"email": data.Email,
	})

	if err != nil {
		return nil, ErrEmailOrPasswordInvalid
	}

	passHashed := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, ErrEmailOrPasswordInvalid
	}

	if user.Role != "admin" {
		return nil, common.ErrNoPermission(nil)
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
