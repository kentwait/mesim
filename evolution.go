package mesim

import (
	"math"
	"math/rand"
	"mesim/sampler"
	"mesim/utils"
	"sort"
)

// FitnessFunc
type FitnessFunc func([]int, [][]float64) float64

// SeqSpaceToFitSpace
func SeqSpaceToFitSpace(seqSpace [][]int, fitnessMatrix [][]float64, totalFitnessFunc FitnessFunc, normalized bool) []float64 {
	fitnessSpace := seqSpaceToFitSpace(seqSpace, fitnessMatrix, totalFitnessFunc)
	if normalized == true {
		fitnessDenominator := utils.Sum(fitnessSpace...)
		for i := range fitnessSpace {
			fitnessSpace[i] = fitnessSpace[i] / fitnessDenominator
		}
	}
	return fitnessSpace
}

// SeqSpaceToLogFitSpace
func SeqSpaceToLogFitSpace(seqSpace [][]int, fitnessMatrix [][]float64, totalFitnessFunc FitnessFunc, normalized bool) []float64 {
	fitnessSpace := seqSpaceToFitSpace(seqSpace, fitnessMatrix, totalFitnessFunc)
	if normalized == true {
		fitnessDenominator := math.Log(utils.Sum(fitnessSpace...))
		for i := range fitnessSpace {
			fitnessSpace[i] = fitnessSpace[i] - fitnessDenominator
		}
	}
	return fitnessSpace
}

func seqSpaceToFitSpace(seqSpace [][]int, fitnessMatrix [][]float64, totalFitnessFunc FitnessFunc) []float64 {
	if len(seqSpace) == 0 {
		panic("Length of seqSpace must be greater than zero")
	} else {
		if len(seqSpace[0]) == 0 {
			panic("Length of rows in seqSpace must be greater than zero")
		}
	}
	if len(fitnessMatrix) == 0 {
		panic("Length of fitnessMatrix must be greater than zero")
	}

	fitnessSpace := make([]float64, len(seqSpace))

	for i, seq := range seqSpace {
		fitnessSpace[i] = totalFitnessFunc(seq, fitnessMatrix)
	}
	return fitnessSpace
}

// ReplicateSelect
func ReplicateSelect(ancSeqSpace [][]int, nextPopSize int, fitnessMatrix [][]float64, totalFitnessFunc FitnessFunc) [][]int {
	normedFitSpace := SeqSpaceToFitSpace(ancSeqSpace, fitnessMatrix, totalFitnessFunc, true)
	ancSeqSpaceCnts := sampler.MultinomialSample(nextPopSize, normedFitSpace)

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

// RecombineSeqSpace
func RecombineSeqSpace(seqSpace *[][]int, r float64) {
	// Randomly pick (by permutation) sequence pairs
	popSize := len(*seqSpace)
	numSites := len((*seqSpace)[0]) - 1 // One less site because we are counting breakpoints
	permSampleIndexes := rand.Perm(popSize)

	// For each sequence pair, determine number of recombination events e
	var numEvents, seqID1, seqID2, startPos int
	var s1Ptr, s2Ptr *[]int
	var newS1, newS2, permSites []int
	orientation := true
	for i := 0; i < popSize-1; i += 2 {
		// Processing each pair could be made into a goroutine
		numEvents = sampler.BinomialSample(numSites, r)

		// For each sequence pair, randomly pick (by permutation) breakpoints
		if numEvents > 0 {
			seqID1 = permSampleIndexes[i]
			seqID2 = permSampleIndexes[i+1]
			startPos = 0
			newS1 = []int{}
			newS2 = []int{}
			orientation = true

			s1Ptr, s2Ptr = &(*seqSpace)[seqID1], &(*seqSpace)[seqID2]

			permSites = rand.Perm(numSites)[:numEvents] // +1 so that lowest breakpoint is [0:1] and highest is [-1:]
			sort.Ints(permSites)

			for _, pos := range permSites {
				pos++ // Lowest pos == 1, highest pos == len - 1
				if orientation == true {
					newS1 = append(newS1, (*s1Ptr)[startPos:pos]...)
					newS2 = append(newS2, (*s2Ptr)[startPos:pos]...)
					orientation = false
				} else {
					newS1 = append(newS1, (*s2Ptr)[startPos:pos]...)
					newS2 = append(newS2, (*s1Ptr)[startPos:pos]...)
					orientation = true
				}
				startPos = pos
			}
			if orientation == true {
				newS1 = append(newS1, (*s1Ptr)[startPos:]...)
				newS2 = append(newS2, (*s2Ptr)[startPos:]...)
				orientation = false
			} else {
				newS1 = append(newS1, (*s2Ptr)[startPos:]...)
				newS2 = append(newS2, (*s1Ptr)[startPos:]...)
				orientation = true
			}
			(*seqSpace)[seqID1] = newS1
			(*seqSpace)[seqID2] = newS2
		}
	}
}

// EvolveSeqSpaceConstPop
func EvolveSeqSpaceConstPop(seqSpace *[][]int, mutationRate float64, recombinationRate float64, charTransitionMatrix [][]float64, fitnessMatrix [][]float64, fitnessFunc FitnessFunc) {
	popSize := len(*seqSpace)
	*seqSpace = ReplicateSelect(*seqSpace, popSize, fitnessMatrix, fitnessFunc)
	MutateSeqSpace(seqSpace, mutationRate, charTransitionMatrix)
	RecombineSeqSpace(seqSpace, recombinationRate)
}
