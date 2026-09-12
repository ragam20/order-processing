package router

import (
	"net/http"

	"ops/payment/delivery/rest/handler"

	"github.com/gin-gonic/gin"
)

func NewRoutes(g *gin.Engine, h *handler.Handler) {
	api := g.Group("/api")

	internal := api.Group("/internal/payment")
	{
		internal.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "payment service - ok",
			})
		})
		internal.POST("/", h.CreatePayment)
		internal.GET("/status/:providerReference", h.GetStatusPayment)
	}
}
