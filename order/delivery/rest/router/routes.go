package router

import (
	"net/http"

	"ops/order/delivery/rest/handler"

	"github.com/gin-gonic/gin"
)

func NewRoutes(g *gin.Engine, h *handler.Handler) {
	api := g.Group("/api")
	{
		api.POST("/checkout", h.Checkout)
	}

	internal := api.Group("/internal/order")
	{
		internal.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "order service - ok",
			})
		})
	}
}
