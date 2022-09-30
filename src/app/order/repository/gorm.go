package repository

import (
	"errors"
	"order-api/src/app/order"
	"order-api/src/app/order/handler/response"
	"order-api/src/app/order/repository/record"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// CreateOrder create new order from request
func (repo *repository) CreateOrder(data record.Order) error {
	return repo.db.Create(&data).Error
}

// GetOrders get list of orders
func (repo *repository) GetOrders() ([]response.Order, error) {
	var records []record.Order
	if err := repo.db.Preload("Items").Find(&records).Error; err != nil {
		return nil, err
	}
	return record.MapToBatchResponse(records), nil
}

// UpdateOrder update order data by the given id
func (repo *repository) UpdateOrder(id int, data record.Order) error {
	query := repo.db.Omit("Items").Where("order_id", id).Updates(&data)
	if query.RowsAffected < 1 && query.Error == nil {
		return errors.New("record not found")
	}
	return query.Error
}

// UpdateOrderItems update items data that related with the given order id
func (repo *repository) UpdateOrderItems(orderID int, orderItems []record.Item) error {
	order := record.Order{OrderID: orderID}
	if err := repo.db.Model(&order).Association("Items").Replace(orderItems); err != nil {
		return err
	}
	return nil
}

// DeleteOrder delete order data by the given id
func (repo *repository) DeleteOrder(id int) error {
	query := repo.db.Delete(new(record.Order), "order_id", id)
	if query.RowsAffected < 1 && query.Error == nil {
		return errors.New("record not found")
	}
	return query.Error
}

// DeleteOrderItems delete items data that related with the given order id
func (repo *repository) DeleteOrderItems(orderID int) error {
	query := repo.db.Delete(new(record.Item), "order_id", orderID)
	if query.RowsAffected < 1 && query.Error == nil {
		return errors.New("record not found")
	}
	return query.Error
}

func NewGORMRepository(db *gorm.DB) order.Repository {
	return &repository{db}
}
