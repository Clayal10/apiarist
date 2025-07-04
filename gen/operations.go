package gen

import (
	"encoding/binary"
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

func PSOSineGen(swarm *Swarm, u *UserInput) {

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
