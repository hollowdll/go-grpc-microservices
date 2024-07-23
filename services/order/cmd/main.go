package main

import (
	"log"

	"github.com/hollowdll/go-grpc-microservices/services/order/config"
	"github.com/hollowdll/go-grpc-microservices/services/order/internal/adapters/inventory"
	"github.com/hollowdll/go-grpc-microservices/services/order/internal/adapters/payment"
	"github.com/hollowdll/go-grpc-microservices/services/order/internal/application/core/api"
)

func main() {
	log.Println("starting order service ...")

	config.InitConfig()
	cfg := config.NewConfig()
	log.Printf("running application in %s mode", cfg.ApplicationMode)

	inventoryAdapter, err := inventory.NewAdapter(cfg)
	if err != nil {
		log.Fatalf("failed to initialize gRPC client for inventory service: %v", err)
	}

	paymentAdapter, err := payment.NewAdapter(cfg)
	if err != nil {
		log.Fatalf("failed to initialize gRPC client for payment service: %v", err)
	}

	application := api.NewApplication(inventoryAdapter, paymentAdapter)
}
