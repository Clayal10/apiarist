package network

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

// SineGen is a placeholder for what the output of an estimated function will be.
func SineGen(input float64) float64 {
	swarm := swarm{}

	swarm.initSwarm()
	swarm.iterateSwarmNoConc()

	fd, err := os.Create("data-output/data.csv")
	if err != nil {
		fmt.Printf("Couldn't create csv file: %v", err)
		return swarm.bestParticle.runNetwork(input) // just return the value
	}
	defer fd.Close()

	writer := csv.NewWriter(fd)
	defer writer.Flush()

	var data [][]string
	for i := -3 * math.Pi; i < 3*math.Pi; i += 0.05 {
		data = append(data, []string{strconv.FormatFloat(i, 'g', -1, 64), strconv.FormatFloat(swarm.bestParticle.runNetwork(i), 'g', -1, 64)})
	}

	err = writer.WriteAll(data)
	if err != nil {
		fmt.Printf("Error writing data: %v", err)
	}
	return swarm.bestParticle.runNetwork(input)
}

// This function will take a particle, run the network on the interval, and send
// it to the JS.
func visualDisplay(p *particle) {

}
