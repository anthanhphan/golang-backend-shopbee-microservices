package userbiz

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"shopbee/common"
	"shopbee/component/hasher"
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

func sendMailRequest(email, password string) {
	url := "http://13.54.238.78/api/v1/mail/send"
	type Mail struct {
		Receiver string `json:"receiver"`
		Subject  string `json:"subject"`
		Body     string `json:"body"`
	}
	mail := Mail{
		Receiver: email,
		Subject:  "New Shopbee Password",
		Body:     "Your new password is " + password,
	}
	payload := mail

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling payload:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadJSON))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)

	// Read and display the response body
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Println("Response Body:", buf.String())
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
	fmt.Print(password)
	sendMailRequest(email, password)
	return nil
}
