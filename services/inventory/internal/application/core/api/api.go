package api

import (
	"context"

	"github.com/hollowdll/go-grpc-microservices/services/inventory/internal/application/core/domain"
	"github.com/hollowdll/go-grpc-microservices/services/inventory/internal/ports"
)

// Application contains core business logic.
type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a *Application) GetProductDetails(ctx context.Context, productCodes []string) ([]*domain.Product, error) {
	// Some business logic.

	return []*domain.Product{}, nil
}

func (a *Application) CheckProductStockQuantity(ctx context.Context, products []*domain.ProductQuantity) ([]*domain.ProductStock, error) {
	// Some business logic.

	return []*domain.ProductStock{}, nil
}

func (a *Application) ReduceProductStockQuantity(ctx context.Context, products []*domain.ProductQuantity) ([]*domain.ProductStock, error) {
	// Some business logic.

	return []*domain.ProductStock{}, nil
}
