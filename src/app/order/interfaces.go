package order

import (
	"order-api/src/app/order/repository/record"
)

type Repository interface {
	CreateOrder(data record.Order) error
	GetOrders() ([]record.Order, error)
	UpdateOrder(id int, data record.Order) error
	UpdateOrderItems(orderID int, orderItems []record.Item) error
	DeleteOrder(id int) error
	DeleteOrderItems(orderID int) error
}

type Service interface {
	CreateOrder(data record.Order) error
	GetOrders() ([]record.Order, error)
	UpdateOrder(id int, data record.Order) error
	DeleteOrder(id int) error
}
