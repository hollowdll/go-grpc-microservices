package main

import (
	"context"
	"log"
	"time"

	"github.com/hollowdll/go-grpc-microservices/services/inventory/config"
	"github.com/hollowdll/go-grpc-microservices/services/inventory/internal/adapters/db"
	"github.com/hollowdll/go-grpc-microservices/services/inventory/internal/adapters/grpc"
	"github.com/hollowdll/go-grpc-microservices/services/inventory/internal/application/core/api"
)

func initApplication(application *api.Application, cfg *config.Config) {
	if cfg.IsDevelopmentMode() {
		log.Println("development mode detected: populating test data ...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := application.PopulateTestData(ctx); err != nil {
			log.Fatalf("failed to populate test data: %v", err)
		}
	}
}

func main() {
	log.Println("starting inventory service ...")

	config.InitConfig()
	cfg := config.LoadConfig()
	log.Printf("running application in %s mode", cfg.ApplicationMode)

	dbAdapter := db.NewInMemoryDBAdapter()
	application := api.NewApplication(dbAdapter)
	initApplication(application, cfg)

	grpcAdapter := grpc.NewAdapter(application, cfg)
	grpcAdapter.Run()
}
