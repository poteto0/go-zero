//go:build ignore

package main

import (
	"math"
	"math/rand"

	gridworld "github.com/poteto0/go-zero/dp/gridWorld"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette/moreland"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type plottable struct {
	grid       [][]float64
	N          int
	M          int
	resolution float64
	minX       float64
	minY       float64
}

func (p plottable) Dims() (c, r int) {
	return p.N, p.M
}
func (p plottable) X(c int) float64 {
	return p.minX + float64(c)*p.resolution
}
func (p plottable) Y(r int) float64 {
	return p.minY + float64(r)*p.resolution
}
func (p plottable) Z(c, r int) float64 {
	return p.grid[c][r]
}

func main() {
	gridWorld := gridworld.DefaultGridWorld
	var V [][]float64
	for y := 0; y < gridWorld.Height(); y++ {
		var vRow []float64
		for x := 0; x < gridWorld.Width(); x++ {
			vRow = append(vRow, rand.Float64())
		}
		V = append(V, vRow)
	}

	threshold := 0.001
	gamma := 0.9
	for {
		var oldV [][]float64
		for y := 0; y < gridWorld.Height(); y++ {
			var oldVRow []float64
			for x := 0; x < gridWorld.Width(); x++ {
				oldVRow = append(oldVRow, V[y][x])
			}
			oldV = append(oldV, oldVRow)
		}
		gridWorld.Run(V, gamma)

		delta := 0.0
		for y := 0; y < gridWorld.Height(); y++ {
			for x := 0; x < gridWorld.Width(); x++ {
				t := math.Abs(V[y][x] - oldV[y][x])
				if delta < t {
					delta = t
				}
			}
		}

		if delta < threshold {
			break
		}
	}

	V4Plot := make([][]float64, 4)
	for i := 0; i < len(V4Plot); i++ {
		V4Plot[i] = make([]float64, 3)
	}

	for x := 0; x < gridWorld.Width(); x++ {
		for y := 0; y < gridWorld.Height(); y++ {
			V4Plot[x][gridWorld.Height()-y-1] = V[y][x]
		}
	}

	plotData := plottable{
		grid:       V4Plot,
		N:          gridWorld.Width(),
		M:          gridWorld.Height(),
		minX:       0.5,
		minY:       0.5,
		resolution: 1.0,
	}
	pal := moreland.SmoothBlueRed().Palette(255)
	h := plotter.NewHeatMap(plotData, pal)
	//h.Rasterized = true

	p := plot.New()
	p.Add(h)

	if err := p.Save(5*vg.Inch, 5*vg.Inch, "./out/fig.png"); err != nil {
		panic(err)
	}
}
