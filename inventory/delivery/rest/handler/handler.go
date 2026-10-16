package handler

import (
	"net/http"
	"strconv"

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
	createProductsUC    *usecase.CreateProductsUseCase
}

func NewHandler(getProductsUC *usecase.GetProductsUseCase, getProductByIDUC *usecase.GetProductByIDUseCase, checkAvailabilityUC *usecase.CheckAvailabilityUseCase, createProductsUC *usecase.CreateProductsUseCase) *Handler {
	return &Handler{
		getProductsUC:       getProductsUC,
		getProductByIDUC:    getProductByIDUC,
		checkAvailabilityUC: checkAvailabilityUC,
		createProductsUC:    createProductsUC,
	}
}

func (h *Handler) CreateProducts(c *gin.Context) {
	var req request.CreateProductsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(&app.AppError{
			Code:    app.CodeInvalidRequest,
			Message: "invalid request body",
			Err:     err,
		})
		return
	}

	input := make([]in.CreateProductInput, 0, len(req.Items))
	for _, item := range req.Items {
		input = append(input, in.CreateProductInput{
			Name:      item.Name,
			ImageURL:  item.ImageURL,
			Price:     item.Price,
			Available: item.Available,
		})
	}

	products, err := h.createProductsUC.Execute(c.Request.Context(), input)
	if err != nil {
		c.Error(err)
		return
	}

	resp := make([]response.ProductResponse, 0, len(products))
	for _, product := range products {
		resp = append(resp, response.ProductResponse{
			ID:        product.ID,
			Name:      product.Name,
			ImageURL:  product.ImageURL,
			Price:     product.Price,
			Available: product.Available,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		})
	}

	c.JSON(http.StatusCreated, &app.AppResponse[[]response.ProductResponse]{
		Success: true,
		Data:    &resp,
	})
}

func (h *Handler) GetProducts(c *gin.Context) {
	page, err := parseQueryInt(c, "page")
	if err != nil {
		c.Error(&app.AppError{
			Code:    app.CodeInvalidRequest,
			Message: "page must be a valid integer",
			Err:     err,
		})
		return
	}

	limit, err := parseQueryInt(c, "limit")
	if err != nil {
		c.Error(&app.AppError{
			Code:    app.CodeInvalidRequest,
			Message: "limit must be a valid integer",
			Err:     err,
		})
		return
	}

	result, appErr := h.getProductsUC.Execute(c.Request.Context(), page, limit)
	if appErr != nil {
		c.Error(appErr)
		return
	}

	products := make([]response.ProductResponse, 0, len(result.Products))
	for _, product := range result.Products {
		products = append(products, response.ProductResponse{
			ID:        product.ID,
			Name:      product.Name,
			ImageURL:  product.ImageURL,
			Price:     product.Price,
			Available: product.Available,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, &app.AppResponse[[]response.ProductResponse]{
		Success: true,
		Data:    &products,
		Meta: map[string]any{
			"page":     result.Page,
			"limit":    result.Limit,
			"nextPage": result.NextPage,
			"hasMore":  result.HasMore,
		},
	})
}

func parseQueryInt(c *gin.Context, key string) (int, error) {
	value := c.Query(key)
	if value == "" {
		return 0, nil
	}

	return strconv.Atoi(value)
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
