package service

import (
	"order-api/src/app/order"
	"order-api/src/app/order/handler/response"
	"order-api/src/app/order/repository/record"
)

type service struct {
	repo order.Repository
}

// CreateOrder create new order from request
func (s *service) CreateOrder(data record.Order) error {
	return s.repo.CreateOrder(data)
}

// GetOrders list of orders
func (s *service) GetOrders() ([]response.Order, error) {
	return s.repo.GetOrders()
}

// UpdateOrder update order data and the related items
func (s *service) UpdateOrder(id int, data record.Order) error {
	if err := s.repo.UpdateOrder(id, data); err != nil {
		return err
	}
	return s.repo.UpdateOrderItems(id, data.Items)
}

// DeleteOrder delete order and the related items
func (s *service) DeleteOrder(id int) error {
	if err := s.repo.DeleteOrderItems(id); err != nil {
		return err
	}
	return s.repo.DeleteOrder(id)
}

func NewService(repo order.Repository) order.Service {
	return &service{repo}
}
