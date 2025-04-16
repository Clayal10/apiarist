package parser

import (
	"time"

	"github.com/Clayal10/mathGen/lib/network"
)

// Input will be which math function they want to use and the value to input
type UserInput struct {
	Function string
	InputVal float64
	Learning string
}

type UserOutput struct { // This struct will change with increased functionality. What we want to give back
	Function  string
	OutputVal float64
	Learning  string
	Time      time.Duration
}

func TakeUserInput(u UserInput) UserOutput {
	opeationResult, waitTime := network.PSOSineGen(u.InputVal)

	out := UserOutput{
		Function:  u.Function,
		OutputVal: opeationResult,
		Learning:  u.Learning,
		Time:      waitTime,
	}

	return out
}
