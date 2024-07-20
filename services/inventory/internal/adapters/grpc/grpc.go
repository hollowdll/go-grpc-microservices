package grpc

import (
	"context"
	"fmt"
	"log"

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

	products, err := a.api.GetProductDetails(ctx, req.ProductCodes)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to get product details: %v", err))
	}

	var productDetails = []*inventorypb.ProductDetails{}
	for _, product := range products {
		productDetails = append(productDetails, &inventorypb.ProductDetails{
			ProductCode:    product.ProductCode,
			Name:           product.Name,
			Description:    product.Description,
			UnitPriceCents: product.UnitPriceCents,
		})
	}

	return &inventorypb.GetProductDetailsResponse{ProductDetails: productDetails}, nil
}
