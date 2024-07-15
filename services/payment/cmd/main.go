package main

import (
	"log"

	"github.com/hollowdll/go-grpc-microservices/services/payment/config"
	"github.com/hollowdll/go-grpc-microservices/services/payment/internal/adapters/grpc"
	"github.com/hollowdll/go-grpc-microservices/services/payment/internal/application/core/api"
)

func main() {
	cfg := config.NewConfig()
	log.Printf("starting payment service in %s mode ...", cfg.ApplicationMode)

	application := api.NewApplication()
	grpcAdapter := grpc.NewAdapter(application, cfg)
	grpcAdapter.Run()
}
