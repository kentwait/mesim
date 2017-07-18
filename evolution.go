package mesim

import (
	sampler "mesim/sampler"
	utils "mesim/utils"
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

// SeqSpaceToFitSpace
func SeqSpaceToFitSpace(ancSeqSpace [][]int64, fitnessMatrix [][]float64, totalFitnessFunc func([]int64) float64, normalized bool) []float64 {
	var popSize, numSites int64
	fitnessSpace := make([]float64, len(ancSeqSpace))
	if len(fitnessMatrix) == 0 {
		panic("Length of fitnessMatrix must be greater than zero")
	} else {
		popSize = int64(len(fitnessMatrix))
		numSites = int64(len(fitnessMatrix[0]))
	}

	for i, seq := range ancSeqSpace {
		fitnessSpace[i] = totalFitnessFunc(seq)
	}

	if normalized == true {
		fitnessDenominator := utils.Sum(fitnessSpace...)
		for i := range fitnessSpace {
			fitnessSpace[i] = fitnessSpace[i] / fitnessDenominator
		}
	}
	return fitnessSpace
}

// ReplicateSelect
func ReplicateSelect(ancSeqSpace [][]int64, nextPopSize int64, fitnessMatrix [][]float64, totalFitnessFunc func(int64) int64) [][]int64 {

}
