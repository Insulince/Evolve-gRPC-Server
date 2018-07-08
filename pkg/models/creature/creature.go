package creature_model

import (
	"evolve-rpc/pkg/pb"
	"math/rand"
	"log"
	"evolve-rpc/pkg/models/population"
	"evolve-rpc/pkg"
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
			ChanceOfMutation:                0,
			FitnessValue:                    0,
			SimulatedThisGeneration:         false,
			Outcome:                         "UNSET",
			NaturallySelectedThisGeneration: false,
			FitnessIndex:                    0,
			MutatedThisGeneration:           false,
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
		speed := c.Speed
		stamina := c.Stamina
		health := c.Health
		greed := c.Greed

		distance := 0.0
		for i := 0; i < 100; i++ {
			if stamina > 0 {
				stamina -= 0.05
				distance += speed
			} else {
				if float64(i % 10) / 10.0 < greed + 0.01 {
					health -= 0.5
					distance += speed

					if health <= 0 {
						c.FitnessValue = 0
						c.SimulatedThisGeneration = true
						return
					}
				} else {
					// TODO: Take a break instead of oscillating between tired and not tired.
					stamina += 0.05
				}
			}
		}

		c.FitnessValue = distance / 100.0
		c.SimulatedThisGeneration = true
	} else {
		log.Fatalf("Creature \"%v\" already simulated...\n", c.Name)
	}
}

func (c *Creature) NaturallySelect(p *population_model.Population) () {
	if !c.NaturallySelectedThisGeneration {
		if c.FitnessValue > 0 {
			//if p.Size < p.CarryingCapacity {
			//	if rand.Float64() >= float64(c.FitnessIndex)/float64(p.Size) {
			//		c.Outcome = "SUCCESS"
			//	} else {
			//		c.Outcome = "FAILURE"
			//	}
			//} else {
				if rand.Float64() >= float64(c.FitnessIndex)/float64(p.CarryingCapacity) {
					c.Outcome = "SUCCESS"
				} else {
					c.Outcome = "FAILURE"
				}
			//}
		} else {
			c.Outcome = "FAILURE"
		}
		c.NaturallySelectedThisGeneration = true
	} else {
		log.Fatalf("Creature \"%v\" already naturall selected...\n", c.Name)
	}
}

func (c *Creature) Kill() () {
	// Currently there's nothing to be done when a creature is killed, it just "fails" to be sent back to the UI. This function is included for completeness.
}

// TODO: Refactor.
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
		child.ChanceOfMutation = c.ChanceOfMutation

		const MAX_DELTA = 0.1
		mutations := 0
		for rand.Float64() < c.ChanceOfMutation && mutations < 5 {
			attributeIndex := rand.Intn(5)
			mutationAmount := rand.Float64() * MAX_DELTA
			for mutationAmount == 0 {
				mutationAmount = rand.Float64() * MAX_DELTA
			}
			if rand.Float64() < 0.5 {
				mutationAmount *= -1
			}

			validMutation := true
			switch attributeIndex {
			case 0:
				child.Speed += mutationAmount
				if child.Speed > 1 {
					child.Speed = 1
					validMutation = false
				} else if child.Speed < 0 {
					child.Speed = 0
					validMutation = false
				}
			case 1:
				child.Health += mutationAmount
				if child.Health > 1 {
					child.Health = 1
					validMutation = false
				} else if child.Health < 0 {
					child.Health = 0
					validMutation = false
				}
			case 2:
				child.Stamina += mutationAmount
				if child.Stamina > 1 {
					child.Stamina = 1
					validMutation = false
				} else if child.Stamina < 0 {
					child.Stamina = 0
					validMutation = false
				}
			case 3:
				child.Greed += mutationAmount
				if child.Greed > 1 {
					child.Greed = 1
					validMutation = false
				} else if child.Greed < 0 {
					child.Greed = 0
					validMutation = false
				}
			case 4:
				child.ChanceOfMutation += mutationAmount
				if child.ChanceOfMutation > 1 {
					child.ChanceOfMutation = 1
					validMutation = false
				} else if child.ChanceOfMutation < 0.001 {
					child.ChanceOfMutation = 0.001
					validMutation = false
				}
			default:
				log.Fatalf("Unrecognized attributeIndex amount: %v\n", attributeIndex)
			}

			if validMutation {
				mutations++
				if child.MutatedThisGeneration == false {
					child.Generation = 0
					child.Name = util.MutateName(c.Name)
					child.MutatedThisGeneration = true
				}
			}
		}

		offspring = append(offspring, child)
	}

	return offspring
}
