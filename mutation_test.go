package mesim

import (
	"testing"
)

// Following tests different scenarios for the MutateChar function

// Test change from character 0 to 1 given that this is the only change
// possible based on the given transition rate matrix.
func TestMutateChar0To1(t *testing.T) {
	ancChar := 0
	// transition rate matrix guarantees change from 0 to 1
	rateMatrix := [][]float64{
		[]float64{0.0, 1.0, 0.0, 0.0},
		[]float64{0.0, 0.0, 0.0, 0.0},
		[]float64{0.0, 0.0, 0.0, 0.0},
		[]float64{0.0, 0.0, 0.0, 0.0},
	}

	var newChar int
	// EvolveChar for 10 rounds
	for i := 0; i < 10; i++ {
		newChar = ancChar
		MutateChar(&newChar, rateMatrix)
		if newChar != 1 {
			t.Errorf("EvolveChar(%d): expected 1, actual (%d)", ancChar, newChar)
		}
	}
}

// Test change from character 0 to 1 or 2 given that this is the only
// change possible based on the given transition rate matrix.
// There is an equal probability of changing to 1 or to 2.
func TestMutateChar0To1Or2(t *testing.T) {
	ancChar := 0
	// transition rate matrix guarantees change from 0 to 1 or 2
	// with equal probability
	rateMatrix := [][]float64{
		[]float64{0.0, 0.5, 0.5, 0.0},
		[]float64{0.0, 0.0, 0.0, 0.0},
		[]float64{0.0, 0.0, 0.0, 0.0},
		[]float64{0.0, 0.0, 0.0, 0.0},
	}

	var newChar int
	// EvolveChar for 10 rounds
	for i := 0; i < 10; i++ {
		newChar = ancChar
		MutateChar(&newChar, rateMatrix)
		if newChar != 1 && newChar != 2 {
			t.Errorf("EvolveChar(%d): expected 1 or 2, actual (%d)", ancChar, newChar)
		}
	}
}

// Test if character 0 will change given that the transition rate matrix
// does not allow it.
func TestMutateCharNoChange(t *testing.T) {
	ancChar := 0
	// transition rate matrix guarantees no transition
	rateMatrix := [][]float64{
		[]float64{1.0, 0.0, 0.0, 0.0},
		[]float64{0.0, 1.0, 0.0, 0.0},
		[]float64{0.0, 0.0, 1.0, 0.0},
		[]float64{0.0, 0.0, 0.0, 1.0},
	}

	var newChar int
	// EvolveChar for 10 rounds
	for i := 0; i < 10; i++ {
		newChar = ancChar
		MutateChar(&newChar, rateMatrix)
		if newChar != ancChar {
			t.Errorf("EvolveChar(%d): expected 0, actual (%d)", ancChar, newChar)
		}
	}
}

// Test if character 0 will change to any of the other characters given that
// the given transition rate matrix only allows transitions.

func TestMutateCharMustChange(t *testing.T) {
	ancChar := 0
	// transition rate matrix guarantees a transition
	// always happens
	rateMatrix := [][]float64{
		[]float64{0.0, 0.3, 0.3, 0.4},
		[]float64{0.4, 0.0, 0.3, 0.3},
		[]float64{0.3, 0.4, 0.0, 0.3},
		[]float64{0.3, 0.3, 0.4, 0.0},
	}

	var newChar int
	// EvolveChar for 10 rounds
	for i := 0; i < 10; i++ {
		newChar = ancChar
		MutateChar(&newChar, rateMatrix)
		if newChar == ancChar {
			t.Errorf("EvolveChar(%d): expected 1, 2, or 3, actual (%d)", ancChar, newChar)
		}
	}
}

// Test if character 0 after 10 rounds of successive mutations.
// The given transition rate matrix allows both remaining the same character
// or changing into a different character. However remaining the same has a
// very low probability compared to changing.
func TestMutateCharRandomChange(t *testing.T) {
	ancChar := 0
	// Small possibility that character will not change
	// after one step
	rateMatrix := [][]float64{
		[]float64{0.001, 0.333, 0.333, 0.333},
		[]float64{0.333, 0.001, 0.333, 0.333},
		[]float64{0.333, 0.333, 0.001, 0.333},
		[]float64{0.333, 0.333, 0.333, 0.001},
	}
	newChar := ancChar
	// Compound mutation for 10 rounds
	for i := 0; i < 10; i++ {
		MutateChar(&newChar, rateMatrix)
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
		MutateSeqExplicitly(&evolvedArray, rateMatrix)
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
		MutateSeqFast(&evolvedArray, 0.99, rateMatrix)
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
