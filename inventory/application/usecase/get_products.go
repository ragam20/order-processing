package usecase

import (
	"context"

	"ops/inventory/application/out"
	"ops/inventory/domain"
	"ops/shared/app"
)

type GetProductsUseCase struct {
	repo domain.Repository
}

const (
	defaultProductLimit = 10
	maxProductLimit     = 100
)

func NewGetProductsUseCase(repo domain.Repository) *GetProductsUseCase {
	return &GetProductsUseCase{repo: repo}
}

func (u *GetProductsUseCase) Execute(ctx context.Context, page int, limit int) (*out.GetProductsResult, *app.AppError) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = defaultProductLimit
	}
	if limit < 0 || limit > maxProductLimit {
		return nil, &app.AppError{
			Code:    app.CodeInvalidInput,
			Message: "limit must be between 1 and 100",
		}
	}
	if page < 0 {
		return nil, &app.AppError{
			Code:    app.CodeInvalidInput,
			Message: "page must be greater than 0",
		}
	}

	offset := (page - 1) * limit
	products, err := u.repo.GetProducts(ctx, limit+1, offset)
	if err != nil {
		return nil, &app.AppError{
			Code:    app.CodeInternalError,
			Message: "failed to get products",
			Err:     err,
		}
	}

	hasMore := len(products) > limit
	nextPage := page
	if hasMore {
		products = products[:limit]
		nextPage++
	}

	return &out.GetProductsResult{
		Products: products,
		Page:     page,
		Limit:    limit,
		NextPage: nextPage,
		HasMore:  hasMore,
	}, nil
}
