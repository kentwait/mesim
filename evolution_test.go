package mesim

import (
	"fmt"
	"testing"
)

func TestEvolveChar(t *testing.T) {
	ancChar := 0
	rateMatrix := [][]float64{
		[]float64{0.001, 0.333, 0.333, 0.333},
		[]float64{0.333, 0.001, 0.333, 0.333},
		[]float64{0.333, 0.333, 0.001, 0.333},
		[]float64{0.333, 0.333, 0.333, 0.001},
	}
	newChar := ancChar
	// EvolveChar for 10 rounds
	for i := 0; i < 10; i++ {
		newChar = EvolveChar(newChar, rateMatrix)
	}
	if ancChar == newChar {
		t.Errorf("EvolveChar(%d): expected 1, 2, or 3, actual (%d)", ancChar, newChar)
	}
}
func TestEvolveExplicit(t *testing.T) {
	ancArray := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0} // 10 0's
	rateMatrix := [][]float64{
		[]float64{0.001, 0.333, 0.333, 0.333},
		[]float64{0.333, 0.001, 0.333, 0.333},
		[]float64{0.333, 0.333, 0.001, 0.333},
		[]float64{0.333, 0.333, 0.333, 0.001},
	}
	// Deepcopy ancArray
	evolvedArray := make([]int, len(ancArray))
	for i := range evolvedArray {
		evolvedArray[i] = ancArray[i]
	}

	// EvolveExplicit for 10 rounds
	for i := 0; i < 10; i++ {
		EvolveExplicit(&evolvedArray, rateMatrix)
	}
	diffCnt := 0
	for i := range ancArray {
		if ancArray[i] != evolvedArray[i] {
			diffCnt++
		}
	}
	if diffCnt == 0 {
		t.Error("EvolveExplicit(ancArray, rateMatrix): ancArray should not be equal to result evolvedArray")
		t.Error(ancArray, evolvedArray)
	}
}

func TestEvolveFast(t *testing.T) {
	ancArray := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0} // 10 0's
	rateMatrix := [][]float64{
		[]float64{0.00, 0.34, 0.33, 0.33},
		[]float64{0.33, 0.00, 0.34, 0.33},
		[]float64{0.33, 0.33, 0.00, 0.34},
		[]float64{0.34, 0.33, 0.33, 0.00},
	}
	// Deepcopy ancArray
	evolvedArray := make([]int, len(ancArray))
	for i := range evolvedArray {
		evolvedArray[i] = ancArray[i]
	}

	// EvolveExplicit for 10 rounds
	for i := 0; i < 10; i++ {
		EvolveFast(&evolvedArray, 0.99, rateMatrix)
	}
	diffCnt := 0
	for i := range ancArray {
		if ancArray[i] != evolvedArray[i] {
			diffCnt++
		}
	}
	if diffCnt == 0 {
		t.Error("TestEvolveFast(ancArray, 0.99, rateMatrix),  ancArray should not be equal to result evolvedArray")
		t.Error(ancArray, evolvedArray)
	}
}

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
	fitnessFunction := func(seq []int) (multSum float64) {
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
	fitnessFunction := func(seq []int) (multSum float64) {
		multSum = float64(1)
		for i, char := range seq {
			multSum *= float64(fitnessMatrix[i][char])
		}
		return
	}

	var tempSeqSpace [][]int
	for i := 0; i < 100; i++ {
		fmt.Println(tempSeqSpace)
		tempSeqSpace = ReplicateSelect(newSeqSpace, nextPopSize, fitnessMatrix, fitnessFunction)
		newSeqSpace = tempSeqSpace
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

func TestMutateSeqSpace(t *testing.T) {
	ancSeqSpace := [][]int{
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	rateMatrix := [][]float64{
		[]float64{0.00, 0.34, 0.33, 0.33},
		[]float64{0.33, 0.00, 0.34, 0.33},
		[]float64{0.33, 0.33, 0.00, 0.34},
		[]float64{0.34, 0.33, 0.33, 0.00},
	}
	mu := 0.1

	evolvedSeqSpace := make([][]int, len(ancSeqSpace))
	for i := range ancSeqSpace {
		evolvedSeqSpace[i] = make([]int, len(ancSeqSpace[i]))
		for j := range ancSeqSpace[i] {
			evolvedSeqSpace[i][j] = ancSeqSpace[i][j]
		}
	}

	for i := 0; i < 10; i++ {
		MutateSeqSpace(&evolvedSeqSpace, mu, rateMatrix)
	}

	diffCnt := 0
	for i := range ancSeqSpace {
		for j := range ancSeqSpace[i] {
			if ancSeqSpace[i][j] != evolvedSeqSpace[i][j] {
				diffCnt++
			}
		}
	}
	if diffCnt == 0 {
		t.Errorf("TestMutateSeqSpace(ancSeqSpace, %f, rateMatrix): ancSeqSpace should not be equal to result evolvedSeqSpace", mu)
		t.Error(ancSeqSpace, evolvedSeqSpace)
	}
}
