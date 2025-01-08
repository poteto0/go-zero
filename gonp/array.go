package gonp

func Average(targets []float64) float64 {
	sum := 0.0
	for _, target := range targets {
		sum += target
	}
	return sum / float64(len(targets))
}

func AverageHorizontal(targets [][]float64) []float64 {
	vals := make([]float64, len(targets[0]))
	for i := range targets {
		for j := range targets[i] {
			vals[j] += (targets[i][j] - vals[j]) / float64(i+1)
		}
	}

	return vals
}

func MaxIndex(targets []float64) int {
	maxId := 0
	maxVal := 0.0

	for id, target := range targets {
		if target >= maxVal {
			maxVal = target
			maxId = id
		}
	}

	return maxId
}
