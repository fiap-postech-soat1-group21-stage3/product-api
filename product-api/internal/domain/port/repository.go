package port

import (
	"context"

	"github.com/fiap-postech-soat1-group21/product-api/product-api/internal/domain/entity"
)

// ProductRepository is the interface for product database
type ProductRepository interface {
	Create(ctx context.Context, product *entity.Product) (*entity.Product, error)
	Update(ctx context.Context, product *entity.Product) error
	Delete(ctx context.Context, pdt *entity.Product) error
	GetProducts(context.Context) (*entity.ProductResponseList, error)
}
