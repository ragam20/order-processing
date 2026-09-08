package handler

import (
	"net/http"

	"ops/inventory/application/in"
	"ops/inventory/application/usecase"
	"ops/inventory/delivery/rest/request"
	"ops/inventory/delivery/rest/response"
	"ops/shared/app"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	getProductsUC       *usecase.GetProductsUseCase
	getProductByIDUC    *usecase.GetProductByIDUseCase
	checkAvailabilityUC *usecase.CheckAvailabilityUseCase
}

func NewHandler(getProductsUC *usecase.GetProductsUseCase, getProductByIDUC *usecase.GetProductByIDUseCase, checkAvailabilityUC *usecase.CheckAvailabilityUseCase) *Handler {
	return &Handler{
		getProductsUC:       getProductsUC,
		getProductByIDUC:    getProductByIDUC,
		checkAvailabilityUC: checkAvailabilityUC,
	}
}

func (h *Handler) GetProducts(c *gin.Context) {
}

func (h *Handler) GetProductByID(c *gin.Context) {
}

func (h *Handler) CheckAvailabilityItems(c *gin.Context) {
	var req request.CheckAvailabilityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(&app.AppError{
			Code:    app.CodeInvalidRequest,
			Message: "invalid request body",
			Err:     err,
		})
		return
	}

	input := make([]in.ItemRequest, 0, len(req.Items))
	for _, v := range req.Items {
		input = append(input, in.ItemRequest{
			ProductID: v.ProductID,
			Quantity:  v.Quantity,
		})
	}

	result, err := h.checkAvailabilityUC.Execute(c.Request.Context(), input)
	if err != nil {
		c.Error(err)
		return
	}

	resp := make([]response.ItemAvailabilityResponse, 0, len(result))
	for _, v := range result {
		resp = append(resp, response.ItemAvailabilityResponse{
			ProductID:   v.ProductID,
			Name:        v.Name,
			Price:       v.Price,
			Quantity:    v.Quantity,
			Available:   v.Available,
			IsAvailable: v.IsAvailable,
		})
	}

	c.JSON(http.StatusOK, &app.AppResponse[[]response.ItemAvailabilityResponse]{
		Success: true,
		Data:    &resp,
	})
}
