package inventory

import (
	"context"
	"fmt"
	"log"

	"github.com/hollowdll/go-grpc-microservices/services/order/config"
	"github.com/hollowdll/go-grpc-microservices/services/order/internal/application/core/domain"
	"github.com/hollowdll/grpc-microservices-proto/gen/golang/inventorypb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	inventory inventorypb.InventoryServiceClient
	conn      *grpc.ClientConn
}

func NewAdapter(cfg *config.Config) (*Adapter, error) {
	log.Printf("creating gRPC client for inventory service ...")

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	address := fmt.Sprintf("%s:%d", cfg.InventoryServiceHost, cfg.InventoryServicePort)

	log.Printf("using endpoint %s", address)
	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		return nil, err
	}
	client := inventorypb.NewInventoryServiceClient(conn)

	return &Adapter{
		inventory: client,
		conn:      conn,
	}, nil
}

func (a *Adapter) CloseConnection() {
	if a.conn != nil {
		if err := a.conn.Close(); err != nil {
			log.Printf("failed to close gRPC client connection to inventory service: %v", err)
		}
	}
}

func (a *Adapter) GetProductPrices(ctx context.Context, productCodes []string) ([]*domain.ProductPrice, error) {
	response, err := a.inventory.GetProductDetails(ctx, &inventorypb.GetProductDetailsRequest{
		ProductCodes: productCodes,
	})
	if err != nil {
		return nil, err
	}

	var productPrices = []*domain.ProductPrice{}
	for _, product := range response.ProductDetails {
		productPrices = append(productPrices, &domain.ProductPrice{
			ProductCode:    product.ProductCode,
			UnitPriceCents: product.UnitPriceCents,
		})
	}

	return productPrices, nil
}

func (a *Adapter) CheckProductStockQuantities(ctx context.Context, orderItems []*domain.OrderItem) ([]*domain.ProductStock, error) {
	var productQuantities = []*inventorypb.ProductQuantity{}
	for _, orderItem := range orderItems {
		productQuantities = append(productQuantities, &inventorypb.ProductQuantity{
			ProductCode: orderItem.ProductCode,
			Quantity:    orderItem.Quantity,
		})
	}

	response, err := a.inventory.CheckProductStockQuantity(ctx, &inventorypb.CheckProductStockQuantityRequest{
		Products: productQuantities,
	})
	if err != nil {
		return nil, err
	}

	var productStocks = []*domain.ProductStock{}
	for _, product := range response.Products {
		productStocks = append(productStocks, &domain.ProductStock{
			ProductCode:       product.ProductCode,
			AvailableQuantity: product.AvailableQuantity,
			IsAvailable:       product.IsAvailable,
		})
	}

	return productStocks, nil
}

func (a *Adapter) ReduceProductStockQuantities(ctx context.Context, orderItems []*domain.OrderItem) error {
	var productQuantities = []*inventorypb.ProductQuantity{}
	for _, orderItem := range orderItems {
		productQuantities = append(productQuantities, &inventorypb.ProductQuantity{
			ProductCode: orderItem.ProductCode,
			Quantity:    orderItem.Quantity,
		})
	}

	_, err := a.inventory.ReduceProductStockQuantity(ctx, &inventorypb.ReduceProductStockQuantityRequest{
		Products: productQuantities,
	})

	return err
}
