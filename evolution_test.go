package mesim

import (
	"fmt"
	"testing"
)

func TestSeqSpaceToFitSpace(t *testing.T) {
	seqSpace := [][]int{
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	fitnessMatrix := [][]float64{
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
	}
	fitnessFunction := func(seq []int, fitnessMatrix [][]float64) (multSum float64) {
		multSum = float64(1)
		for i, char := range seq {
			multSum *= float64(fitnessMatrix[i][char])
		}
		return
	}
	expectedFitSpace := []float64{
		1 / float64(5),
		1 / float64(5),
		1 / float64(5),
		1 / float64(5),
		1 / float64(5),
	}

	actualFitSpace := SeqSpaceToFitSpace(seqSpace, fitnessMatrix, fitnessFunction, true)

	diffCnt := 0
	for i := range expectedFitSpace {
		if expectedFitSpace[i] != actualFitSpace[i] {
			diffCnt++
		}
	}
	if diffCnt > 0 {
		t.Errorf("SeqSpaceToFitSpace(seqSpace, fitnessMatrix, fitnessFunction, true): expected %v, actual %v", expectedFitSpace, actualFitSpace)
	}

}

func TestReplicateSelect(t *testing.T) {
	seqSpace := [][]int{
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	newSeqSpace := [][]int{
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	nextPopSize := len(newSeqSpace) // 6
	fitnessMatrix := [][]float64{
		[]float64{1.0, 1.5, 1.0, 1.0}, // Fitness of 1 at pos 0 is 2x baseline
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
	}
	fitnessFunction := func(seq []int, fitnessMatrix [][]float64) (multSum float64) {
		multSum = float64(1)
		for i, char := range seq {
			multSum *= float64(fitnessMatrix[i][char])
		}
		return
	}

	for i := 0; i < 100; i++ {
		newSeqSpace = ReplicateSelect(newSeqSpace, nextPopSize, fitnessMatrix, fitnessFunction)
	}

	diffCnt := 0
	for i := range newSeqSpace {
		for j := range newSeqSpace[i] {
			if seqSpace[i][j] != newSeqSpace[i][j] {
				diffCnt++
			}
		}
	}
	if diffCnt == 0 {
		t.Errorf("TestReplicateSelect(newSeqSpace, %d, fitnessMatrix, fitnessFunction): original seqSpace should not be equal to result newSeqSpace", nextPopSize)
		t.Error(seqSpace, newSeqSpace)
	}

}

func TestRecombineSeqSpace(t *testing.T) {
	ancSeqSpace := [][]int{
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		[]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		[]int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		[]int{4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
	}
	r := 0.1
	evolvedSeqSpace := make([][]int, len(ancSeqSpace))
	for i := range ancSeqSpace {
		evolvedSeqSpace[i] = make([]int, len(ancSeqSpace[i]))
		for j := range ancSeqSpace[i] {
			evolvedSeqSpace[i][j] = ancSeqSpace[i][j]
		}
	}
	// fmt.Println("anc", ancSeqSpace)
	RecombineSeqSpace(&evolvedSeqSpace, r)
	// fmt.Println("evolved", evolvedSeqSpace)

	// TODO
	// Break up RecombineSeqSpace to test these components individually
	// Test whether evolvedSeqSpace has same shape as ancSeqSpace
	// Test if recombination occurred
	// Test if given a set of breakpoints, function uses
}

func TestEvolveSeqSpaceConstPop(t *testing.T) {
	ancSeqSpace := [][]int{
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		[]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		[]int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
	}
	mutationRate := 0.1
	recombinationRate := 0.1
	rateMatrix := [][]float64{
		[]float64{0.00, 0.34, 0.33, 0.33},
		[]float64{0.33, 0.00, 0.34, 0.33},
		[]float64{0.33, 0.33, 0.00, 0.34},
		[]float64{0.34, 0.33, 0.33, 0.00},
	}
	fitnessMatrix := [][]float64{
		[]float64{1.0, 1.0, 1.0, 1.0}, // Fitness of 1 at pos 0 is 2x baseline
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
		[]float64{1.0, 1.0, 1.0, 1.0},
	}
	fitnessFunc := func(seq []int, fitnessMatrix [][]float64) (multSum float64) {
		multSum = float64(1)
		for i, char := range seq {
			multSum *= float64(fitnessMatrix[i][char])
		}
		return
	}
	evolvedSeqSpace := make([][]int, len(ancSeqSpace))
	for i := range ancSeqSpace {
		evolvedSeqSpace[i] = make([]int, len(ancSeqSpace[i]))
		for j := range ancSeqSpace[i] {
			evolvedSeqSpace[i][j] = ancSeqSpace[i][j]
		}
	}

	fmt.Println(evolvedSeqSpace)
	EvolveSeqSpaceConstPop(&evolvedSeqSpace, mutationRate, recombinationRate, rateMatrix, fitnessMatrix, fitnessFunc)
	fmt.Println(evolvedSeqSpace)

}
