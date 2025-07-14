package swarm

import (
	"math"
	"math/rand"
	"sync"
)

var Function = math.Sin

const (
	swarmSize = 30
	// 1 -> 5 -> 5 -> 1
	networkSize = 12
)

// This swarm will be the driving force behind the neurons.
type Swarm struct {
	networkCollection [swarmSize]particle // Size of the swarm
	bestParticle      particle
	mu                sync.Mutex
	shouldStop        bool

	inertia float64
	c1      float64
	c2      float64
}

// Each particle holds a neural network
type particle struct {
	fitness    float64
	weight     [networkSize]float64 // weight of each neuron
	bestWeight [networkSize]float64
	velocity   [networkSize]float64 // velocity of each neuron
}

func initParticle() particle {
	p := new(particle)
	for i := range networkSize {
		p.weight[i] = rand.Float64()
		p.velocity[i] = rand.Float64()
	}
	p.bestWeight = p.weight
	p.fitness = math.MaxFloat64
	return *p
}

// This function is a little 'learning algorithm specific', but nil can be passed in if you're
// not using PSO
func (p *particle) fitnessFunction(s *Swarm) {
	errorBuf := 0.0
	counter := 0.0
	for i := -3 * math.Pi; i < 3*math.Pi; i += 0.05 {
		predicted := p.runNetwork(float64(i))
		real := Function(i)

		errorBuf += (predicted - real) * (predicted - real)
		counter += 1
	}

	// Just 'mean' error works a little better right now

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
	// get ready for matrix multiplication
	x = bipolar(x * p.weight[0]) // 0

	var bufSingle float64
	for i := range 5 {
		bufSingle += bipolar(x * p.weight[i+1]) // 1, 2, 3, 4, 5: first layer

	}

	bufSingle = bipolar(bufSingle)

	bufSingle *= p.weight[6] // 6: last node

	return bufSingle
}

// TODO create more activiation functions and include them in various layers.
func bipolar(x float64) float64 {
	return float64((1 - math.Pow(math.E, float64(-x))) / (1 + math.Pow(math.E, float64(-x))))
}

func (s *Swarm) iterateSwarmConc() {
	var wg sync.WaitGroup
	for i := range swarmSize { // Spin up a go routine for each particle
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.networkCollection[i].updateVelocity(s)
			s.networkCollection[i].updateWeight()
			// Need a mutex lock for this function
			s.networkCollection[i].fitnessFunction(s)
		}()
	}
	wg.Wait()
}

func (s *Swarm) findBestParticle() particle {
	bestIndex := 0
	bestFitness := float64(math.MaxFloat64)
	for i := 0; i < swarmSize; i++ {
		if bf := s.networkCollection[i].fitness; bf < bestFitness {
			bestFitness = bf
			bestIndex = i
		}
	}
	return s.networkCollection[bestIndex]
}

func (p *particle) updateVelocity(s *Swarm) {
	r1 := rand.Float64()
	r2 := rand.Float64()

	for i := range p.velocity {
		vBuf := s.inertia*p.velocity[i] + s.c1*r1*(p.bestWeight[i]-p.weight[i]) +
			s.c2*r2*(s.bestParticle.weight[i]-p.weight[i])
		switch {
		case vBuf < -1:
			vBuf = -1
		case vBuf > 1:
			vBuf = 1

		}

		p.velocity[i] = vBuf
	}
}

func (p *particle) updateWeight() {
	for i := 0; i < len(p.weight); i++ {
		p.weight[i] = p.weight[i] + p.velocity[i]
	}
}
