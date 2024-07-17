package grpc

import (
	"context"
	"fmt"
	"log"

	"github.com/hollowdll/go-grpc-microservices/services/payment/internal/application/core/domain"
	"github.com/hollowdll/grpc-microservices-proto/gen/golang/paymentpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const createPaymentPrefix string = "CreatePayment"

func (a *Adapter) CreatePayment(ctx context.Context, req *paymentpb.CreatePaymentRequest) (res *paymentpb.CreatePaymentResponse, err error) {
	log.Printf("call RPC %s: request = %v", createPaymentPrefix, req)
	defer func() {
		if err != nil {
			log.Printf("RPC %s failed: request = %v; error = %v", createPaymentPrefix, req, err)
		} else {
			log.Printf("RPC %s success: request = %v; response = %v", createPaymentPrefix, req, res)
		}
	}()

	newPayment := domain.NewPayment(req.CustomerId, req.OrderId, req.TotalPriceCents)
	result, err := a.api.Charge(ctx, newPayment)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to charge: %v", err))
	}

	return &paymentpb.CreatePaymentResponse{PaymentId: result.ID}, nil
}
