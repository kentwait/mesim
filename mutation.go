package mesim

import (
	"math/rand"
	"mesim/sampler"
)

// MutateChar
func MutateChar(charPtr *int, rateMatrix [][]float64) {
	*charPtr = sampler.MultinomialWhere(1, rateMatrix[*charPtr], 1)[0]
}

// MutateSeqExplicitly
func MutateSeqExplicitly(ancArrayPtr *[]int, rateMatrix [][]float64) {
	for i, char := range *ancArrayPtr {
		(*ancArrayPtr)[i] = sampler.MultinomialWhere(1, rateMatrix[char], 1)[0]
	}
}

// MutateSeqFast
func MutateSeqFast(ancArrayPtr *[]int, mu float64, zeroedRateMatrix [][]float64) {
	mutCoords := sampler.PoissonMutCoords(mu, len(*ancArrayPtr), 1)
	var tmpArray []int
	if len(mutCoords) > 0 {
		for _, yPos := range mutCoords[1] {
			tmpArray = append(tmpArray, (*ancArrayPtr)[yPos])
		}
		MutateSeqExplicitly(&tmpArray, zeroedRateMatrix)
		for i, xPos := range mutCoords[0] {
			(*ancArrayPtr)[xPos] = tmpArray[i]
		}
	}

}

// MutateSeqSpace
func MutateSeqSpace(seqSpacePtr *[][]int, mu float64, rateMatrix [][]float64) {
	if len(*seqSpacePtr) == 0 {
		panic("Length of seqSpace must be greater than zero")
	} else {
		if len((*seqSpacePtr)[0]) == 0 {
			panic("Length of rows in seqSpace must be greater than zero")
		}
	}
	popSize := len(*seqSpacePtr)
	numSites := len((*seqSpacePtr)[0])
	muPerSeq := mu * float64(numSites)

	// Returns two arrays, array[0] is always 0, array[1] is column coords, array[2] is number of hits
	hitsPerSeq := sampler.PoissonMutCoords(muPerSeq, popSize, 1)

	var permSites []int
	var seqIdx, char, newChar int
	for i, hits := range hitsPerSeq[2] {
		permSites = rand.Perm(numSites)
		seqIdx = hitsPerSeq[0][i]
		for _, siteIdx := range permSites[:hits] {
			char = (*seqSpacePtr)[seqIdx][siteIdx]
			newChar = sampler.MultinomialWhere(1, rateMatrix[char], 1)[0]
			(*seqSpacePtr)[seqIdx][siteIdx] = newChar
		}
	}
}
