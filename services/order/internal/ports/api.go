package ports

import (
	"context"

	"github.com/hollowdll/go-grpc-microservices/services/order/internal/application/core/domain"
)

type APIPort interface {
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
}
