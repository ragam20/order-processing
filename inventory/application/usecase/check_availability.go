package usecase

import (
	"context"

	"ops/inventory/application/in"
	"ops/inventory/application/out"
	"ops/inventory/domain"
	"ops/shared/app"
)

type CheckAvailabilityUseCase struct {
	repo domain.Repository
}

func NewCheckAvailability(repo domain.Repository) *CheckAvailabilityUseCase {
	return &CheckAvailabilityUseCase{repo: repo}
}

func (u *CheckAvailabilityUseCase) Execute(ctx context.Context, input []in.ItemRequest) ([]out.ItemAvailability, *app.AppError) {
	if len(input) == 0 {
		return nil, &app.AppError{
			Code:    app.CodeInvalidInput,
			Message: "items must not be empty",
		}
	}

	ids := make([]int64, 0, len(input))
	seen := make(map[int64]struct{}, len(input))

	for _, item := range input {
		if item.Quantity <= 0 {
			return nil, &app.AppError{
				Code:    app.CodeInvalidInput,
				Message: domain.ErrProductQuantity.Error(),
				Err:     domain.ErrProductQuantity,
			}
		}
		if _, exists := seen[item.ProductID]; exists {
			return nil, &app.AppError{
				Code:    app.CodeInvalidInput,
				Message: "duplicate product id",
			}
		}

		seen[item.ProductID] = struct{}{}
		ids = append(ids, item.ProductID)
	}

	products, err := u.repo.GetProductByIDs(ctx, ids)
	if err != nil {
		return nil, &app.AppError{
			Code:    app.CodeInternalError,
			Message: "internal error",
			Err:     err,
		}
	}

	productsMap := make(map[int64]domain.Product)
	for _, item := range products {
		productsMap[item.ID] = item
	}

	result := make([]out.ItemAvailability, 0, len(input))

	for _, item := range input {
		product, found := productsMap[item.ProductID]
		if !found {
			return nil, &app.AppError{
				Code:    app.CodeNotFound,
				Message: domain.ErrProductNotFound.Error(),
				Err:     domain.ErrProductNotFound,
			}
		}

		result = append(result, out.ItemAvailability{
			ProductID:   product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Quantity:    item.Quantity,
			Available:   product.Available,
			IsAvailable: product.Available >= int(item.Quantity),
		})
	}

	return result, nil
}
