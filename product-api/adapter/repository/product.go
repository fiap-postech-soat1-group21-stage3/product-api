package repository

import (
	"context"

	"github.com/fiap-postech-soat1-group21-stage4/product-api/product-api/internal/domain/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Product struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *Product {
	return &Product{db}
}

func (p *Product) Create(ctx context.Context, pr *entity.Product) (*entity.Product, error) {
	dbFn := p.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true})

	if result := dbFn.Table("product").Create(pr); result.Error != nil {
		return nil, result.Error
	}

	return pr, nil
}

func (p *Product) Update(ctx context.Context, pr *entity.Product) error {
	dbFn := p.db.WithContext(ctx)

	filter := pr.ID

	return dbFn.Table("product").Where(filter).Updates(pr).Error
}

func (p *Product) Delete(ctx context.Context, pr *entity.Product) error {
	dbFn := p.db.WithContext(ctx)

	id := pr.ID

	return dbFn.Table("product").Where("id = ?", id).Delete(&pr).Error
}

func (p *Product) GetProducts(ctx context.Context) (*entity.ProductResponseList, error) {
	dbFn := p.db.WithContext(ctx)

	var count int64
	var products []*entity.Product

	result := dbFn.Table("product").Find(&products).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}
	return &entity.ProductResponseList{
		Result: products,
		Count:  count,
	}, nil
}
