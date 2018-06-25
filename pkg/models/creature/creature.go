package creature_model

import (
	"evolve-rpc/pkg/pb"
	"fmt"
)

type Creature struct {
	*pb.CreatureMessage
}

func FromMessage(creatureMessage *pb.CreatureMessage) (creature Creature) {
	return Creature{
		CreatureMessage: creatureMessage,
	}
}

func ToMessage(creature Creature) (creatureMessage *pb.CreatureMessage) {
	return creature.CreatureMessage
}

func (c *Creature) Simulate() () {
	if !c.SimulatedThisGeneration {
		c.FitnessValue = c.Speed + c.Stamina + c.Health + c.Greed
		c.SimulatedThisGeneration = true
	} else {
		fmt.Println("Creature already simulated...")
		// TODO: Should throw an error.
	}
}

func (c *Creature) Reproduce() () {
	// TODO: Implement.

	c.Generation++
}
