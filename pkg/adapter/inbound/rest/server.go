package rest

import (
	postgresDb "ftgo-order/pkg/adapter/outbound/postgres_db"
	"ftgo-order/pkg/core/service"
	"ftgo-order/pkg/interface/outbound/repository"
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
	pgConn, _ := postgresDb.Init()
	orderRepo := repository.NewOrderRepo(pgConn)
	orderService := service.NewOrderService(orderRepo)
	services := BusinessService{
		OrderService: orderService,
	}
	server.InitBusinessService(services)
	server.InitMiddleware()
	server.InitRoute()
	server.Run()
}
