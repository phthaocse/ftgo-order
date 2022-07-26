package rest

import (
	"ftgo-order/pkg/core/service"
	"github.com/gin-gonic/gin"
)

type ginServer struct {
	engine *gin.Engine
	service.BusinessService
}

func NewGinServer() *ginServer {
	server := &ginServer{}
	server.engine = gin.Default()
	return server
}

func (gs *ginServer) HandlerFnWrapper(fn service.BusinessServiceFn) gin.HandlerFunc {
	return func(c *gin.Context) {
		fn()
	}
}

func (gs *ginServer) InitOrderRoute() {
	orderGroup := gs.engine.Group("order")
	{
		orderGroup.POST("", gs.HandlerFnWrapper(gs.OrderService.CreateOrder))
	}
}

func (gs *ginServer) InitRoute() {
	gs.InitOrderRoute()
}

func (gs *ginServer) InitBusinessService(services service.BusinessService) {
	gs.BusinessService = services
}

func (gs *ginServer) InitMiddleware() {
	gs.engine.Use(authMiddleware)
}

func (gs *ginServer) Run() {
	gs.engine.Run(":8080")
}
