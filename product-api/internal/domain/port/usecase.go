package port

import (
	"context"

	"github.com/fiap-postech-soat1-group21-stage4/product-api/product-api/internal/domain/entity"
)

// ProductUseCase is the interface for product repository
type ProductUseCase interface {
	Create(ctx context.Context, product *entity.Product) (*entity.Product, error)
	Update(ctx context.Context, product *entity.Product) error
	Delete(ctx context.Context, pdt *entity.Product) error
	GetProducts(context.Context) (*entity.ProductResponseList, error)
}
