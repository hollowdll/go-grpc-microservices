package api

import (
	"context"
	"errors"
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
	products, err := a.db.GetProductsByCode(ctx, productCodes)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (a *Application) CheckProductStockQuantity(ctx context.Context, productQuantities []*domain.ProductQuantity) ([]*domain.ProductStock, error) {
	var productCodes = []string{}
	for _, productQuantity := range productQuantities {
		productCodes = append(productCodes, productQuantity.ProductCode)
	}

	products, err := a.db.GetProductsByCode(ctx, productCodes)
	if err != nil {
		return nil, err
	}

	var productStocks = []*domain.ProductStock{}
	for _, productQuantity := range productQuantities {
		for _, product := range products {
			if productQuantity.ProductCode == product.ProductCode {
				if productQuantity.Quantity <= product.QuantityInStock {
					productStocks = append(productStocks, &domain.ProductStock{
						ProductCode:       product.ProductCode,
						AvailableQuantity: product.QuantityInStock,
						IsAvailable:       true,
					})
				} else {
					productStocks = append(productStocks, &domain.ProductStock{
						ProductCode:       product.ProductCode,
						AvailableQuantity: product.QuantityInStock,
						IsAvailable:       false,
					})
				}
				break
			}
		}
	}

	return productStocks, nil
}

func (a *Application) ReduceProductStockQuantity(ctx context.Context, productQuantities []*domain.ProductQuantity) ([]*domain.ProductStock, error) {
	var productCodes = []string{}
	for _, productQuantity := range productQuantities {
		productCodes = append(productCodes, productQuantity.ProductCode)
	}

	products, err := a.db.GetProductsByCode(ctx, productCodes)
	if err != nil {
		return nil, err
	}

	var productStocks = []*domain.ProductStock{}
	var updatedQuantities = []*domain.ProductQuantity{}
	for _, productQuantity := range productQuantities {
		for _, product := range products {
			if productQuantity.ProductCode == product.ProductCode {
				updatedQuantity := product.QuantityInStock - productQuantity.Quantity
				if updatedQuantity < 0 {
					return nil, errors.New("negative product stock quantity")
				}

				updatedQuantities = append(updatedQuantities, &domain.ProductQuantity{
					ProductCode: product.ProductCode,
					Quantity:    updatedQuantity,
				})
				productStock := &domain.ProductStock{
					ProductCode:       product.ProductCode,
					AvailableQuantity: updatedQuantity,
					IsAvailable:       true,
				}

				if updatedQuantity == 0 {
					productStock.IsAvailable = false
				}
				productStocks = append(productStocks, productStock)
			}
			break
		}
	}

	err = a.db.UpdateProductStockQuantities(ctx, updatedQuantities)
	if err != nil {
		return nil, err
	}

	return productStocks, nil
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

	log.Println("saving test data to database ...")
	for _, product := range products {
		log.Printf("product = %v", product)
	}

	err := a.db.SaveProducts(ctx, products)
	if err != nil {
		return err
	}

	return nil
}
