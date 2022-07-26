package service

type BusinessServiceFn func()

type BusinessService struct {
	OrderService OrderServiceI
}
