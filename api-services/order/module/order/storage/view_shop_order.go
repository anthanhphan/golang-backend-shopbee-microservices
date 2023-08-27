package orderstorage

import (
	"context"
	ordermodel "shopbee/module/order/model"
)

func (s *orderMySql) ViewShopOrder(
	ctx context.Context,
	shopId int,
) ([]ordermodel.Order, error) {
	var result []ordermodel.Order

	db := s.db.Table(ordermodel.Order{}.TableName())

	db.Preload("Buyer")

	if err := db.Where("shop_id = ?", shopId).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (s *orderMySql) ViewShopOrderDetail(
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
