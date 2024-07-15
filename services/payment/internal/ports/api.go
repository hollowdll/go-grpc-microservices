package ports

import (
	"context"

	"github.com/hollowdll/go-grpc-microservices/services/payment/internal/application/core/domain"
)

// APIPort is an inbound port interface for the core business logic
// that the application can implement.
type APIPort interface {
	Charge(ctx context.Context, payment *domain.Payment) (*domain.Payment, error)
}
