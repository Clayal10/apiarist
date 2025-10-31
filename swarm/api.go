package swarm

import (
	"math"
	"time"
)

type UserInput struct {
	Inertia float64
	// How much personal best impacts the particle.
	CogCoef float64
	// How much global best impacts the particle.
	SocCoef float64
}

type UserOutput struct {
	Time time.Duration
}

// These values are hard coded for the moment
func (s *Swarm) InitSwarm(u *UserInput) {
	s.inertia = u.Inertia
	s.c1 = u.CogCoef
	s.c2 = u.SocCoef

	for i := 0; i < swarmSize; i++ {
		s.networkCollection[i] = initParticle()
	}
	s.bestParticle = s.findBestParticle()
}

func (swarm *Swarm) PSOSineGen() {
	for {
		if swarm.shouldStop {
			break
		}
		swarm.iterateSwarmConc()
	}
}

func (s *Swarm) GetValues() (data []float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p := s.findBestParticle()
	for i := -3 * math.Pi; i < 3*math.Pi; i += 0.05 {
		data = append(data, p.runNetwork(i))
	}
	return
}

func (s *Swarm) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.shouldStop = true
}
