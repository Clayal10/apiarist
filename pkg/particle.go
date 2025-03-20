package network

import (
	"math"
	"math/rand"
)

const (
	iterations = 1000
	swarmSize  = 30
	inertia    = 0.95
	c1         = 0.5
	c2         = 0.5
	// 1 -> 3 -> 3 -> 1
	networkSize = 8
)

// This swarm will be the driving force behind the neurons.
type swarm struct {
	networkCollection [swarmSize]particle // Size of the swarm
	bestParticle      particle
}

// These values are hard coded for the moment
func (s *swarm) initSwarm() {
	for i := 0; i < swarmSize; i++ {
		s.networkCollection[i] = initParticle()
	}
	s.bestParticle = s.findBestParticle()
}

func (s *swarm) findBestParticle() particle {
	bestIndex := 0
	bestFitness := float32(math.MaxFloat32)
	for i := 0; i < swarmSize; i++ {
		if bf := s.networkCollection[i].fitness; bf < bestFitness {
			bestFitness = bf
			bestIndex = i
		}
	}
	return s.networkCollection[bestIndex]
}

// Each particle holds a neural network
type particle struct {
	fitness    float32
	weight     [networkSize]float32 // weight of each neuron
	bestWeight [networkSize]float32
	velocity   [networkSize]float32 // velocity of each neuron
}

func initParticle() particle {
	p := new(particle)
	for i := 0; i < networkSize; i++ {
		p.weight[i] = rand.Float32()
		p.velocity[i] = rand.Float32()
	}
	p.bestWeight = p.weight
	return *p
}

func (p *particle) updateVelocity() {

}

func (p *particle) updateWeight() {

}

func (p *particle) fitnessFunction() {

}
