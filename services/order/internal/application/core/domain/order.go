package domain

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	OrderPending   OrderStatus = "Pending"
	OrderCompleted OrderStatus = "Completed"
	OrderFailed    OrderStatus = "Failed"
	OrderCancelled OrderStatus = "Cancelled"
)

type OrderItem struct {
	ProductCode string `json:"product_code"`
	Quantity    int32  `json:"quantity"`
}

type Order struct {
	ID              string      `json:"id"`
	CustomerID      string      `json:"customer_id"`
	Status          OrderStatus `json:"status"`
	OrderItems      []OrderItem `json:"order_items"`
	CreatedAtMillis int64       `json:"created_at_millis"`
	UpdatedAtMillis int64       `json:"updated_at_millis"`
}

func NewOrder(customerID string, orderItems []OrderItem) (*Order, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &Order{
		ID:              id.String(),
		CustomerID:      customerID,
		Status:          OrderPending,
		OrderItems:      orderItems,
		CreatedAtMillis: time.Now().UnixMilli(),
		UpdatedAtMillis: time.Now().UnixMilli(),
	}, nil
}

type ProductPrice struct {
	ProductCode    string `json:"product_code"`
	UnitPriceCents int32  `json:"unit_price_cents"`
}

type ProductStock struct {
	ProductCode       string `json:"product_code"`
	AvailableQuantity int32  `json:"available_quantity"`
	IsAvailable       bool   `json:"is_available"`
}
