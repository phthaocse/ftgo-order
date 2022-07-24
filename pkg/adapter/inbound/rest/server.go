package rest

import (
	"ftgo-order/pkg/core/service"
)

type Server interface {
	Run()
	InitRoute()
	InitMiddleware()
	InitBusinessService(businessService BusinessService)
}

type BusinessService struct {
	OrderService service.OrderServiceI
}

func StartHTTPServer(server Server) {
	orderService := service.NewOrderService()
	services := BusinessService{
		OrderService: orderService,
	}
	server.InitBusinessService(services)
	server.InitMiddleware()
	server.InitRoute()
	server.Run()
}
