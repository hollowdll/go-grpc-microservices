package payment

import (
	"context"
	"fmt"
	"log"

	"github.com/hollowdll/go-grpc-microservices/services/order/config"
	"github.com/hollowdll/go-grpc-microservices/services/order/internal/application/core/domain"
	"github.com/hollowdll/grpc-microservices-proto/gen/golang/paymentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment paymentpb.PaymentServiceClient
	conn    *grpc.ClientConn
}

func NewAdapter(cfg *config.Config) (*Adapter, error) {
	log.Printf("creating gRPC client for payment service ...")

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	address := fmt.Sprintf("%s:%d", cfg.PaymentServiceHost, cfg.PaymentServicePort)

	log.Printf("using endpoint %s", address)
	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		return nil, err
	}
	client := paymentpb.NewPaymentServiceClient(conn)

	return &Adapter{
		payment: client,
		conn:    conn,
	}, nil
}

func (a *Adapter) CloseConnection() {
	if a.conn != nil {
		if err := a.conn.Close(); err != nil {
			log.Printf("failed to close gRPC client connection to payment service: %v", err)
		}
	}
}

func (a *Adapter) CreatePayment(ctx context.Context, order *domain.Order, totalPriceCents int32) error {
	return nil
}
