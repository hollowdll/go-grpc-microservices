package main

import (
	"log"

	"github.com/hollowdll/go-grpc-microservices/services/inventory/config"
	"github.com/hollowdll/go-grpc-microservices/services/inventory/internal/adapters/db"
	"github.com/hollowdll/go-grpc-microservices/services/inventory/internal/adapters/grpc"
	"github.com/hollowdll/go-grpc-microservices/services/inventory/internal/application/core/api"
)

func main() {
	log.Println("starting inventory service ...")

	cfg := config.NewConfig()
	log.Printf("running application in %s mode", cfg.ApplicationMode)

	dbAdapter := db.NewInMemoryDBAdapter()
	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, cfg)
	grpcAdapter.Run()
}
