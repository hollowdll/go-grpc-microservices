package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/hollowdll/go-grpc-microservices/services/inventory/config"
	"github.com/hollowdll/go-grpc-microservices/services/inventory/internal/ports"
	"github.com/hollowdll/grpc-microservices-proto/gen/golang/inventorypb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api    ports.APIPort
	config *config.Config
	server *grpc.Server
	inventorypb.UnimplementedInventoryServiceServer
}

func NewAdapter(api ports.APIPort, config *config.Config) *Adapter {
	return &Adapter{api: api, config: config}
}

// Run runs the gRPC server and starts to listen for requests.
func (a *Adapter) Run() {
	log.Printf("initializing inventory service gRPC server on port %d ...", a.config.GrpcPort)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.config.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", a.config.GrpcPort, err)
	}

	grpcServer := grpc.NewServer()
	inventorypb.RegisterInventoryServiceServer(grpcServer, a)

	// this enables gRPC services to be tested with e.g. grpcurl
	if a.config.IsDevelopmentMode() {
		log.Println("development mode detected: enabling gRPC server reflection ...")
		reflection.Register(grpcServer)
	}

	log.Printf("starting inventory service gRPC server ...")
	log.Printf("inventory service gRPC server listening at %v", lis.Addr())

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC at %v", lis.Addr())
	}
}
