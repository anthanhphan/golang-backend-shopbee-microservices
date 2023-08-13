package orderstorage

import (
	"context"
	ordermodel "shopbee/module/order/model"
)

func (s *orderMySql) ViewOrder(
	ctx context.Context,
	userId int,
) ([]ordermodel.Order, error) {
	var result []ordermodel.Order

	db := s.db.Table(ordermodel.Order{}.TableName())

	if err := db.Where("user_id = ?", userId).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (s *orderMySql) ViewOrderDetail(
	ctx context.Context,
	orderId int,
) ([]ordermodel.OrderDetail, error) {
	var result []ordermodel.OrderDetail

	db := s.db.Table(ordermodel.OrderDetail{}.TableName())

	if err := db.Where("order_id = ?", orderId).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
