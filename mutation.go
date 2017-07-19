package mesim

import (
	"math/rand"
	"mesim/sampler"
)

// MutateChar
func MutateChar(char *int, rateMatrix [][]float64) {
	*char = sampler.MultinomialWhere(1, rateMatrix[*char], 1)[0]
}

// MutateSeqExplicitly
func MutateSeqExplicitly(ancArray *[]int, rateMatrix [][]float64) {
	for i, char := range *ancArray {
		(*ancArray)[i] = sampler.MultinomialWhere(1, rateMatrix[char], 1)[0]
	}
}

// MutateSeqFast
func MutateSeqFast(ancArray *[]int, mu float64, zeroedRateMatrix [][]float64) {
	mutCoords := sampler.PoissonMutCoords(mu, len(*ancArray), 1)
	var tmpArray []int
	if len(mutCoords) > 0 {
		for _, yPos := range mutCoords[1] {
			tmpArray = append(tmpArray, (*ancArray)[yPos])
		}
		MutateSeqExplicitly(&tmpArray, zeroedRateMatrix)
		for i, xPos := range mutCoords[0] {
			(*ancArray)[xPos] = tmpArray[i]
		}
	}

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
