package creature_model

import (
	"evolve-rpc/pkg/pb"
	"math/rand"
	"log"
	"evolve-rpc/pkg/models/population"
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

func FromMessage(creatureMessage *pb.CreatureMessage) (creature *Creature) {
	return &Creature{
		CreatureMessage: creatureMessage,
	}
}

func ToMessage(creature *Creature) (creatureMessage *pb.CreatureMessage) {
	return creature.CreatureMessage
}

func (c *Creature) Simulate() () {
	if !c.SimulatedThisGeneration {
		c.FitnessValue = c.Speed + c.Stamina + c.Health + c.Greed
		c.SimulatedThisGeneration = true
	} else {
		log.Fatalf("Creature \"%v\" already simulated...\n", c.Name)
	}
}

func (c *Creature) NaturallySelect(p *population_model.Population) () {
	if !c.NaturallySelectedThisGeneration {
		if p.Size < p.CarryingCapacity {
			if rand.Float64() >= float64(c.FitnessIndex)/float64(p.Size) {
				c.Outcome = "SUCCESS"
			} else {
				c.Outcome = "FAILURE"
			}
		} else {
			if rand.Float64() >= float64(c.FitnessIndex)/float64(p.CarryingCapacity) {
				c.Outcome = "SUCCESS"
			} else {
				c.Outcome = "FAILURE"
			}
		}
		c.NaturallySelectedThisGeneration = true
	} else {
		log.Fatalf("Creature \"%v\" already naturall selected...\n", c.Name)
	}
}

func (c *Creature) Kill() () {
	// Currently there's nothing to be done when a creature is killed, it just "fails" to be sent back to the UI. This function is included for completeness.
}

func (c *Creature) Reproduce() (offspring []*Creature) {
	quantityChildren := rand.Intn(3) + 1 // Reproduce between 1 and 3 children (inclusive: [1, 3]).

	for i := 0; i < quantityChildren; i++ {
		child := New()

		child.Name = c.Name
		child.Generation = c.Generation + 1
		child.Speed = c.Speed
		child.Health = c.Health
		child.Stamina = c.Stamina
		child.Greed = c.Greed

		// TODO: Mutations.

		offspring = append(offspring, child)
	}

	return offspring
}
