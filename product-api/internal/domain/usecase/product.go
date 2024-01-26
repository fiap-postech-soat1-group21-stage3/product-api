package usecase

import (
	"context"

	"github.com/fiap-postech-soat1-group21/product-api/product-api/internal/domain/entity"
	"github.com/fiap-postech-soat1-group21/product-api/product-api/internal/domain/port"
)

type useCaseProduct struct {
	repository port.ProductRepository
}

// NewProductUseCase is responsible for all use cases for products
func NewProductUseCase(p port.ProductRepository) port.ProductUseCase {
	return &useCaseProduct{
		repository: p,
	}
}

// Create create and persist product data
func (u *useCaseProduct) Create(ctx context.Context, input *entity.Product) (*entity.Product, error) {
	p, err := u.repository.Create(ctx, input)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Update updates a product and persist new data
func (u *useCaseProduct) Update(ctx context.Context, input *entity.Product) error {
	if err := u.repository.Update(ctx, input); err != nil {
		return err
	}

	return nil
}

// Delete remove a product data
func (u *useCaseProduct) Delete(ctx context.Context, input *entity.Product) error {
	if err := u.repository.Delete(ctx, input); err != nil {
		return err
	}

	return nil
}

// List retrieves all products
func (u *useCaseProduct) GetProducts(ctx context.Context) (*entity.ProductResponseList, error) {
	res, err := u.repository.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}
