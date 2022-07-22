package rest

import (
	"ftgo-order/pkg/core/service"
	"github.com/gin-gonic/gin"
)

type ginServer struct {
	Engine *gin.Engine
	BusinessService
}

func NewGinServer() *ginServer {
	server := &ginServer{}
	server.Engine = gin.Default()
	return server
}

func (gs *ginServer) HandlerFnWrapper(fn service.BusinessServiceFn) gin.HandlerFunc {
	return func(c *gin.Context) {
		fn()
	}
}

func (gs *ginServer) InitOrderRoute() {
	orderGroup := gs.Engine.Group("order")
	{
		orderGroup.POST("", gs.HandlerFnWrapper(gs.OrderService.CreateOrder))
	}
}

func (gs *ginServer) InitRoute() {
	gs.InitOrderRoute()
}

func (gs *ginServer) InitBusinessService(services BusinessService) {
	gs.BusinessService = services
}

func (gs *ginServer) Run() {
	gs.Engine.Run(":8080")
}
