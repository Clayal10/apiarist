package network

import (
	"encoding/binary"
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/Clayal10/mathGen/lib/user"
)

// SineGen is a placeholder for what the output of an estimated function will be.
func PSOSineGen(u user.UserInput) time.Duration {

	swarm := swarm{}

	startTime := time.Now()
	swarm.initSwarm(u)
	//swarm.iterateSwarmNoConc()
	swarm.iterateSwarmConc() // About 4 times faster.

	endTime := time.Now()

	fd, err := os.Create("data-output/data.csv")
	if err != nil {
		fmt.Printf("Couldn't create csv file: %v", err)
		return endTime.Sub(startTime) // just return the value
	}
	defer fd.Close()

	writer := csv.NewWriter(fd)
	defer writer.Flush()

	var data [][]string
	var webData []byte
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
		webData = append(webData, buf[:8]...)
	}
	fmt.Println(webData)
	// Start serving the web socket
	go visualDisplay(webData)

	err = writer.WriteAll(data)
	if err != nil {
		fmt.Printf("Error writing data: %v", err)
	}

	return endTime.Sub(startTime)

}
