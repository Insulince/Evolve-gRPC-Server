package services

import (
	"golang.org/x/net/context"
	"log"
	"google.golang.org/grpc/metadata"
	"math/rand"
	"evolve-rpc/pkg/pb"
	"evolve-rpc/pkg"
	"sync"
	"evolve-rpc/pkg/models/creature"
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
			Name:                    name,
			Generation:              generation,
			Speed:                   speed,
			Stamina:                 stamina,
			Health:                  health,
			Greed:                   greed,
			FitnessValue:            0,
			SimulatedThisGeneration: false,
		},
	}, nil
}

func (creatureService *CreatureService) GenerateCreaturesRpc(context context.Context, request *pb.GenerateCreaturesRpcRequest) (response *pb.GenerateCreaturesRpcResponse, err error) {
	log.Printf("GenerateCreaturesRpc: Interaction started.\n")
	defer log.Printf("GenerateCreaturesRpc: Interaction complete.\n")

	contextMetadata, ok := metadata.FromIncomingContext(context)
	if ok {
		log.Printf("GenerateCreaturesRpc: Metadata received: \"%v\".\n", contextMetadata)
	} else {
		log.Fatalf("GenerateCreaturesRpc: Unable to read metadata!\n")
	}

	log.Printf("GenerateCreaturesRpc: RpcRequest received: \"%v\".\n", request)

	var creatureMessages []*pb.CreatureMessage

	var wg sync.WaitGroup
	wg.Add(int(request.Quantity))

	creatureMessageChannel := make(chan *pb.CreatureMessage)

	for i := int64(0); i < request.Quantity; i++ {
		go func() {
			name := util.GenerateRandomName()
			generation := int64(0)
			speed := rand.Float64()
			stamina := rand.Float64()
			health := rand.Float64()
			greed := rand.Float64()

			creatureMessageChannel <- &pb.CreatureMessage{
				Name:                    name,
				Generation:              generation,
				Speed:                   speed,
				Stamina:                 stamina,
				Health:                  health,
				Greed:                   greed,
				FitnessValue:            0,
				SimulatedThisGeneration: false,
			}

		}()
	}

	go func() {
		for creatureMessage := range creatureMessageChannel {
			creatureMessages = append(creatureMessages, creatureMessage)
			wg.Done()
		}
	}()

	wg.Wait()

	return &pb.GenerateCreaturesRpcResponse{
		CreatureMessages: creatureMessages,
	}, nil
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

	creature := creature_model.FromMessage(request.CreatureMessage)

	creature.Simulate()

	creatureMessage := creature_model.ToMessage(creature)

	log.Printf("SimulateCreatureRpc: Sending response to client.\n")
	return &pb.SimulateCreatureRpcResponse{
		CreatureMessage: creatureMessage,
	},
		nil
}

func (creatureService *CreatureService) SimulateCreaturesRpc(context context.Context, request *pb.SimulateCreaturesRpcRequest) (response *pb.SimulateCreaturesRpcResponse, err error) {
	log.Printf("SimulateCreaturesRpc: Interaction started.\n")
	defer log.Printf("SimulateCreaturesRpc: Interaction complete.\n")

	contextMetadata, ok := metadata.FromIncomingContext(context)
	if ok {
		log.Printf("SimulateCreaturesRpc: Metadata received: \"%v\".\n", contextMetadata)
	} else {
		log.Fatalf("SimulateCreaturesRpc: Unable to read metadata!\n")
	}

	log.Printf("SimulateCreaturesRpc: RpcRequest received: \"%v\".\n", request)

	var creatureMessages []*pb.CreatureMessage

	var wg sync.WaitGroup
	wg.Add(len(request.CreatureMessages))

	creatureMessageChannel := make(chan *pb.CreatureMessage)

	for _, creatureMessage := range request.CreatureMessages {
		go func(creatureMessage *pb.CreatureMessage) {
			creature := creature_model.FromMessage(creatureMessage)

			creature.Simulate()

			creatureMessage = creature_model.ToMessage(creature)

			creatureMessageChannel <- creatureMessage
		}(creatureMessage)
	}

	go func() {
		for creatureMessage := range creatureMessageChannel {
			creatureMessages = append(creatureMessages, creatureMessage)
			wg.Done()
		}
	}()

	wg.Wait()

	return &pb.SimulateCreaturesRpcResponse{
		CreatureMessages: creatureMessages,
	},
		nil
}
