package service

import (
	"fmt"
	"ftgo-order/pkg/interface/outbound/repository"
)

type OrderServiceI interface {
	CreateOrder()
	CancelOrder()
	ReviseOrder()
}

type OrderService struct {
	orderRepo repository.OrderRepo
}

func NewOrderService(orderRepo repository.OrderRepo) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

func (os *OrderService) CreateOrder() {
	fmt.Println("Implement me")
}

func (os *OrderService) CancelOrder() {
	fmt.Println("Implement me")
}

func (os *OrderService) ReviseOrder() {
	fmt.Println("Implement me")
}
