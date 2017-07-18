package mesim

import (
	"math"
	"math/rand"
	"mesim/sampler"
	"mesim/utils"
)

// EvolveChar
func EvolveChar(char int, rateMatrix [][]float64) (newChar int) {
	newChar = sampler.MultinomialWhere(1, rateMatrix[char], 1)[0]
	return
}

// EvolveExplicit
func EvolveExplicit(ancArray *[]int, rateMatrix [][]float64) {
	for i, char := range *ancArray {
		(*ancArray)[i] = sampler.MultinomialWhere(1, rateMatrix[char], 1)[0]
	}
}

// EvolveFast
func EvolveFast(ancArray *[]int, mu float64, zeroedRateMatrix [][]float64) {
	mutCoords := sampler.PoissonMutCoords(mu, len(*ancArray), 1)
	var tmpArray []int
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
type FitnessFunc func([]int) float64

// SeqSpaceToFitSpace
func SeqSpaceToFitSpace(ancSeqSpace [][]int, fitnessMatrix [][]float64, totalFitnessFunc FitnessFunc, normalized bool) []float64 {
	fitnessSpace := seqSpaceToFitSpace(ancSeqSpace, fitnessMatrix, totalFitnessFunc)
	if normalized == true {
		fitnessDenominator := utils.Sum(fitnessSpace...)
		for i := range fitnessSpace {
			fitnessSpace[i] = fitnessSpace[i] / fitnessDenominator
		}
	}
	return fitnessSpace
}

// SeqSpaceToLogFitSpace
func SeqSpaceToLogFitSpace(ancSeqSpace [][]int, fitnessMatrix [][]float64, totalFitnessFunc FitnessFunc, normalized bool) []float64 {
	fitnessSpace := seqSpaceToFitSpace(ancSeqSpace, fitnessMatrix, totalFitnessFunc)
	if normalized == true {
		fitnessDenominator := math.Log(utils.Sum(fitnessSpace...))
		for i := range fitnessSpace {
			fitnessSpace[i] = fitnessSpace[i] - fitnessDenominator
		}
	}
	return fitnessSpace
}

func seqSpaceToFitSpace(ancSeqSpace [][]int, fitnessMatrix [][]float64, totalFitnessFunc FitnessFunc) []float64 {
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
	return fitnessSpace
}

// ReplicateSelect
func ReplicateSelect(ancSeqSpace [][]int, nextPopSize int, fitnessMatrix [][]float64, totalFitnessFunc FitnessFunc) [][]int {
	normedFitSpace := SeqSpaceToFitSpace(ancSeqSpace, fitnessMatrix, totalFitnessFunc, true)
	ancSeqSpaceCnts := sampler.Multinomial(nextPopSize, normedFitSpace)

	newSeqSpace := make([][]int, len(ancSeqSpace))
	idxOffset := 0
	for ancPos, cnt := range ancSeqSpaceCnts {
		for i := 0 + idxOffset; i < cnt+idxOffset; i++ {
			newSeqSpace[i] = ancSeqSpace[ancPos]
		}
		idxOffset += cnt
	}
	return newSeqSpace
}

// MutateSeqSpace
func MutateSeqSpace(seqSpace *[][]int, mu float64, rateMatrix [][]float64) {
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

	// Returns two arrays, array[0] is always 0, array[1] is column coords, array[2] is number of hits
	hitsPerSeq := sampler.PoissonMutCoords(muPerSeq, popSize, 1)

	var permSites []int
	var seqIdx, char, newChar int
	for i, hits := range hitsPerSeq[2] {
		permSites = rand.Perm(numSites)
		seqIdx = hitsPerSeq[0][i]
		for _, siteIdx := range permSites[:hits] {
			char = (*seqSpace)[seqIdx][siteIdx]
			newChar = sampler.MultinomialWhere(1, rateMatrix[char], 1)[0]
			(*seqSpace)[seqIdx][siteIdx] = newChar
		}
	}
}
