package handler

import (
	"net/http"
	"strconv"

	"ops/order/application/in"
	"ops/order/application/out"
	"ops/order/application/usecase"
	"ops/order/delivery/rest/request"
	"ops/order/delivery/rest/response"
	"ops/order/domain"
	"ops/shared/app"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	cancelOrderUC       *usecase.CancelOrderUseCase
	getDetailOrderUC    *usecase.GetDetailOrderUseCase
	getOrdersUC         *usecase.GetOrdersUseCase
	getOrdersByStatusUC *usecase.GetOrdersByStatusUseCase
	checkoutUC          *usecase.CheckoutUseCase
}

func NewHandler(cancelOrderUseCase *usecase.CancelOrderUseCase, getDetailOrderUseCase *usecase.GetDetailOrderUseCase, getOrdersUC *usecase.GetOrdersUseCase, getOrdersByStatusUC *usecase.GetOrdersByStatusUseCase, checkoutUC *usecase.CheckoutUseCase) *Handler {
	return &Handler{
		cancelOrderUC:       cancelOrderUseCase,
		getDetailOrderUC:    getDetailOrderUseCase,
		getOrdersUC:         getOrdersUC,
		getOrdersByStatusUC: getOrdersByStatusUC,
		checkoutUC:          checkoutUC,
	}
}

func (h *Handler) GetOrders(c *gin.Context) {
	var (
		orders []out.OrderWithItemsResult
		err    *app.AppError
	)
	if status := c.Query("status"); status != "" {
		orders, err = h.getOrdersByStatusUC.Execute(c.Request.Context(), domain.OrderStatus(status))
	} else {
		orders, err = h.getOrdersUC.Execute(c.Request.Context())
	}
	if err != nil {
		c.Error(err)
		return
	}

	resp := make([]response.OrderResponse, 0, len(orders))
	for _, orderWithItems := range orders {
		resp = append(resp, mapOrderResponse(orderWithItems))
	}

	c.JSON(http.StatusOK, &app.AppResponse[[]response.OrderResponse]{
		Success: true,
		Data:    &resp,
	})
}

func (h *Handler) GetOrderDetail(c *gin.Context) {
	orderID, err := strconv.ParseInt(c.Param("orderId"), 10, 64)
	if err != nil || orderID <= 0 {
		c.Error(&app.AppError{
			Code:    app.CodeInvalidRequest,
			Message: "order id must be a valid positive integer",
			Err:     err,
		})
		return
	}

	order, appErr := h.getDetailOrderUC.Execute(c.Request.Context(), orderID)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	resp := mapOrderResponse(*order)
	c.JSON(http.StatusOK, &app.AppResponse[response.OrderResponse]{
		Success: true,
		Data:    &resp,
	})
}

func mapOrderResponse(orderWithItems out.OrderWithItemsResult) response.OrderResponse {
	items := make([]response.OrderItemResponse, 0, len(orderWithItems.Items))
	for _, item := range orderWithItems.Items {
		items = append(items, response.OrderItemResponse{
			ID:        item.ID,
			ProductID: item.ProductID,
			Name:      item.Name,
			Price:     item.Price,
			Quantity:  item.Quantity,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	return response.OrderResponse{
		ID:          orderWithItems.Order.ID,
		Total:       orderWithItems.Order.Total,
		Status:      string(orderWithItems.Order.Status),
		CancelledAt: orderWithItems.Order.CancelledAt,
		CompletedAt: orderWithItems.Order.CompletedAt,
		CreatedAt:   orderWithItems.Order.CreatedAt,
		UpdatedAt:   orderWithItems.Order.UpdatedAt,
		Items:       items,
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
