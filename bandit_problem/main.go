package main

import (
	"github.com/poteto0/go-zero/bandit_problem/agent"
	"github.com/poteto0/go-zero/bandit_problem/bandit"
	"github.com/poteto0/go-zero/gonp"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

const (
	N    = 10
	STEP = 1000
	runN = 100
)

func main() {
	var allRates [][]float64

	for n := 0; n < runN; n++ {
		band := bandit.NewBandit(N)
		agent := agent.NewAgent(0.1, 0.8, N)

		rates := make([]float64, STEP)
		tot := 0.0
		for i := 0; i < STEP; i++ {
			action := agent.SelectAction()
			reward := band.Play(action)
			agent.Update(action, float64(reward))
			tot += float64(reward)
			rates[i] = tot / float64((i + 1))
		}

		allRates = append(allRates, rates)
	}

	pts := make(plotter.XYs, len(allRates[0]))

	aveRates := gonp.AverageHorizontal(allRates)

	for i, rate := range aveRates {
		pts[i].X = float64(i)
		pts[i].Y = rate
	}

	p := plot.New()

	if err := plotutil.AddLinePoints(p, pts); err != nil {
		panic(err)
	}

	if err := p.Save(5*vg.Inch, 5*vg.Inch, "./out/fig.png"); err != nil {
		panic(err)
	}
}
