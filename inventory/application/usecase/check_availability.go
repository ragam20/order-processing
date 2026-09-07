package usecase

import (
	"context"

	"ops/inventory/application/in"
	"ops/inventory/application/out"
	"ops/inventory/domain"
)

type CheckAvailablity struct {
	repo domain.Repository
}

func NewCheckAvailablity(repo domain.Repository) *CheckAvailablity {
	return &CheckAvailablity{repo: repo}
}

func (u *CheckAvailablity) Execute(ctx context.Context, input []in.ItemRequest) ([]out.ItemAvailability, error) {
	return nil, nil
}
