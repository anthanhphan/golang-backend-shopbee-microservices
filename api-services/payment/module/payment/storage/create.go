package paymentstorage

import (
	"context"
	"shopbee/common"
	paymentmodel "shopbee/module/payment/model"
)

func (s *paymentMySql) CreatePayment(
	ctx context.Context,
	data *paymentmodel.Payment,
) error {
	db := s.db.Table(paymentmodel.Payment{}.TableName())

	if err := db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
