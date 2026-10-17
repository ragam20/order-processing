package inventory

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"ops/order/application"
	"ops/order/application/in"
	"ops/order/application/out"
	"ops/shared/config"
)

type inventoryClient struct {
	env *config.Env
}

func New(env *config.Env) application.InventoryService {
	return &inventoryClient{env: env}
}

func (c *inventoryClient) CheckAvailabilityItem(ctx context.Context, input in.CheckoutInput) ([]out.AvailabilityItems, error) {
	itemReq := CheckAvailabilityRequest{
		Items: make([]ItemRequest, 0, len(input.Items)),
	}
	for _, v := range input.Items {
		itemReq.Items = append(itemReq.Items, ItemRequest{
			ProductID: v.ProductID,
			Quantity:  v.Quantity,
		})
	}

	body, err := json.Marshal(itemReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.env.HTTP.InventoryURL+"/api/internal/inventory/check-availability",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"inventory service returned status: %d",
			resp.StatusCode,
		)
	}

	var res CheckAvailabilityResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	result := make([]out.AvailabilityItems, 0, len(res.Data))
	for _, item := range res.Data {
		result = append(result, out.AvailabilityItems{
			ProductID:   item.ProductID,
			Name:        item.Name,
			Quantity:    item.Quantity,
			Price:       item.Price,
			Available:   item.Available,
			IsAvailable: item.IsAvailable,
		})
	}

	return result, nil
}
