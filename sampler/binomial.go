package sampler

// BinomialSample
func BinomialSample(n int, p float64) int {
	pArray := []float64{p, 1 - p}
	result := generalMultinomial(n, pArray, false)
	return result[0]
}

// BinomialMutCoords
func BinomialMutCoords(mu float64, nSites, popSize int) [][]int {
	var xArray, yArray, value []int
	var v int
	for i := 0; i < popSize; i++ {
		for j := 0; j < nSites; j++ {
			v = BinomialSample(1, mu)
			if v > 0 {
				xArray = append(xArray, i)
				yArray = append(yArray, j)
				value = append(value, v)
			}

		}
	}
	result := [][]int{xArray, yArray, value}
	return result
}
