package orderstorage

import (
	"context"
	ordermodel "shopbee/module/order/model"
)

func (s *orderMySql) ChangeOrderStatus(
	ctx context.Context,
	orderId int,
	status string,
) error {
	db := s.db.Table(ordermodel.Order{}.TableName())

	if err := db.
		Where("id = ?", orderId).
		Updates(map[string]interface{}{
			"order_status": status,
		}).Error; err != nil {
		return err
	}

	return nil
}
