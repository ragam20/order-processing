package router

import (
	"ops/payment/delivery/rest/handler"

	"github.com/gin-gonic/gin"
)

func NewWebhooks(g *gin.Engine, w *handler.WebhookHandler) {
	api := g.Group("/api/webhooks")

	midtrans := api.Group("/midtrans")
	{
		midtrans.POST("/dev/notification", w.MidtransNotification)
	}
}
