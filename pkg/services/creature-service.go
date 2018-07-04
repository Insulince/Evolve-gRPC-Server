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
	"evolve-rpc/pkg/models/population"
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
				Outcome:                 "UNSET",
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

func (creatureService *CreatureService) NaturallySelectCreatureRpc(context context.Context, request *pb.NaturallySelectCreatureRpcRequest) (response *pb.NaturallySelectCreatureRpcResponse, err error) {
	log.Printf("NaturallySelectCreatureRpc: Interaction started.\n")
	defer log.Printf("NaturallySelectCreatureRpc: Interaction complete.\n")

	contextMetadata, ok := metadata.FromIncomingContext(context)
	if ok {
		log.Printf("NaturallySelectCreatureRpc: Metadata received: \"%v\".\n", contextMetadata)
	} else {
		log.Fatalf("NaturallySelectCreatureRpc: Unable to read metadata!\n")
	}

	log.Printf("NaturallySelectCreatureRpc: Request received: \"%v\".\n", request)

	creature := creature_model.FromMessage(request.CreatureMessage)

	creature.NaturallySelect(population_model.FromMessage(request.PopulationMessage))

	creatureMessage := creature_model.ToMessage(creature)

	log.Printf("NaturallySelectCreatureRpc: Sending response to client.\n")
	return &pb.NaturallySelectCreatureRpcResponse{
		CreatureMessage: creatureMessage,
	},
		nil
}

