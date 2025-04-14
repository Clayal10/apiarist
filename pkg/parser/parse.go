package parser

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
	out := UserOutput{
		Function:  u.Function,
		OutputVal: u.InputVal,
		Learning:  u.Learning,
	}

	return out
}
