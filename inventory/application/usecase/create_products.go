package usecase

import (
	"context"
	"strings"
	"time"

	"ops/inventory/application/in"
	"ops/inventory/domain"
	"ops/shared/app"
)

type CreateProductsUseCase struct {
	repo domain.Repository
}

func NewCreateProductsUseCase(repo domain.Repository) *CreateProductsUseCase {
	return &CreateProductsUseCase{repo: repo}
}

func (u *CreateProductsUseCase) Execute(ctx context.Context, input []in.CreateProductInput) ([]domain.Product, *app.AppError) {
	if len(input) == 0 {
		return nil, &app.AppError{
			Code:    app.CodeInvalidInput,
			Message: "items must not be empty",
		}
	}

	now := time.Now().UTC()
	products := make([]domain.Product, 0, len(input))
	for _, item := range input {
		if strings.TrimSpace(item.Name) == "" {
			return nil, &app.AppError{
				Code:    app.CodeInvalidInput,
				Message: domain.ErrProductName.Error(),
				Err:     domain.ErrProductName,
			}
		}
		if item.Price < 0 {
			return nil, &app.AppError{
				Code:    app.CodeInvalidInput,
				Message: domain.ErrProductPrice.Error(),
				Err:     domain.ErrProductPrice,
			}
		}
		if item.Available < 0 {
			return nil, &app.AppError{
				Code:    app.CodeInvalidInput,
				Message: domain.ErrProductAvailable.Error(),
				Err:     domain.ErrProductAvailable,
			}
		}

		products = append(products, domain.Product{
			Name:      strings.TrimSpace(item.Name),
			ImageURL:  item.ImageURL,
			Price:     item.Price,
			Available: item.Available,
			CreatedAt: now,
			UpdatedAt: now,
		})
	}

	if err := u.repo.CreateBulk(ctx, products); err != nil {
		return nil, &app.AppError{
			Code:    app.CodeInternalError,
			Message: "failed to create products",
			Err:     err,
		}
	}

	return products, nil
}
