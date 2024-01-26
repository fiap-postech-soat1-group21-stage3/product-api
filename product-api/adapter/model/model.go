package model

import (
	"time"

	"github.com/google/uuid"
)

type ProductCategory string

const (
	Burgers  ProductCategory = "burgers"
	Sides    ProductCategory = "sides"
	Beverage ProductCategory = "beverage"
	Sweets   ProductCategory = "sweets"
)

type ProductRequestDTO struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Category    ProductCategory `json:"category"`
	Price       float64         `json:"price"`
}

type ProductResponseDTO struct {
	ID          uuid.UUID       `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Category    ProductCategory `json:"category"`
	Price       float64         `json:"price"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type ProductResponseList struct {
	Result []*ProductResponseDTO `json:"result"`
	Count  int64                 `json:"count"`
}
