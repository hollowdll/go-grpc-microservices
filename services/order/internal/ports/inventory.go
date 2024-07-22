package ports

import (
	"context"

	"github.com/hollowdll/go-grpc-microservices/services/order/internal/application/core/domain"
)

type InventoryPort interface {
	GetProductPrices(ctx context.Context, productCodes []string) ([]*domain.ProductPrice, error)
	CheckProductStockQuantities(ctx context.Context, orderItems []*domain.OrderItem) ([]*domain.ProductStock, error)
	ReduceProductStockQuantities(ctx context.Context, orderItems []*domain.OrderItem) error
}
