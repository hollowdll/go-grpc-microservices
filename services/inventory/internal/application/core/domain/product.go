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

func NewProduct(name string, description string, unitPriceCenters int32) *Product {
	return &Product{
		ProductCode:     uuid.NewString(),
		Name:            name,
		Description:     description,
		UnitPriceCents:  unitPriceCenters,
		QuantityInStock: 0,
		CreatedAtMillis: time.Now().UnixMilli(),
		UpdatedAtMillis: time.Now().UnixMilli(),
	}
}
