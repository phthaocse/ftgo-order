package service

import "fmt"

type OrderServiceI interface {
	CreateOrder()
	CancelOrder()
	ReviseOrder()
}

type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
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
