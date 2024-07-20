package db

import (
	"context"
	"sync"

	"github.com/hollowdll/go-grpc-microservices/services/inventory/internal/application/core/domain"
)

// InMemoryDB is a simple in-memory database to demonstrate and simulate this example microservice.
// With hexagonal architecture, you can make multiple different database implementations
// for different SQL and NoSQL databases that use the same interface.
type InMemoryDB struct {
	products map[string]*domain.Product
	mu       sync.RWMutex
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		products: make(map[string]*domain.Product),
	}
}

func (db *InMemoryDB) AddProducts(products []*domain.Product) {
	db.mu.Lock()
	defer db.mu.Unlock()

	for _, product := range products {
		db.products[product.ProductCode] = product
	}
}

func (db *InMemoryDB) GetProductsByCode(productCodes []string) []*domain.Product {
	db.mu.RLock()
	defer db.mu.RUnlock()

	products := []*domain.Product{}
	for _, code := range productCodes {
		if product, ok := db.products[code]; ok {
			products = append(products, product)
		}
	}

	return products
}

func (db *InMemoryDB) UpdateProductStockQuantities(updatedQuantities map[string]int32) {
	db.mu.Lock()
	defer db.mu.Unlock()

	for productCode, newQuantity := range updatedQuantities {
		if product, ok := db.products[productCode]; ok {
			product.QuantityInStock = newQuantity
			db.products[productCode] = product
		}
	}
}

type InMemoryDBAdapter struct {
	db *InMemoryDB
}

func NewInMemoryDBAdapter() *InMemoryDBAdapter {
	return &InMemoryDBAdapter{
		db: NewInMemoryDB(),
	}
}

func (a *InMemoryDBAdapter) GetProductsByCode(ctx context.Context, productCodes []string) ([]*domain.Product, error) {
	products := a.db.GetProductsByCode(productCodes)

	return products, nil
}

func (a *InMemoryDBAdapter) UpdateProductStockQuantities(ctx context.Context, products []*domain.ProductQuantity) error {
	var updatedQuantities = make(map[string]int32)
	for _, product := range products {
		updatedQuantities[product.ProductCode] = product.Quantity
	}
	a.db.UpdateProductStockQuantities(updatedQuantities)

	return nil
}

func (a *InMemoryDBAdapter) SaveProducts(ctx context.Context, products []*domain.Product) error {
	a.db.AddProducts(products)

	return nil
}
