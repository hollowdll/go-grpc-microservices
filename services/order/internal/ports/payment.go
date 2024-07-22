package ports

import (
	"context"

	"github.com/hollowdll/go-grpc-microservices/services/order/internal/application/core/domain"
)

type PaymentPort interface {
	CreatePayment(ctx context.Context, order *domain.Order, totalPriceCents int32) error
}
