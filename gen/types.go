package gen

import "time"

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
