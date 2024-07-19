package api

import (
	"context"
	"log"

	"github.com/hollowdll/go-grpc-microservices/services/inventory/internal/application/core/domain"
	"github.com/hollowdll/go-grpc-microservices/services/inventory/internal/ports"
)

// Application is the core of the application. It contains all business logic.
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

// PopulateTestData is used to save some test data to database.
func (a *Application) PopulateTestData(ctx context.Context) error {
	products := []*domain.Product{
		domain.NewProduct("Bike", "Mountain bike.", 80000, 5),
		domain.NewProduct("Laptop", "Business Laptop.", 150000, 10),
		domain.NewProduct("Cable", "HDMI cable.", 1200, 12),
		domain.NewProduct("Keyboard", "Mechanical keyboard.", 8000, 8),
		domain.NewProduct("Monitor", "OLED monitor.", 11000, 3),
		domain.NewProduct("Chair", "Office chair.", 15000, 0),
	}

	log.Printf("saving the following test data to database: products = %v", products)

	err := a.db.SaveProducts(ctx, products)
	if err != nil {
		return err
	}

	return nil
}
