package service

type OrderService interface {
	createOrder()
	cancelOrder()
	reviseOrder()
}
