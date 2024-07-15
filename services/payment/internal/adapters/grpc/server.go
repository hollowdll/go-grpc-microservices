package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/hollowdll/go-grpc-microservices/services/payment/internal/ports"
	"github.com/hollowdll/grpc-microservices-proto/gen/golang/paymentpb"
	"google.golang.org/grpc"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *grpc.Server
	paymentpb.UnimplementedPaymentServiceServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

// Run runs the gRPC server and starts to listen for requests.
func (a *Adapter) Run() {
	log.Printf("initializing payment service gRPC server on port %d ...", a.port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	paymentpb.RegisterPaymentServiceServer(grpcServer, a)

	log.Printf("starting payment service gRPC server ...")
	log.Printf("payment service gRPC server listening at %v", lis.Addr())

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC at %v", lis.Addr())
	}
}
