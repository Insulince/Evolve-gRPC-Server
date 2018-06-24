package services

import (
	"golang.org/x/net/context"
	"log"
	"google.golang.org/grpc/metadata"
	"math/rand"
	"evolve-rpc/pkg/pb"
	"evolve-rpc/pkg"
	"sync"
)

type CreatureService struct {
}

func (evolveService *CreatureService) GenerateCreatureRpc(context context.Context, request *pb.GenerateCreatureRpcRequest) (response *pb.GenerateCreatureRpcResponse, err error) {
	log.Printf("GenerateCreatureRpc: Interaction started.\n")
	defer log.Printf("GenerateCreatureRpc: Interaction complete.\n")

	contextMetadata, ok := metadata.FromIncomingContext(context)
	if ok {
		log.Printf("GenerateCreatureRpc: Metadata received: \"%v\".\n", contextMetadata)
	} else {
		log.Fatalf("GenerateCreatureRpc: Unable to read metadata!\n")
	}

	log.Printf("GenerateCreatureRpc: Request received: \"%v\".\n", request)

	name := util.GenerateRandomName()
	generation := int64(0)
	speed := rand.Float64()
	stamina := rand.Float64()
	health := rand.Float64()
	greed := rand.Float64()

	log.Printf("GenerateCreatureRpc: Sending response to client.\n")
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

func (evolveService *CreatureService) GenerateCreaturesRpc(request *pb.GenerateCreaturesRpcRequest, stream pb.CreatureService_GenerateCreaturesRpcServer) (err error) {
	log.Printf("GenerateCreaturesRpc:  Interaction started.\n")
	defer log.Printf("GenerateCreaturesRpc:  Interaction complete.\n")

	contextMetadata, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		log.Printf("GenerateCreaturesRpc:  Metadata received: \"%v\".\n", contextMetadata)
	} else {
		log.Fatalf("GenerateCreaturesRpc:  Unable to read metadata!\n")
	}

	log.Printf("GenerateCreaturesRpc:  RpcRequest received: \"%v\".\n", request)

	var wg sync.WaitGroup
	wg.Add(int(request.Quantity))
	for i := int64(0); i < request.Quantity; i++ {
		go func() {
			defer wg.Done()

			name := util.GenerateRandomName()
			generation := int64(0)
			speed := rand.Float64()
			stamina := rand.Float64()
			health := rand.Float64()
			greed := rand.Float64()

			stream.Send(
				&pb.GenerateCreaturesRpcResponse{
					Creature: &pb.CreatureMessage{
						Name:       name,
						Generation: generation,
						Speed:      speed,
						Stamina:    stamina,
						Health:     health,
						Greed:      greed,
					},
				},
			)
		}()
	}
	wg.Wait()

	return nil
}
