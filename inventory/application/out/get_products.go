package out

import "ops/inventory/domain"

type GetProductsResult struct {
	Products []domain.Product
	Page     int
	Limit    int
	NextPage int
	HasMore  bool
}
