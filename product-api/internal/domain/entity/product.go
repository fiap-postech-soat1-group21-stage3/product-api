package entity

import (
	"time"

	"github.com/google/uuid"
)

type Category string

const (
	Burguers Category = "burguers"
	Sides    Category = "sides"
	Beverage Category = "beverage"
	Sweets   Category = "sweets"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Category    Category  `gorm:"type:enum('burguers', 'sides', 'beverage', 'sweets')"`
	Price       float64   `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"not null;autoUpdateTime"`
}

// ProductResponseList summary list
type ProductResponseList struct {
	Result []*Product
	Count  int64
}
