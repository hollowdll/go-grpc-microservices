package domain

import (
	"time"

	"github.com/google/uuid"
)

type PaymentStatus string

const (
	PaymentPending   PaymentStatus = "Pending"
	PaymentCompleted PaymentStatus = "Completed"
	PaymentFailed    PaymentStatus = "Failed"
	PaymentCancelled PaymentStatus = "Cancelled"
)

func (s PaymentStatus) String() string {
	return string(s)
}

type Payment struct {
	ID              string        `json:"id"`
	CustomerID      string        `json:"customer_id"`
	Status          PaymentStatus `json:"status"`
	OrderID         string        `json:"order_id"`
	TotalPriceCents int32         `json:"total_price_cents"`
	CreatedAtMillis int64         `json:"created_at_millis"`
	UpdatedAtMillis int64         `json:"updated_at_millis"`
}

func NewPayment(customerID string, orderID string, totalPriceCents int32) *Payment {
	return &Payment{
		ID:              uuid.NewString(),
		CustomerID:      customerID,
		Status:          PaymentPending,
		OrderID:         orderID,
		TotalPriceCents: totalPriceCents,
		CreatedAtMillis: time.Now().UnixMilli(),
		UpdatedAtMillis: time.Now().UnixMilli(),
	}
}
