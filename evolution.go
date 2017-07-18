package mesim

import (
	rand "math/rand"
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

// FitnessFunc
type FitnessFunc func([]int64) float64

// SeqSpaceToFitSpace
func SeqSpaceToFitSpace(ancSeqSpace [][]int64, fitnessMatrix [][]float64, totalFitnessFunc FitnessFunc, normalized bool) []float64 {
	if len(ancSeqSpace) == 0 {
		panic("Length of ancSeqSpace must be greater than zero")
	} else {
		if len(ancSeqSpace[0]) == 0 {
			panic("Length of rows in ancSeqSpace must be greater than zero")
		}
	}
	if len(fitnessMatrix) == 0 {
		panic("Length of fitnessMatrix must be greater than zero")
	}

	fitnessSpace := make([]float64, len(ancSeqSpace))

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
func ReplicateSelect(ancSeqSpace [][]int64, nextPopSize int64, fitnessMatrix [][]float64, totalFitnessFunc FitnessFunc) [][]int64 {
	normedFitSpace := SeqSpaceToFitSpace(ancSeqSpace, fitnessMatrix, totalFitnessFunc, true)
	ancSeqSpaceCnts := sampler.Multinomial(nextPopSize, normedFitSpace)

	newSeqSpace := make([][]int64, len(ancSeqSpace))
	for ancPos, cnt := range ancSeqSpaceCnts {
		for i := int64(0); i < cnt; i++ {
			newSeqSpace[i] = ancSeqSpace[ancPos]
		}
	}
	return newSeqSpace
}

// MutateSeqSpace
func MutateSeqSpace(seqSpace *[][]int64, mu float64, rateMatrix [][]float64) {
	if len(*seqSpace) == 0 {
		panic("Length of ancSeqSpace must be greater than zero")
	} else {
		if len((*seqSpace)[0]) == 0 {
			panic("Length of rows in ancSeqSpace must be greater than zero")
		}
	}
	popSize := len(*seqSpace)
	numSites := len((*seqSpace)[0])
	muPerSeq := mu * float64(numSites)

	// Returns two arrays, array[0] is sequence ID, array[1] always 0, array[2] is numbe rof hits
	hitsPerSeq := sampler.PoissonMutCoords(muPerSeq, int64(popSize), int64(1))

	var permSites, seqIdx []int
	for i, hits := range hitsPerSeq[2] {
		permSites = rand.Perm(numSites)
		seqIdx = int(hitsPerSeq[0][i])
		for _, siteIdx := range permSites[:hits] {
			char = (*seqSpace)[seqIdx][siteIdx]
			newChar = sampler.MultinomialWhere(int64(1), rateMatrix[char], int64(1))[0]
			(*seqSpace)[seqIdx][siteIdx] = newChar
		}
	}
}
