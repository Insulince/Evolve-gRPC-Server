package services

import (
	"../pb"
	"golang.org/x/net/context"
	"log"
	"google.golang.org/grpc/metadata"
)

type EvolveService struct {
}

func (evolveService *EvolveService) Evolve(context context.Context, request *pb.EvolveRequest) (response *pb.EvolveResponse, err error) {
	log.Printf("Evolve: Interaction started.\n")
	defer log.Printf("Evolve: Interaction complete.\n")

	contextMetadata, ok := metadata.FromIncomingContext(context)
	if ok {
		log.Printf("Evolve: Metadata received: \"%v\".\n", contextMetadata)
	} else {
		log.Fatalf("Evolve: Unable to read metadata!\n")
	}

	log.Printf("Evolve: Request received: \"%v\".\n", request)

	log.Printf("Evolve: Sending response to client.\n")
	return &pb.EvolveResponse{Success: true}, nil
}
