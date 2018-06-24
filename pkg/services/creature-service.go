package services

import (
	"golang.org/x/net/context"
	"log"
	"google.golang.org/grpc/metadata"
	"math/rand"
	"evolve-rpc/pkg/pb"
	"evolve-rpc/pkg"
	"sync"
	"io"
)

type CreatureService struct {
}

func (creatureService *CreatureService) GenerateCreatureRpc(context context.Context, request *pb.GenerateCreatureRpcRequest) (response *pb.GenerateCreatureRpcResponse, err error) {
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
		CreatureMessage: &pb.CreatureMessage{
			Name:       name,
			Generation: generation,
			Speed:      speed,
			Stamina:    stamina,
			Health:     health,
			Greed:      greed,
		},
	}, nil
}

func (creatureService *CreatureService) GenerateCreaturesRpc(request *pb.GenerateCreaturesRpcRequest, stream pb.CreatureService_GenerateCreaturesRpcServer) (err error) {
	log.Printf("GenerateCreaturesRpc: Interaction started.\n")
	defer log.Printf("GenerateCreaturesRpc: Interaction complete.\n")

	contextMetadata, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		log.Printf("GenerateCreaturesRpc: Metadata received: \"%v\".\n", contextMetadata)
	} else {
		log.Fatalf("GenerateCreaturesRpc: Unable to read metadata!\n")
	}

	log.Printf("GenerateCreaturesRpc: RpcRequest received: \"%v\".\n", request)

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
					CreatureMessage: &pb.CreatureMessage{
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

func (creatureService *CreatureService) SimulateCreatureRpc(context context.Context, request *pb.SimulateCreatureRpcRequest) (response *pb.SimulateCreatureRpcResponse, err error) {
	log.Printf("SimulateCreatureRpc: Interaction started.\n")
	defer log.Printf("SimulateCreatureRpc: Interaction complete.\n")

	contextMetadata, ok := metadata.FromIncomingContext(context)
	if ok {
		log.Printf("SimulateCreatureRpc: Metadata received: \"%v\".\n", contextMetadata)
	} else {
		log.Fatalf("SimulateCreatureRpc: Unable to read metadata!\n")
	}

	log.Printf("SimulateCreatureRpc: Request received: \"%v\".\n", request)

	// TODO: Do Business Logic Here.

	log.Printf("SimulateCreatureRpc: Sending response to client.\n")
	return &pb.SimulateCreatureRpcResponse{
		CreatureMessage: request.CreatureMessage,
	}, nil
}

func (creatureService *CreatureService) SimulateCreaturesRpc(stream pb.CreatureService_SimulateCreaturesRpcServer) (err error) {
	log.Printf("SimulateCreaturesRpc: Interaction started.\n")
	defer log.Printf("SimulateCreaturesRpc: Interaction complete.\n")

	contextMetadata, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		log.Printf("SimulateCreaturesRpc: Metadata received: \"%v\".\n", contextMetadata)
	} else {
		log.Fatalf("SimulateCreaturesRpc: Unable to read metadata!\n")
	}

	log.Printf("SimulateCreaturesRpc: BiDirectional RpcRequest received, opening client stream.\n")

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			log.Printf("SimulateCreaturesRpc: Client stream closed.\n")
			break
		}
		if err != nil {
			log.Printf("SimulateCreaturesRpc: Request error: \"%v\".\n", err.Error())
			return err
		}
		log.Printf("SimulateCreaturesRpc: RpcRequest received: \"%v\".\n", request)

		// TODO: Do Business Logic Here.

		log.Printf("SimulateCreaturesRpc: Sending response to cllient.")
		stream.Send(
			&pb.SimulateCreaturesRpcResponse{
				CreatureMessage: request.CreatureMessage,
			},
		)
	}

	log.Printf("SimulateCreaturesRpc: Closing server stream.\n")
	return nil
}
