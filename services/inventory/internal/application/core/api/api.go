package api

import (
	"context"

	"github.com/hollowdll/go-grpc-microservices/services/inventory/internal/application/core/domain"
)

// Application contains core business logic.
type Application struct {
	// Here we can add outbound ports for e.g. database
	// that business logic can call.
}

func NewApplication() *Application {
	return &Application{}
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
