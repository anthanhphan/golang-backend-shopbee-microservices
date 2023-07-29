package userbiz

import (
	"context"
	"fmt"
	"shopbee/common"
	"shopbee/component/hasher"
	mailservice "shopbee/module/sendmail"
	usermodel "shopbee/module/user/model"
)

type ForgotPWStorage interface {
	FindUserByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*usermodel.User, error)

	ForgotPassword(
		ctx context.Context,
		email string,
		data *usermodel.User,
	) error
}

type forgotPWUserBiz struct {
	store ForgotPWStorage
}

func NewForgotPWBiz(store ForgotPWStorage) *forgotPWUserBiz {
	return &forgotPWUserBiz{
		store: store,
	}
}

func (biz *forgotPWUserBiz) ForgotPassword(
	ctx context.Context,
	email string,
) error {
	data, err := biz.store.FindUserByCondition(ctx, map[string]interface{}{"email": email})

	if err != nil {
		return common.ErrInternal(err)
	}

	salt := common.GenSalt(50)
	password := common.GenSalt(10)

	md5 := hasher.NewMd5Hash()
	data.Password = md5.Hash(password + salt)
	data.Salt = salt

	if err := biz.store.ForgotPassword(ctx, email, data); err != nil {
		return err
	}

	fmt.Print(password)

	body := "Your new password is -> " + password
	mailservice.SendMail(data.Email, "Your new password", body)

	return nil
}
