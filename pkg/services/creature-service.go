package services

import (
	"golang.org/x/net/context"
	"log"
	"google.golang.org/grpc/metadata"
	"math/rand"
	"experimental/evolve-rpc/pkg/models/creature"
	"experimental/evolve-rpc/pkg/pb"
)

type CreatureService struct {
}


func (evolveService *CreatureService) GenerateCreatureRpc(context context.Context, request *pb.GenerateCreatureRpcRequest) (response *pb.GenerateCreatureRpcResponse, err error) {
	log.Printf("Creature: Interaction started.\n")
	defer log.Printf("Creature: Interaction complete.\n")

	contextMetadata, ok := metadata.FromIncomingContext(context)
	if ok {
		log.Printf("Creature: Metadata received: \"%v\".\n", contextMetadata)
	} else {
		log.Fatalf("Creature: Unable to read metadata!\n")
	}

	log.Printf("Creature: RpcRequest received: \"%v\".\n", request)

	name := ""
	generation := int64(0)
	speed := 0.1
	stamina := 0.1
	health := 0.1
	greed := 0.1

	log.Printf("Creature: Sending response to client.\n")
	return &pb.GenerateCreatureRpcResponse{
		Creature: &pb.CreatureMessage{
			Name:       name,
			Generation: generation,
			Speed:      speed,
			Stamina:    stamina,
			Health:     health,
			Greed:      greed,
		},
	}, nil
}
