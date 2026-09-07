package handler

import (
	"ops/inventory/application/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	getProductsUseCase    *usecase.GetProductsUseCase
	getProductByIDUseCase *usecase.GetProductByIDUseCase
}

func NewHandler(getProductsUseCase *usecase.GetProductsUseCase, getProductByIDUseCase *usecase.GetProductByIDUseCase) *Handler {
	return &Handler{
		getProductsUseCase:    getProductsUseCase,
		getProductByIDUseCase: getProductByIDUseCase,
	}
}

func (h *Handler) GetProducts(c *gin.Context) {
}

func (h *Handler) GetProductByID(c *gin.Context) {
}
