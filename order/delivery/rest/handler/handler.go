package handler

import (
	"net/http"

	"ops/order/application/in"
	"ops/order/application/out"
	"ops/order/application/usecase"
	"ops/order/delivery/rest/request"
	"ops/order/delivery/rest/response"
	"ops/shared/app"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	cancelOrderUC    *usecase.CancelOrderUseCase
	getDetailOrderUC *usecase.GetDetailOrderUseCase
	checkoutUC       *usecase.CheckoutUseCase
}

func NewHandler(cancelOrderUseCase *usecase.CancelOrderUseCase, getDetailOrderUseCase *usecase.GetDetailOrderUseCase, checkoutUC *usecase.CheckoutUseCase) *Handler {
	return &Handler{
		cancelOrderUC:    cancelOrderUseCase,
		getDetailOrderUC: getDetailOrderUseCase,
		checkoutUC:       checkoutUC,
	}
}

func (h *Handler) Checkout(c *gin.Context) {
	var req request.CheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(&app.AppError{
			Code:    app.CodeInvalidRequest,
			Message: "invalid request body",
			Err:     err,
		})
		return
	}

	ckInputItems := make([]in.CheckoutItemInput, 0)
	for _, v := range req.Items {
		ckInputItems = append(ckInputItems, in.CheckoutItemInput{
			ProductID: v.ProductID,
			Quantity:  v.Quantity,
		})
	}

	ckInput := in.CheckoutInput{
		Items: ckInputItems,
		Customer: in.CustomerInput{
			ID:    req.Customer.ID,
			Name:  req.Customer.Name,
			Email: req.Customer.Email,
			Phone: req.Customer.Phone,
		},
		FinishURL: req.FinishURL,
	}

	res, err := h.checkoutUC.Execute(c.Request.Context(), ckInput)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, &app.AppResponse[response.CheckoutResponse]{
		Success: true,
		Data: &response.CheckoutResponse{
			OrderID: res.OrderID,
			Status:  res.Status,
			Total:   res.Total,
			Items: func(items []out.CheckoutItemResult) []response.CheckoutItemResponse {
				r := make([]response.CheckoutItemResponse, 0)
				for _, item := range items {
					r = append(r, response.CheckoutItemResponse{
						ProductID: item.ProductID,
						Name:      item.Name,
						Price:     item.Price,
						Quantity:  item.Quantity,
					})
				}
				return r
			}(res.Items),
			Payment: response.OrderPaymentResponse{
				Status:       res.Payment.Status,
				CheckoutURL:  res.Payment.CheckoutURL,
				PaymentToken: res.Payment.PaymentToken,
			},
		},
	})
}
