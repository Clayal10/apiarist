package parser

import (
	"github.com/Clayal10/mathGen/lib/network"
	"github.com/Clayal10/mathGen/lib/user"
)

func TakeUserInput(u user.UserInput) user.UserOutput {
	waitTime := network.PSOSineGen(u)

	out := user.UserOutput{
		Time: waitTime,
	}

	return out
}
