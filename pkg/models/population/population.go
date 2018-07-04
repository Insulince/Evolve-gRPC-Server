package population_model

import (
	"evolve-rpc/pkg/pb"
)

type Population struct {
	*pb.PopulationMessage
}

func New() (population *Population) {
	return &Population{
		&pb.PopulationMessage{
			Size: 0,
		},
	}
}

func FromMessage(populationMessage *pb.PopulationMessage) (population *Population) {
	return &Population{
		PopulationMessage: populationMessage,
	}
}

func ToMessage(population Population) (populationMessage *pb.PopulationMessage) {
	return population.PopulationMessage
}
