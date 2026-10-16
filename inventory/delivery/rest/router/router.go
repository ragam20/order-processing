package router

import (
	"net/http"

	"ops/inventory/delivery/rest/handler"

	"github.com/gin-gonic/gin"
)

func NewRoutes(g *gin.Engine, h *handler.Handler) {
	api := g.Group("/api")

	internal := api.Group("/internal/inventory")
	{
		internal.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "inventory service - ok",
			})
		})

		internal.POST("/check-availability", h.CheckAvailabilityItems)
	}

	api.POST("/products/bulk", h.CreateProducts)
}
