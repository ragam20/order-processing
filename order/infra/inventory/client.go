package inventory

import (
	"context"

	"ops/order/application"
)

type Client struct{}

func New() application.InventoryService {
	return &Client{}
}

func (c *Client) CheckavailabilityItem(ctx context.Context) {}
