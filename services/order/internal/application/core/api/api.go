package api

import (
	"context"

	"github.com/hollowdll/go-grpc-microservices/services/order/internal/application/core/domain"
	"github.com/hollowdll/go-grpc-microservices/services/order/internal/ports"
)

// Application is the core of the application.
// It contains business logic and outbound port interfaces.
// Business logic can call outbound adapters using these interfaces.
type Application struct {
	inventory ports.InventoryPort
	payment   ports.PaymentPort
}

func NewApplication(inventory ports.InventoryPort, payment ports.PaymentPort) *Application {
	return &Application{
		inventory: inventory,
		payment:   payment,
	}
}

func CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	// Not yet implemented.
	return nil, nil
}
