package network

import (
	"math"
	"math/rand"
)

var Function = math.Sin

// Each particle holds a neural network
type particle struct {
	fitness    float64
	weight     [networkSize]float64 // weight of each neuron
	bestWeight [networkSize]float64
	velocity   [networkSize]float64 // velocity of each neuron
}

func initParticle() particle {
	p := new(particle)
	for i := 0; i < networkSize; i++ {
		p.weight[i] = rand.Float64()
		p.velocity[i] = rand.Float64()
	}
	p.bestWeight = p.weight
	return *p
}

func (p *particle) fitnessFunction(s *swarm) {
	errorBuf := 0.0
	counter := 0.0
	for i := -3 * math.Pi; i < 3*math.Pi; i += 0.05 {
		predicted := p.runNetwork(float64(i))
		real := Function(i)

		errorBuf += (predicted - real) * (predicted - real)
		counter += 1
	}

	errorBuf = errorBuf / counter
	if errorBuf < p.fitness {
		p.bestWeight = p.weight
	}
	p.fitness = errorBuf

	if s != nil { // Should be able to reuse outside of PSO
		s.mu.Lock()
		if p.fitness < s.bestParticle.fitness {
			s.bestParticle = *p
		}
		s.mu.Unlock()
	}

}

func (p *particle) runNetwork(x float64) float64 {
	// There are specific hard-coded numbers that depend on the network size
	x = bipolar(x * p.weight[0]) // 0

	var bufRow [3]float64
	for i := range bufRow {
		bufRow[i] = bipolar(x * p.weight[i+1]) // 1, 2, 3: first layer

	}

	var bufSingle float64 = 0
	for i := range bufRow {
		bufSingle += bufRow[i] * p.weight[i+4] // 4, 5, 6: second layer
	}
	bufSingle = bipolar(bufSingle)

	bufSingle *= p.weight[7] // 7: last node

	return bufSingle
}

func bipolar(x float64) float64 {
	return float64((1 - math.Pow(math.E, float64(-x))) / (1 + math.Pow(math.E, float64(-x))))
}
