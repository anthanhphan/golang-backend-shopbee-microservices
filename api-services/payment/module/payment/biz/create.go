package paymentbiz

import (
	"context"
	"shopbee/common"
	paymentmodel "shopbee/module/payment/model"
)

type CreatePaymentStorage interface {
	CreatePayment(
		ctx context.Context,
		data *paymentmodel.Payment,
	) error
}

type createPaymentBiz struct {
	store     CreatePaymentStorage
	requester common.Requester
}

func NewCreatePaymentBiz(
	store CreatePaymentStorage,
	requester common.Requester,
) *createPaymentBiz {
	return &createPaymentBiz{
		store:     store,
		requester: requester,
	}
}

func (biz *createPaymentBiz) CreatePayment(
	ctx context.Context,
	data *paymentmodel.Payment,
) (*paymentmodel.Payment, error) {
	if data.PaymenMethod == "card" {
		data.PaymenStatus = "paid"
	} else {
		data.PaymenStatus = "pending"
	}

	if err := biz.store.CreatePayment(ctx, data); err != nil {
		return nil, err
	}

	return data, nil
}
