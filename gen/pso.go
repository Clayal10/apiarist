package gen

import (
	"encoding/binary"
	"math"
	"math/rand"
	"sync"
)

var (
	inertia = 0.95
	c1      = 0.5
	c2      = 0.5
)

const (
	swarmSize = 30
	// 1 -> 5 -> 5 -> 1
	networkSize = 12
)

// This swarm will be the driving force behind the neurons.
type swarm struct {
	networkCollection [swarmSize]particle // Size of the swarm
	bestParticle      particle
	mu                sync.Mutex
	shouldStop        bool
}

func (s *swarm) GetValues() (data []byte) {
	s.shouldStop = true
	for i := -3 * math.Pi; i < 3*math.Pi; i += 0.05 {
		buf := []byte{}
		floatOutput := s.bestParticle.runNetwork(i)
		binary.LittleEndian.PutUint64(buf[:], math.Float64bits(floatOutput))
		data = append(data, buf[:8]...)
	}
	return
}

func (s *swarm) iterateSwarmConc() {
	var wg sync.WaitGroup
	for i := range swarmSize { // Spin up a go routine for each particle
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.networkCollection[i].updateVelocity(s.bestParticle)
			s.networkCollection[i].updateWeight()
			// Need a mutex lock for this function
			s.networkCollection[i].fitnessFunction(s)
		}()
	}
	wg.Wait()
}

// These values are hard coded for the moment
func (s *swarm) initSwarm(u UserInput) {
	inertia = u.Inertia
	c1 = u.CogCoef
	c2 = u.SocCoef

	for i := 0; i < swarmSize; i++ {
		s.networkCollection[i] = initParticle()
	}
	s.bestParticle = s.findBestParticle()
}

func (s *swarm) findBestParticle() particle {
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

func (p *particle) updateVelocity(bestP particle) {
	r1 := rand.Float64()
	r2 := rand.Float64()

	for i := range p.velocity {
		vBuf := inertia*p.velocity[i] + c1*r1*(p.bestWeight[i]-p.weight[i]) +
			c2*r2*(bestP.weight[i]-p.weight[i])
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
