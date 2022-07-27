package rest

import (
	"github.com/gin-gonic/gin"
)

func (gs *ginServer) createOrder(c *gin.Context) {
	gs.BusinessService.OrderService.CreateOrder()
}
