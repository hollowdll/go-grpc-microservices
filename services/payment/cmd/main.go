package main

import (
	"github.com/hollowdll/go-grpc-microservices/services/payment/config"
	"github.com/hollowdll/go-grpc-microservices/services/payment/internal/adapters/grpc"
	"github.com/hollowdll/go-grpc-microservices/services/payment/internal/application/core/api"
)

func main() {
	cfg := config.NewConfig()
	application := api.NewApplication()
	grpcAdapter := grpc.NewAdapter(application, cfg)
	grpcAdapter.Run()
}
