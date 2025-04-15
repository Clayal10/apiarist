package parser

import "github.com/Clayal10/mathGen/lib/network"

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
}

func TakeUserInput(u UserInput) UserOutput {
	opeationResult := network.SineGen(u.InputVal)

	out := UserOutput{
		Function:  u.Function,
		OutputVal: opeationResult,
		Learning:  u.Learning,
	}

	return out
}
