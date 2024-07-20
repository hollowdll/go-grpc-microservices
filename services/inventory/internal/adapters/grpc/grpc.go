package grpc

import (
	"context"
	"fmt"
	"log"

	"github.com/hollowdll/go-grpc-microservices/services/inventory/internal/application/core/domain"
	"github.com/hollowdll/grpc-microservices-proto/gen/golang/inventorypb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	getProductDetailsPrefix          = "GetProductDetails"
	checkProductStockQuantityPrefix  = "CheckProductStockQuantity"
	reduceProductStockQuantityPrefix = "ReduceProductStockQuantity"
)

func (a *Adapter) GetProductDetails(ctx context.Context, req *inventorypb.GetProductDetailsRequest) (res *inventorypb.GetProductDetailsResponse, err error) {
	log.Printf("call RPC %s: request = %v", getProductDetailsPrefix, req)
	defer func() {
		if err != nil {
			log.Printf("RPC %s failed: request = %v; error = %v", getProductDetailsPrefix, req, err)
		} else {
			log.Printf("RPC %s success: request = %v; response = %v", getProductDetailsPrefix, req, res)
		}
	}()

	results, err := a.api.GetProductDetails(ctx, req.ProductCodes)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to get product details: %v", err))
	}

	var productDetails = []*inventorypb.ProductDetails{}
	for _, product := range results {
		productDetails = append(productDetails, &inventorypb.ProductDetails{
			ProductCode:    product.ProductCode,
			Name:           product.Name,
			Description:    product.Description,
			UnitPriceCents: product.UnitPriceCents,
		})
	}

	return &inventorypb.GetProductDetailsResponse{ProductDetails: productDetails}, nil
}

func (a *Adapter) CheckProductStockQuantity(ctx context.Context, req *inventorypb.CheckProductStockQuantityRequest) (res *inventorypb.CheckProductStockQuantityResponse, err error) {
	log.Printf("call RPC %s: request = %v", checkProductStockQuantityPrefix, req)
	defer func() {
		if err != nil {
			log.Printf("RPC %s failed: request = %v; error = %v", checkProductStockQuantityPrefix, req, err)
		} else {
			log.Printf("RPC %s success: request = %v; response = %v", checkProductStockQuantityPrefix, req, res)
		}
	}()

	var productQuantities = []*domain.ProductQuantity{}
	for _, product := range req.Products {
		productQuantities = append(productQuantities, &domain.ProductQuantity{
			ProductCode: product.ProductCode,
			Quantity:    product.Quantity,
		})
	}
	results, err := a.api.CheckProductStockQuantity(ctx, productQuantities)

	var productStocks = []*inventorypb.ProductStock{}
	for _, result := range results {
		productStocks = append(productStocks, &inventorypb.ProductStock{
			ProductCode:       result.ProductCode,
			AvailableQuantity: result.AvailableQuantity,
			IsAvailable:       result.IsAvailable,
		})
	}

	return &inventorypb.CheckProductStockQuantityResponse{Products: productStocks}, nil
}

func (a *Adapter) ReduceProductStockQuantity(ctx context.Context, req *inventorypb.ReduceProductStockQuantityRequest) (res *inventorypb.ReduceProductStockQuantityResponse, err error) {
	log.Printf("call RPC %s: request = %v", reduceProductStockQuantityPrefix, req)
	defer func() {
		if err != nil {
			log.Printf("RPC %s failed: request = %v; error = %v", reduceProductStockQuantityPrefix, req, err)
		} else {
			log.Printf("RPC %s success: request = %v; response = %v", reduceProductStockQuantityPrefix, req, res)
		}
	}()

	var productQuantities = []*domain.ProductQuantity{}
	for _, product := range req.Products {
		productQuantities = append(productQuantities, &domain.ProductQuantity{
			ProductCode: product.ProductCode,
			Quantity:    product.Quantity,
		})
	}
	results, err := a.api.ReduceProductStockQuantity(ctx, productQuantities)

	var productStocks = []*inventorypb.ProductStock{}
	for _, result := range results {
		productStocks = append(productStocks, &inventorypb.ProductStock{
			ProductCode:       result.ProductCode,
			AvailableQuantity: result.AvailableQuantity,
			IsAvailable:       result.IsAvailable,
		})
	}

	return &inventorypb.ReduceProductStockQuantityResponse{Products: productStocks}, nil
}
