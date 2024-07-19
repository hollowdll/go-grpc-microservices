package ports

import (
	"context"

	"github.com/hollowdll/go-grpc-microservices/services/inventory/internal/application/core/domain"
)

type DBPort interface {
	GetProductsByCode(ctx context.Context, productCodes []string) ([]*domain.Product, error)
	UpdateProductStockQuantities(ctx context.Context, products []*domain.ProductQuantity) error
	SaveProduct(ctx context.Context, product *domain.Product) error
}
