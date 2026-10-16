package handler

import (
	"net/http"

	"ops/payment/application/in"
	"ops/payment/application/usecase"
	"ops/payment/delivery/rest/request"
	"ops/payment/delivery/rest/response"
	"ops/shared/app"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	createPaymentUC    *usecase.CreatePaymentUseCase
	getStatusPaymentUC *usecase.GetPaymentStatus
}

func NewHandler(createPaymentUC *usecase.CreatePaymentUseCase, getStatusPaymentUC *usecase.GetPaymentStatus) *Handler {
	return &Handler{
		createPaymentUC:    createPaymentUC,
		getStatusPaymentUC: getStatusPaymentUC,
	}
}

func (h *Handler) CreatePayment(c *gin.Context) {
	var req request.CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, app.AppResponse[any]{
			Success: false,
			Error:   &app.ErrorInfo{Code: app.CodeInvalidRequest, Message: "invalid request body"},
		})
		return
	}

	items := make([]in.Item, 0, len(req.Items))
	for _, item := range req.Items {
		items = append(items, in.Item{
			ID:       item.ID,
			Name:     item.Name,
			Price:    item.Price,
			Quantity: item.Quantity,
		})
	}

	result, err := h.createPaymentUC.Execute(c.Request.Context(), in.CreatePaymentInput{
		OrderID:   req.OrderID,
		Amount:    req.TotalAmount,
		Items:     items,
		Customer:  in.Customer{ID: req.Customer.ID, Name: req.Customer.Name, Email: req.Customer.Email, Phone: req.Customer.Phone},
		FinishURL: req.FinishURL,
	})
	if err != nil {
		c.JSON(http.StatusBadGateway, app.AppResponse[any]{
			Success: false,
			Error:   &app.ErrorInfo{Code: err.Code, Message: err.Message},
		})
		return
	}

	c.JSON(http.StatusCreated, app.AppResponse[response.CreatePaymentResponse]{
		Success: true,
		Data: &response.CreatePaymentResponse{
			Status:                result.Status,
			ProviderReference:     result.ProviderReference,
			ProviderTransactionID: result.ProviderTransactionID,
			CheckoutURL:           result.CheckoutURL,
			PaymentToken:          result.PaymentToken,
		},
	})
}

func (h *Handler) GetStatusPayment(c *gin.Context) {
	providerReference := c.Param("providerReference")
	if providerReference == "" {
		c.JSON(http.StatusBadRequest, app.AppResponse[any]{
			Success: false,
			Error:   &app.ErrorInfo{Code: app.CodeInvalidRequest, Message: "provider reference is required"},
		})
		return
	}

	result, err := h.getStatusPaymentUC.Execute(c.Request.Context(), providerReference)
	if err != nil {
		c.JSON(http.StatusBadGateway, app.AppResponse[any]{
			Success: false,
			Error:   &app.ErrorInfo{Code: err.Code, Message: err.Message},
		})
		return
	}

	c.JSON(http.StatusOK, app.AppResponse[response.GetStatusPaymentResponse]{
		Success: true,
		Data: &response.GetStatusPaymentResponse{
			Status:                result.Status,
			Provider:              result.Provider,
			ProviderReference:     result.ProviderReference,
			ProviderTransactionID: result.ProviderTransactionID,
			Amount:                result.Amount,
			PaidAt:                result.PaidAt,
		},
	})
}
