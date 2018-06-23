package main

import (
	"log"
	"fmt"
	"net"
	"google.golang.org/grpc"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"net/http"
	"experimental/evolve-rpc/pkg/services"
	"experimental/evolve-rpc/pkg/pb"
	"experimental/evolve-rpc/pkg/configurations"
)

func main() () {
	log.Printf("Evolve API starting up.\n")
	defer log.Printf("Evolve API shut down.\n")

	config := configurations.GetConfigurations()

	listen, err := net.Listen(config.Protocol, fmt.Sprintf(":%v", config.Port))
	if err != nil {
		log.Fatalf("Error opening \"%v\" port \"%v\": \"%v\".\n", config.Protocol, config.Port, err.Error())
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCreatureServiceServer(grpcServer, new(services.CreatureService))

	wrappedGRPCServer := grpcweb.WrapServer(grpcServer)
	httpServer := http.Server{}
	httpServer.Handler = http.HandlerFunc(wrappedGRPCServer.ServeHTTP)

	log.Printf("Server started on \"%v\" port \"%v\".\n", config.Protocol, config.Port)
	err = httpServer.Serve(listen)

	if err != nil {
		log.Fatalf("gRPC-Web Server Serve error: \"%v\"\n", err.Error())
	}

	log.Printf("Evolve API shutting down.\n")
}