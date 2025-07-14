package swarm

import (
	"encoding/binary"
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
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

func (swarm *Swarm) PSOSineGen(u *UserInput) {

	for {
		if swarm.shouldStop {
			break
		}
		swarm.iterateSwarmConc()
	}

	fd, err := os.Create("data-output/data.csv")
	if err != nil {
		fmt.Printf("Couldn't create csv file: %v", err)
		return
	}
	defer fd.Close()

	writer := csv.NewWriter(fd)
	defer writer.Flush()

	var data [][]string
	for i := -3 * math.Pi; i < 3*math.Pi; i += 0.05 {
		floatOutput := swarm.bestParticle.runNetwork(i)

		data = append(
			data, []string{strconv.FormatFloat(i, 'g', -1, 64),
				strconv.FormatFloat(floatOutput, 'g', -1, 64),
				strconv.FormatFloat(math.Sin(i), 'g', -1, 64),
			})
		// Convert output to a byte slice and append to webData
		var buf [8]byte
		binary.LittleEndian.PutUint64(buf[:], math.Float64bits(floatOutput))
	}

	err = writer.WriteAll(data)
	if err != nil {
		fmt.Printf("Error writing data: %v", err)
	}
}

func (s *Swarm) GetValues() (data []float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.shouldStop = true
	for i := -3 * math.Pi; i < 3*math.Pi; i += 0.05 {
		data = append(data, s.bestParticle.runNetwork(i))
	}
	return
}
