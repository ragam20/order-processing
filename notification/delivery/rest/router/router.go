package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRoutes(g *gin.Engine) {
	api := g.Group("/api")

	internal := api.Group("/internal/notification")
	{
		internal.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "notification service - ok",
			})
		})
	}
}