func (creatureService *CreatureService) NaturallySelectCreaturesRpc(context context.Context, request *pb.NaturallySelectCreaturesRpcRequest) (response *pb.NaturallySelectCreaturesRpcResponse, err error) {
	log.Printf("NaturallySelectCreaturesRpc: Interaction started.\n")
	defer log.Printf("NaturallySelectCreaturesRpc: Interaction complete.\n")

	contextMetadata, ok := metadata.FromIncomingContext(context)
	if ok {
		log.Printf("NaturallySelectCreaturesRpc: Metadata received: \"%v\".\n", contextMetadata)
	} else {
		log.Fatalf("NaturallySelectCreaturesRpc: Unable to read metadata!\n")
	}

	log.Printf("NaturallySelectCreaturesRpc: RpcRequest received: \"%v\".\n", request)

	var creatureMessages []*pb.CreatureMessage

	var wg sync.WaitGroup
	wg.Add(len(request.CreatureMessages))

	creatureMessageChannel := make(chan *pb.CreatureMessage)

	for _, creatureMessage := range request.CreatureMessages {
		go func(creatureMessage *pb.CreatureMessage) {
			creature := creature_model.FromMessage(creatureMessage)

			creature.NaturallySelect(population_model.FromMessage(request.PopulationMessage))

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

	return &pb.NaturallySelectCreaturesRpcResponse{
		CreatureMessages: creatureMessages,
	},
		nil
}

func (creatureService *CreatureService) KillFailedCreatureRpc(context context.Context, request *pb.KillFailedCreatureRpcRequest) (response *pb.KillFailedCreatureRpcResponse, err error) {
	log.Printf("KillFailedCreatureRpc: Interaction started.\n")
	defer log.Printf("KillFailedCreatureRpc: Interaction complete.\n")

	contextMetadata, ok := metadata.FromIncomingContext(context)
	if ok {
		log.Printf("KillFailedCreatureRpc: Metadata received: \"%v\".\n", contextMetadata)
	} else {
		log.Fatalf("KillFailedCreatureRpc: Unable to read metadata!\n")
	}

	log.Printf("KillFailedCreatureRpc: Request received: \"%v\".\n", request)

	creature := creature_model.FromMessage(request.CreatureMessage)

	creature.Kill()

	log.Printf("KillFailedCreatureRpc: Sending response to client.\n")
	return &pb.KillFailedCreatureRpcResponse{
	},
		nil
}

func (creatureService *CreatureService) KillFailedCreaturesRpc(context context.Context, request *pb.KillFailedCreaturesRpcRequest) (response *pb.KillFailedCreaturesRpcResponse, err error) {
	log.Printf("KillFailedCreaturesRpc: Interaction started.\n")
	defer log.Printf("KillFailedCreaturesRpc: Interaction complete.\n")

	contextMetadata, ok := metadata.FromIncomingContext(context)
	if ok {
		log.Printf("KillFailedCreaturesRpc: Metadata received: \"%v\".\n", contextMetadata)
	} else {
		log.Fatalf("KillFailedCreaturesRpc: Unable to read metadata!\n")
	}

	log.Printf("KillFailedCreaturesRpc: RpcRequest received: \"%v\".\n", request)

	var wg sync.WaitGroup
	wg.Add(len(request.CreatureMessages))

	creatureMessageChannel := make(chan struct{})

	for _, creatureMessage := range request.CreatureMessages {
		go func(creatureMessage *pb.CreatureMessage) {
			creature := creature_model.FromMessage(creatureMessage)

			creature.Kill()

			creatureMessageChannel <- struct{}{}
		}(creatureMessage)
	}

	go func() {
		for range creatureMessageChannel {
			wg.Done()
		}
	}()

	wg.Wait()

	return &pb.KillFailedCreaturesRpcResponse{
	},
		nil
}

func (creatureService *CreatureService) ReproduceSuccessfulCreatureRpc(context context.Context, request *pb.ReproduceSuccessfulCreatureRpcRequest) (response *pb.ReproduceSuccessfulCreatureRpcResponse, err error) {
	log.Printf("ReproduceSuccessfulCreatureRpc: Interaction started.\n")
	defer log.Printf("ReproduceSuccessfulCreatureRpc: Interaction complete.\n")

	contextMetadata, ok := metadata.FromIncomingContext(context)
	if ok {
		log.Printf("ReproduceSuccessfulCreatureRpc: Metadata received: \"%v\".\n", contextMetadata)
	} else {
		log.Fatalf("ReproduceSuccessfulCreatureRpc: Unable to read metadata!\n")
	}

	log.Printf("ReproduceSuccessfulCreatureRpc: Request received: \"%v\".\n", request)

	creature := creature_model.FromMessage(request.CreatureMessage)

	offspring := creature.Reproduce()

	var creatureMessages []*pb.CreatureMessage
	for _, child := range offspring {
		creatureMessages = append(creatureMessages, creature_model.ToMessage(child))
	}

	log.Printf("ReproduceSuccessfulCreatureRpc: Sending response to client.\n")
	return &pb.ReproduceSuccessfulCreatureRpcResponse{
		CreatureMessages: creatureMessages,
	},
		nil
}

func (creatureService *CreatureService) ReproduceSuccessfulCreaturesRpc(context context.Context, request *pb.ReproduceSuccessfulCreaturesRpcRequest) (response *pb.ReproduceSuccessfulCreaturesRpcResponse, err error) {
	log.Printf("ReproduceSuccessfulCreaturesRpc: Interaction started.\n")
	defer log.Printf("ReproduceSuccessfulCreaturesRpc: Interaction complete.\n")

	contextMetadata, ok := metadata.FromIncomingContext(context)
	if ok {
		log.Printf("ReproduceSuccessfulCreaturesRpc: Metadata received: \"%v\".\n", contextMetadata)
	} else {
		log.Fatalf("ReproduceSuccessfulCreaturesRpc: Unable to read metadata!\n")
	}

	log.Printf("ReproduceSuccessfulCreaturesRpc: RpcRequest received: \"%v\".\n", request)

	var offspring []*pb.CreatureMessage

	var wg sync.WaitGroup
	wg.Add(len(request.CreatureMessages))

	creatureMessagesChannel := make(chan []*pb.CreatureMessage)

	for _, creatureMessage := range request.CreatureMessages {
		go func(creatureMessage *pb.CreatureMessage) {
			creature := creature_model.FromMessage(creatureMessage)

			offspring := creature.Reproduce()

			var creatureMessages []*pb.CreatureMessage
			for _, child := range offspring {
				creatureMessages = append(creatureMessages, creature_model.ToMessage(child))
			}
			creatureMessagesChannel <- creatureMessages
		}(creatureMessage)
	}

	go func() {
		for creatureMessages := range creatureMessagesChannel {
			offspring = append(offspring, creatureMessages...)
			wg.Done()
		}
	}()

	wg.Wait()

	return &pb.ReproduceSuccessfulCreaturesRpcResponse{
		CreatureMessages: offspring,
	},
		nil
}
