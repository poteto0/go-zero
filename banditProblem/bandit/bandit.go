package bandit

import "math/rand"

type Bandit struct {
	Arms  int       `json:"arms"`
	Rates []float64 `json:"rates"`
}

func NewBandit(arms int) *Bandit {
	rates := make([]float64, arms)
	for i := range rates {
		rates[i] = rand.Float64()
	}

	return &Bandit{
		Arms:  arms,
		Rates: rates,
	}
}

func (b *Bandit) Play(armId int) int {
	rate := rand.Float64()
	noise := 0.1 * (0.5 - rand.Float64())
	b.Rates[armId] += noise
	if rate > b.Rates[armId] {
		return 1
	}

	return 0
}
