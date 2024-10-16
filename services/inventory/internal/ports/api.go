package ports

import (
	"context"

	"github.com/hollowdll/go-grpc-microservices/services/inventory/internal/application/core/domain"
)

type APIPort interface {
	GetProductDetails(ctx context.Context, productCodes []string) ([]*domain.Product, error)
	CheckProductStockQuantity(ctx context.Context, productQuantities []*domain.ProductQuantity) ([]*domain.ProductStock, error)
	ReduceProductStockQuantity(ctx context.Context, productQuantities []*domain.ProductQuantity) ([]*domain.ProductStock, error)
	PopulateTestData(ctx context.Context) error
}
