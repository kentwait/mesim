package mesim

import (
	sampler "mesim/sampler"

)

// EvolveExplicit
func EvolveExplicit(ancArray *[]int64, rateMatrix [][]float64) {
	for i, char := range *ancArray {
		(*ancArray)[i] = sampler.MultinomialWhere(int64(1), rateMatrix[char], int64(1))[0]
	}
}

// EvolveFast
func EvolveFast(ancArray *[]int64, mu float64, zeroedRateMatrix [][]float64) {
	mutCoords := sampler.PoissonMutCoords(mu, int64(len(*ancArray)), int64(1))
	var tmpArray []int64
	if len(mutCoords) > 0 {
		for _, yPos := range mutCoords[1] {
			tmpArray = append(tmpArray, (*ancArray)[yPos])
		}
		EvolveExplicit(&tmpArray, zeroedRateMatrix)
		for i, xPos := range mutCoords[0] {
			(*ancArray)[xPos] = tmpArray[i]
		}
	}

}