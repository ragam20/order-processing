package application

import (
	"context"

	"ops/order/application/in"
	"ops/order/application/out"
)

type InventoryService interface {
	CheckAvailabilityItem(ctx context.Context, input in.CheckoutInput) ([]out.ItemAvailability, error)
}
