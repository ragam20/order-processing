package database

import (
	"context"

	"ops/inventory/domain"
	"ops/inventory/infra/database/entity"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, product *domain.Product) error {
	if err := r.db.WithContext(ctx).Create(entity.ProductEntity{
		ID:        product.ID,
		Name:      product.Name,
		Available: product.Available,
		ImageURL:  product.ImageURL,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) GetProductByIDs(ctx context.Context, ids []int64) ([]domain.Product, error) {
	var e []entity.ProductEntity

	if err := r.db.WithContext(ctx).
		Where("id IN ?", ids).Find(&e).Error; err != nil {
		return nil, err
	}

	d := make([]domain.Product, 0, len(e))
	for _, v := range e {
		d = append(d, domain.Product{
			ID:        v.ID,
			Name:      v.Name,
			Available: v.Available,
			ImageURL:  v.ImageURL,
			Price:     v.Price,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return d, nil
}

func (r *repository) Update(ctx context.Context, product *domain.Product) error {
	e := entity.ProductEntity{
		ID:        product.ID,
		Name:      product.Name,
		Available: product.Available,
		ImageURL:  product.ImageURL,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
	return r.db.WithContext(ctx).Model(&entity.ProductEntity{}).Updates(e).Error
}

func (r *repository) GetProducts(ctx context.Context) []domain.Product {
	var e []entity.ProductEntity

	if err := r.db.Find(&e).Error; err != nil {
		return []domain.Product{}
	}

	p := make([]domain.Product, 0)
	for _, v := range e {
		p = append(p, domain.Product{
			ID:        v.ID,
			Name:      v.Name,
			Available: v.Available,
			ImageURL:  v.ImageURL,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return p
}

func (r *repository) GetProductByID(ctx context.Context, productID int64) (*domain.Product, error) {
	var e entity.ProductEntity
	if err := r.db.Where("id = ?", productID).First(&e).Error; err != nil {
		return nil, err
	}

	p := &domain.Product{
		ID:        e.ID,
		Name:      e.Name,
		Available: e.Available,
		ImageURL:  e.ImageURL,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}

	return p, nil
}
