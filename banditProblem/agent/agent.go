package agent

import (
	"math/rand"

	"github.com/poteto0/go-zero/gonp"
)

type Agent struct {
	Epsilon float64   `json:"epsilon"`
	Qs      []float64 `json:"qs"`
	Ns      []int     `json:"ns"`
	Alpha   float64   `json:"alpha"`
}

func NewAgent(epsilon, alpha float64, n int) *Agent {
	return &Agent{
		Epsilon: epsilon,
		Qs:      make([]float64, n),
		Ns:      make([]int, n),
		Alpha:   alpha,
	}
}

func (a *Agent) Update(action int, reward float64) {
	a.Ns[action] += 1
	a.Qs[action] += a.Alpha * (reward - a.Qs[action])
}

func (a *Agent) SelectAction() int {
	randVal := rand.Float64()

	if randVal < a.Epsilon {
		return rand.Intn(len(a.Qs))
	}
	return gonp.MaxIndex(a.Qs)
}
