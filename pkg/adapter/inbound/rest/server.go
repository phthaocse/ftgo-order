package rest

import (
	"ftgo-order/pkg/core/service"
)

type Server interface {
	Run()
	InitRoute()
	InitMiddleware()
	InitBusinessService(businessService service.BusinessService)
}

func StartHTTPServer(server Server, services service.BusinessService) {
	server.InitBusinessService(services)
	server.InitMiddleware()
	server.InitRoute()
	server.Run()
}
