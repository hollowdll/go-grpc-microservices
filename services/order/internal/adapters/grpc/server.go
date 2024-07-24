package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/hollowdll/go-grpc-microservices/services/order/config"
	"github.com/hollowdll/go-grpc-microservices/services/order/internal/ports"
	"github.com/hollowdll/grpc-microservices-proto/gen/golang/orderpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api    ports.APIPort
	cfg    *config.Config
	server *grpc.Server
	orderpb.UnimplementedOrderServiceServer
}

func NewAdapter(api ports.APIPort, cfg *config.Config) *Adapter {
	return &Adapter{api: api, cfg: cfg}
}

func (a *Adapter) Run() {
	log.Printf("initializing order service gRPC server on port %d ...", a.cfg.GrpcPort)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.cfg.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", a.cfg.GrpcPort, err)
	}

	grpcServer := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(grpcServer, a)

	// this enables gRPC services to be tested with e.g. grpcurl
	if a.cfg.IsDevelopmentMode() {
		log.Println("development mode detected: enabling gRPC server reflection ...")
		reflection.Register(grpcServer)
	}

	log.Printf("starting order service gRPC server ...")
	log.Printf("order service gRPC server listening at %v", lis.Addr())

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC at %v", lis.Addr())
	}
}
