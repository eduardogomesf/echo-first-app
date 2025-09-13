package entities

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Price      float64   `json:"price"`
	Categories []string  `json:"categories"`
	IsDisabled bool      `json:"isDisabled"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func NewProduct(
	name string,
	price float64,
	categories []string,
	isDisabled bool,
) *Product {
	now := time.Now()

	return &Product{
		Id:         uuid.New(),
		Name:       name,
		Price:      price,
		Categories: categories,
		IsDisabled: isDisabled,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}
