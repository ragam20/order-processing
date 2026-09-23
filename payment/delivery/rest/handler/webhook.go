package handler

import (
	"io"
	"net/http"

	"ops/payment/application"
	"ops/payment/application/in"
	"ops/payment/application/usecase"

	"github.com/gin-gonic/gin"
)

type WebhookHandler struct {
	processStatusPayment *usecase.ProcessStatusPaymentUseCase
	paymentProvider      application.PaymentProvider
}

func NewWebHookHandler(
	processStatusPayment *usecase.ProcessStatusPaymentUseCase,
	paymentProvider application.PaymentProvider,
) *WebhookHandler {
	return &WebhookHandler{
		processStatusPayment: processStatusPayment,
		paymentProvider:      paymentProvider,
	}
}

func (h *WebhookHandler) MidtransNotification(c *gin.Context) {
	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	notification, err := h.paymentProvider.ParseWebhook(in.ParseWebhook{
		Payload: payload,
	})
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = h.processStatusPayment.Execute(c.Request.Context(), in.ProcessPaymentInput{
		Provider:              notification.Provider,
		ProviderReference:     notification.ProviderReference,
		ProviderTransactionID: notification.ProviderTransactionID,
		Status:                notification.Status,
		GrossAmount:           notification.GrossAmount,
		PaidAt:                notification.PaidAt,
		RawPayload:            notification.RawPayload,
	})
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
