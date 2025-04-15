package network

// SineGen is a placeholder for what the output of an estimated function will be.
func SineGen(input float64) float64 {
	swarm := swarm{}

	swarm.initSwarm()
	swarm.iterateSwarmNoConc()

	return swarm.bestParticle.runNetwork(input)
}

// This function will take a particle, run the network on the interval, and send
// it to the JS.
func visualDisplay(p *particle) {

}
