package application

import "context"

type InventoryService interface {
	CheckavailabilityItem(ctx context.Context)
}
