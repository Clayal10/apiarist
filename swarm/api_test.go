package swarm_test

import (
	"testing"
	"time"

	"github.com/Clayal10/mathGen/swarm"
)

func TestEndToEnd(t *testing.T) {
	t.Run("TestSuccessful", func(_ *testing.T) {
		u := &swarm.UserInput{
			Inertia: .9,
			CogCoef: .9,
			SocCoef: .9,
		}
		s := &swarm.Swarm{}
		s.InitSwarm(u)
		go s.PSOSineGen()
		time.Sleep(time.Millisecond * 500)
		data := s.GetValues()
		if len(data) == 0 {
			t.Fail()
		}
	})
}
