package core

import (
	"errors"
)

// ! Primary Port (order_service.go)

type OrderService interface {
	CreateOrder(order Order) error
}

type orderServiceImpl struct {
	r OrderRepository
}

func NewOrderService(repo OrderRepository) OrderService {
	return &orderServiceImpl{r: repo}
}

func (s *orderServiceImpl) CreateOrder(order Order) error {
	if order.Total <= 0 {
		return errors.New("total must be positive")
	}
	// Business logic...
	if err := s.r.Save(order); err != nil {
		return err
	}
	return nil
}
