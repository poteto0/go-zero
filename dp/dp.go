//go:build ignore

package main

import (
	"fmt"
	"math"
)

func main() {
	V := map[string]float64{
		"L1": 0.0,
		"L2": 0.0,
	}

	T := map[string]float64{
		"L1": 0.0,
		"L2": 0.0,
	}

	cnt := 0
	for {
		T["L1"] = 0.5*(-1+0.9*V["L1"]) + 0.5*(1+0.9*V["L2"])
		delta := math.Abs(T["L1"] - V["L1"])
		V["L1"] = T["L1"]

		T["L2"] = 0.5*(0+0.9*V["L1"]) + 0.5*(-1+0.9*V["L2"])
		delta = math.Max(delta, math.Abs(T["L2"]-V["L2"]))
		V["L2"] = T["L2"]

		cnt++

		if delta < 0.0001 {
			break
		}
	}

	fmt.Println(V)
	fmt.Println(cnt)
}
