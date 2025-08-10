package swarm

import (
	"math"
	"math/rand"
	"sync"

	"github.com/Clayal10/mathGen/lib/mat"
)

var Function = math.Sin

const (
	swarmSize = 30
	// first layer is 1, second and second to last is 1xheight, middle are height*width.
	// NOTE must be more than 4 layers.
	layers         = 5
	width          = 5
	height         = 5
	numberOfValues = ((layers - 4) * width * height) + (height * 2) + 2
)

// This swarm will be the driving force behind the neurons.
type Swarm struct {
	networkCollection [swarmSize]*particle // Size of the swarm
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
	weight     [layers]mat.Matrix // weight of each neuron
	bestWeight [layers]mat.Matrix
	velocity   [numberOfValues]float64 // velocity of each neuron
}

func initParticle() *particle {
	p := &particle{}

	p.weight[0] = mat.NewMatrix([]float64{rand.Float64()}, 1, 1)
	p.weight[1] = mat.NewMatrix(mat.NewRandomMatrixValues(height, 1))

	for i := range layers - 4 {
		p.weight[i+2] = mat.NewMatrix(mat.NewRandomMatrixValues(height, width))
	}

	p.weight[layers-2] = mat.NewMatrix(mat.NewRandomMatrixValues(height, 1))
	p.weight[layers-1] = mat.NewMatrix([]float64{rand.Float64()}, 1, 1)

	for i := range p.velocity {
		p.velocity[i] = rand.Float64()
	}

	p.bestWeight = p.weight
	p.fitness = math.MaxFloat64
	return p
}

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

	if s != nil {
		s.mu.Lock()
		if p.fitness < s.bestParticle.fitness {
			s.bestParticle = *p
		}
		s.mu.Unlock()
	}

}

func (p *particle) getWeightList(best bool) (list []float64) {
	if best {
		for _, mat := range p.bestWeight {
			list = append(list, mat.GetValueList()...)
		}
		return
	}
	for _, mat := range p.weight {
		list = append(list, mat.GetValueList()...)
	}
	return
}

func (p *particle) runNetwork(x float64) float64 {
	input := mat.NewMatrix([]float64{x}, 1, 1)
	for _, matrix := range p.weight {
		input = mat.Mul(input, matrix)
	}
	return input.Values[0][0]
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
	return *s.networkCollection[bestIndex]
}

func (p *particle) updateVelocity(s *Swarm) {
	r1 := rand.Float64()
	r2 := rand.Float64()
	weight := p.getWeightList(false)
	bestWeight := p.getWeightList(true)
	bestParticleWeight := p.getWeightList(true)

	for i := range p.velocity {
		vBuf := s.inertia*p.velocity[i] + s.c1*r1*(bestWeight[i]-weight[i]) +
			s.c2*r2*(bestParticleWeight[i]-weight[i])
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
	weight := p.getWeightList(false)
	for i := 0; i < len(weight); i++ {
		weight[i] = weight[i] + p.velocity[i]
	}

	for i, m := range p.weight {
		values := m.GetValueList()
		p.weight[i] = mat.NewMatrix(values, m.Height, m.Width)
	}
}
