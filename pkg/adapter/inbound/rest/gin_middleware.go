package rest

import "github.com/gin-gonic/gin"

func authMiddleware(c *gin.Context) {
	c.Next()
}
