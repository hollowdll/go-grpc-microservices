package api

import (
	"context"
	"errors"

	"github.com/hollowdll/go-grpc-microservices/services/order/internal/application/core/domain"
	"github.com/hollowdll/go-grpc-microservices/services/order/internal/ports"
)

// Application is the core of the application.
// It contains business logic and outbound port interfaces.
// Business logic can call outbound adapters using these interfaces.
type Application struct {
	inventory ports.InventoryPort
	payment   ports.PaymentPort
}

func NewApplication(inventory ports.InventoryPort, payment ports.PaymentPort) *Application {
	return &Application{
		inventory: inventory,
		payment:   payment,
	}
}

func (a *Application) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	// 1. get product prices
	var productCodes = []string{}
	for _, orderItem := range order.OrderItems {
		productCodes = append(productCodes, orderItem.ProductCode)
	}

	productPrices, err := a.inventory.GetProductPrices(ctx, productCodes)
	if err != nil {
		return nil, err
	}

	// 2. check product stock quantities
	productStocks, err := a.inventory.CheckProductStockQuantities(ctx, order.OrderItems)
	if err != nil {
		return nil, err
	}
	for _, productStock := range productStocks {
		if !productStock.IsAvailable {
			return nil, errors.New("not enough ordered products in stock")
		}
	}

	// 3. make payment
	totalPriceCents := sumPrices(productPrices, order.OrderItems)
	err = a.payment.CreatePayment(ctx, order, totalPriceCents)
	if err != nil {
		return nil, err
	}

	// 4. reduce product stock quantity
	err = a.inventory.ReduceProductStockQuantities(ctx, order.OrderItems)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func sumPrices(productPrices []*domain.ProductPrice, orderItems []*domain.OrderItem) int32 {
	quantities := make(map[string]int32)
	for _, orderItem := range orderItems {
		quantities[orderItem.ProductCode] = orderItem.Quantity
	}

	var sum int32 = 0
	for _, productPrice := range productPrices {
		sum += (productPrice.UnitPriceCents * quantities[productPrice.ProductCode])
	}
	return sum
}
