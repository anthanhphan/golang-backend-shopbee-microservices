package orderstorage

import (
	"context"
	"fmt"
	"shopbee/common"
	ordermodel "shopbee/module/order/model"
)

func (s *orderMySql) CreateOder(
	ctx context.Context,
	data *ordermodel.Order,
) (int, error) {
	db := s.db.Table(ordermodel.Order{}.TableName())

	if err := db.Create(&data).Error; err != nil {
		return 0, common.ErrDB(err)
	}

	return data.Id, nil
}

func (s *orderMySql) CreateOderDetail(
	ctx context.Context,
	orderId int,
	data []common.Product,
) error {

	db := s.db.Table(ordermodel.OrderDetail{}.TableName())

	for i := range data {

		orderDetail := ordermodel.OrderDetail{
			OrderId:       orderId,
			ProductOrigin: &data[i],
		}
		fmt.Println(orderDetail)
		if err := db.Create(&orderDetail).Error; err != nil {
			return err
		}
	}

	return nil
}
