package main

import (
	"net"
	"os"

	"github.com/hown3d/pizza-shop-grpc-example/bakery/bakery"
	bakerypb "github.com/hown3d/pizza-shop-grpc-example/bakery/proto/bakery/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
	//serverPort := os.Getenv("SERVER_PORT")
	serverPort := "9090"
	lis, err := net.Listen("tcp", "0.0.0.0:"+serverPort)
	if err != nil {
		log.Fatalf("Failed to listen on port %v %v", serverPort, err)
	}

	grpcServer := grpc.NewServer()
	srv := bakery.Server{}
	bakerypb.RegisterBakeryServiceServer(grpcServer, &srv)
	reflection.Register(grpcServer)

	log.Infof("Starting grpc server on %v ...", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start grpc Server %v", err)
	}
}
