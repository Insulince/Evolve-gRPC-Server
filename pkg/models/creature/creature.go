package creature_model

import (
	"evolve-rpc/pkg/pb"
	"fmt"
	"math/rand"
)

type Creature struct {
	*pb.CreatureMessage
}

func New() (creature *Creature) {
	return &Creature{
		&pb.CreatureMessage{
			Name:                            "",
			Generation:                      0,
			Speed:                           0,
			Stamina:                         0,
			Health:                          0,
			Greed:                           0,
			FitnessValue:                    0,
			SimulatedThisGeneration:         false,
			Outcome:                         "UNSET",
			NaturallySelectedThisGeneration: false,
		},
	}
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

func (c *Creature) NaturallySelect() () {
	// TODO: Implement correctly.

	if rand.Float64() < 0.5 {
		c.Outcome = "SUCCESS"
	} else {
		c.Outcome = "FAILURE"
	}
	c.NaturallySelectedThisGeneration = true
}

func (c *Creature) Kill() () {
	// TODO: Implement, if needed.
}

func (c *Creature) Reproduce() (offspring []*Creature) {
	child := New()

	child.Name = c.Name
	child.Generation = c.Generation + 1
	child.Speed = c.Speed
	child.Health = c.Health
	child.Stamina = c.Stamina
	child.Greed = c.Greed

	// TODO: Mutations.

	offspring = append(offspring, child)

	// TODO: Multiple children.

	return offspring
}
