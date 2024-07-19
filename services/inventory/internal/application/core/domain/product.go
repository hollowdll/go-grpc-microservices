package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ProductCode     string `json:"product_code"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	UnitPriceCents  int32  `json:"unit_price_cents"`
	QuantityInStock int32  `json:"quantity_in_stock"`
	CreatedAtMillis int64  `json:"created_at_millis"`
	UpdatedAtMillis int64  `json:"updated_at_millis"`
}

func NewProduct(name string, description string, unitPriceCenters int32, quantityInStock int32) *Product {
	return &Product{
		ProductCode:     uuid.NewString(),
		Name:            name,
		Description:     description,
		UnitPriceCents:  unitPriceCenters,
		QuantityInStock: quantityInStock,
		CreatedAtMillis: time.Now().UnixMilli(),
		UpdatedAtMillis: time.Now().UnixMilli(),
	}
}

type ProductQuantity struct {
	ProductCode string `json:"product_code"`
	Quantity    int32  `json:"quantity"`
}

type ProductStock struct {
	ProductCode       string `json:"product_code"`
	AvailableQuantity int32  `json:"available_quantity"`
	IsAvailable       bool   `json:"is_available"`
}
