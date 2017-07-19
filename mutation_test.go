package mesim

import (
	"math/rand"
	"mesim/utils"
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
	rand.Seed(1)

	// MutateChar for 10 rounds
	for i := 0; i < 10; i++ {
		newChar = ancChar
		MutateChar(&newChar, rateMatrix)
		if newChar != 1 {
			t.Errorf("MutateChar(%d): expected 1, actual (%d)", ancChar, newChar)
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
	rand.Seed(1)

	// MutateChar for 10 rounds
	for i := 0; i < 10; i++ {
		newChar = ancChar
		MutateChar(&newChar, rateMatrix)
		if newChar != 1 && newChar != 2 {
			t.Errorf("MutateChar(%d): expected 1 or 2, actual (%d)", ancChar, newChar)
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
	rand.Seed(1)

	// MutateChar for 10 rounds
	for i := 0; i < 10; i++ {
		newChar = ancChar
		MutateChar(&newChar, rateMatrix)
		if newChar != ancChar {
			t.Errorf("MutateChar(%d): expected 0, actual (%d)", ancChar, newChar)
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
	rand.Seed(1)

	// MutateChar for 10 rounds
	for i := 0; i < 10; i++ {
		newChar = ancChar
		MutateChar(&newChar, rateMatrix)
		if newChar == ancChar {
			t.Errorf("MutateChar(%d): expected 1, 2, or 3, actual (%d)", ancChar, newChar)
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
	rand.Seed(1)

	// Compound mutation for 10 rounds
	for i := 0; i < 10; i++ {
		MutateChar(&newChar, rateMatrix)
	}
	if ancChar == newChar {
		t.Errorf("MutateChar(%d): expected 1, 2, or 3, actual (%d)", ancChar, newChar)
	}
}

// Following tests different scenarios for the MutateSeqExplicitly function

// Test if characters change to 1 given that this is the only change
// possible based on the given transition rate matrix. The initial
// sequence does not have 1's.
func TestMutateSeqExplicitlyTo1(t *testing.T) {
	ancSlice := []int{0, 2, 3, 0, 2, 3, 0, 2, 3}
	// transition rate matrix guarantees change to 1
	rateMatrix := [][]float64{
		[]float64{0.0, 1.0, 0.0, 0.0},
		[]float64{0.0, 0.0, 0.0, 0.0},
		[]float64{0.0, 1.0, 0.0, 0.0},
		[]float64{0.0, 1.0, 0.0, 0.0},
	}
	var evolvedSlice []int
	rand.Seed(1)

	// MutateChar for 10 rounds
	for i := 0; i < 10; i++ {
		// Deepcopy ancArray
		evolvedSlice = utils.DeepCopyInts(ancSlice)
		MutateSeqExplicitly(&evolvedSlice, rateMatrix)
		for i := range evolvedSlice {
			if evolvedSlice[i] != 1 {
				t.Errorf("MutateSeqExplicitly(seqArrayPtr, rateMatrix): expected 1, actual (%d)", evolvedSlice[i])
			}
		}
	}
}

// Test if characters change to 1 or 2 given that this is the only
// change possible based on the given transition rate matrix.
// There is an equal probability of changing to 1 or to 2. The initial
// sequence does not have 1's or 2's.
func TestMutateSeqExplicitlyTo1Or2(t *testing.T) {
	ancSlice := []int{0, 3, 0, 3, 0, 3, 0, 3, 0, 3}
	// transition rate matrix guarantees change to 1 or 2
	rateMatrix := [][]float64{
		[]float64{0.0, 0.5, 0.5, 0.0},
		[]float64{0.0, 0.0, 0.0, 0.0},
		[]float64{0.0, 0.0, 0.0, 0.0},
		[]float64{0.0, 0.5, 0.5, 0.0},
	}
	var evolvedSlice []int
	rand.Seed(1)

	// MutateChar for 10 rounds
	for i := 0; i < 10; i++ {
		// Deepcopy ancArray
		evolvedSlice = utils.DeepCopyInts(ancSlice)
		MutateSeqExplicitly(&evolvedSlice, rateMatrix)
		for i := range evolvedSlice {
			if evolvedSlice[i] != 1 && evolvedSlice[i] != 2 {
				t.Errorf("MutateSeqExplicitly(seqArrayPtr, rateMatrix): expected 1 or 2, actual (%d)", evolvedSlice[i])
			}
		}
	}
}

// Test if characters will change given that the transition rate matrix
// does not allow it. All possible characters are present in the original
// sequence.
func TestMutateSeqExplicitlyNoChange(t *testing.T) {
	ancSlice := []int{0, 1, 2, 3, 0, 1, 2, 3}
	// transition rate matrix guarantees change to 1
	rateMatrix := [][]float64{
		[]float64{1.0, 0.0, 0.0, 0.0},
		[]float64{0.0, 1.0, 0.0, 0.0},
		[]float64{0.0, 0.0, 1.0, 0.0},
		[]float64{0.0, 0.0, 0.0, 1.0},
	}
	var evolvedSlice []int
	rand.Seed(1)

	// MutateChar for 10 rounds
	for i := 0; i < 10; i++ {
		// Deepcopy ancArray
		evolvedSlice = utils.DeepCopyInts(ancSlice)
		MutateSeqExplicitly(&evolvedSlice, rateMatrix)
		for i := range evolvedSlice {
			if evolvedSlice[i] != ancSlice[i] {
				t.Errorf("MutateSeqExplicitly(seqArrayPtr, rateMatrix): expected %d, actual (%d)", ancSlice[i], evolvedSlice[i])
			}
		}
	}
}

// Test if characters will change to any of the other characters given that
// the given transition rate matrix only allows transitions. All possible
// characters are present in the original sequence.
func TestMutateSeqExplicitlyMustChange(t *testing.T) {
	ancSlice := []int{0, 1, 2, 3, 0, 1, 2, 3}
	// transition rate matrix guarantees a transition
	// always happens
	rateMatrix := [][]float64{
		[]float64{0.0, 0.3, 0.3, 0.4},
		[]float64{0.4, 0.0, 0.3, 0.3},
		[]float64{0.3, 0.4, 0.0, 0.3},
		[]float64{0.3, 0.3, 0.4, 0.0},
	}
	var evolvedSlice []int
	rand.Seed(1)

	// MutateChar for 10 rounds
	for i := 0; i < 10; i++ {
		// Deepcopy ancArray
		evolvedSlice = utils.DeepCopyInts(ancSlice)
		MutateSeqExplicitly(&evolvedSlice, rateMatrix)
		for i := range evolvedSlice {
			if evolvedSlice[i] == ancSlice[i] {
				t.Errorf("MutateSeqExplicitly(seqArrayPtr, rateMatrix): expected not equal to %d, actual (%d)", ancSlice[i], evolvedSlice[i])
			}
		}
	}
}

// Test if characters change after 10 rounds of successive mutations.
// The given transition rate matrix allows both remaining the same character
// or changing into a different character. However remaining the same has a
// very low probability compared to changing. All possible characters are
// present in the original sequence.
func TestMutateSeqExplicitlyRandom(t *testing.T) {
	ancSlice := []int{0, 1, 2, 3, 0, 1, 2, 3}
	rateMatrix := [][]float64{
		[]float64{0.001, 0.333, 0.333, 0.333},
		[]float64{0.333, 0.001, 0.333, 0.333},
		[]float64{0.333, 0.333, 0.001, 0.333},
		[]float64{0.333, 0.333, 0.333, 0.001},
	}
	// Deepcopy ancArray
	evolvedSlice := utils.DeepCopyInts(ancSlice)
	rand.Seed(1)

	// EvolveExplicit for 10 rounds
	for i := 0; i < 10; i++ {
		MutateSeqExplicitly(&evolvedSlice, rateMatrix)
	}
	sameSlices, _ := utils.CompareIntSlices(ancSlice, evolvedSlice)
	if sameSlices == true {
		t.Errorf("MutateSeqExplicitly(seqArrayPtr, rateMatrix): expected not equal to %v", ancSlice)
	}
}

// Following tests different scenarios for the MutateSeqFast function using
// the same tests for MutateSeqExplicitly

func TestMutateSeqFastTo1(t *testing.T) {
	ancSlice := []int{0, 2, 3, 0, 2, 3, 0, 2, 3}
	// transition rate matrix guarantees change to 1
	rateMatrix := [][]float64{
		[]float64{0.0, 1.0, 0.0, 0.0},
		[]float64{0.0, 0.0, 0.0, 0.0},
		[]float64{0.0, 1.0, 0.0, 0.0},
		[]float64{0.0, 1.0, 0.0, 0.0},
	}
	mutationRate := 1.0
	var evolvedSlice []int
	rand.Seed(1)

	// MutateChar for 10 rounds
	for i := 0; i < 10; i++ {
		// Deepcopy ancArray
		evolvedSlice = utils.DeepCopyInts(ancSlice)
		MutateSeqFast(&evolvedSlice, mutationRate, rateMatrix)
		for i := range evolvedSlice {
			if evolvedSlice[i] != 1 {
				t.Errorf("MutateSeqFast(seqArrayPtr, rateMatrix): expected 1, actual (%d)", evolvedSlice[i])
			}
		}
	}
}

func TestMutateSeqFastTo1Or2(t *testing.T) {
	ancSlice := []int{0, 3, 0, 3, 0, 3, 0, 3, 0, 3}
	// transition rate matrix guarantees change to 1 or 2
	rateMatrix := [][]float64{
		[]float64{0.0, 0.5, 0.5, 0.0},
		[]float64{0.0, 0.0, 0.0, 0.0},
		[]float64{0.0, 0.0, 0.0, 0.0},
		[]float64{0.0, 0.5, 0.5, 0.0},
	}
	mutationRate := 1.0
	var evolvedSlice []int
	rand.Seed(1)

	// MutateChar for 10 rounds
	for i := 0; i < 10; i++ {
		// Deepcopy ancArray
		evolvedSlice = utils.DeepCopyInts(ancSlice)
		MutateSeqFast(&evolvedSlice, mutationRate, rateMatrix)
		for i := range evolvedSlice {
			if evolvedSlice[i] != 1 && evolvedSlice[i] != 2 {
				t.Errorf("MutateSeqFast(seqArrayPtr, rateMatrix): expected 1 or 2, actual (%d)", evolvedSlice[i])
			}
		}
	}
}

func TestMutateSeqFastNoChange(t *testing.T) {
	ancSlice := []int{0, 1, 2, 3, 0, 1, 2, 3}
	// transition rate matrix guarantees change to 1
	rateMatrix := [][]float64{
		[]float64{1.0, 0.0, 0.0, 0.0},
		[]float64{0.0, 1.0, 0.0, 0.0},
		[]float64{0.0, 0.0, 1.0, 0.0},
		[]float64{0.0, 0.0, 0.0, 1.0},
	}
	mutationRate := 0.0
	var evolvedSlice []int
	rand.Seed(1)

	// MutateChar for 10 rounds
	for i := 0; i < 10; i++ {
		// Deepcopy ancArray
		evolvedSlice = utils.DeepCopyInts(ancSlice)
		MutateSeqFast(&evolvedSlice, mutationRate, rateMatrix)
		for i := range evolvedSlice {
			if evolvedSlice[i] != ancSlice[i] {
				t.Errorf("MutateSeqFast(seqArrayPtr, rateMatrix): expected %d, actual (%d)", ancSlice[i], evolvedSlice[i])
			}
		}
	}
}

func TestMutateSeqFastMustChange(t *testing.T) {
	ancSlice := []int{0, 1, 2, 3, 0, 1, 2, 3}
	// transition rate matrix guarantees a transition
	// always happens
	rateMatrix := [][]float64{
		[]float64{0.0, 0.3, 0.3, 0.4},
		[]float64{0.4, 0.0, 0.3, 0.3},
		[]float64{0.3, 0.4, 0.0, 0.3},
		[]float64{0.3, 0.3, 0.4, 0.0},
	}
	mutationRate := 1.0
	var evolvedSlice []int
	rand.Seed(1)

	// MutateChar for 10 rounds
	for i := 0; i < 10; i++ {
		// Deepcopy ancArray
		evolvedSlice = utils.DeepCopyInts(ancSlice)
		MutateSeqFast(&evolvedSlice, mutationRate, rateMatrix)
		for i := range evolvedSlice {
			if evolvedSlice[i] == ancSlice[i] {
				t.Errorf("MutateSeqFast(seqArrayPtr, rateMatrix): expected not equal to %d, actual (%d)", ancSlice[i], evolvedSlice[i])
			}
		}
	}
}

func TestMutateSeqFastRandom(t *testing.T) {
	ancSlice := []int{0, 1, 2, 3, 0, 1, 2, 3}
	rateMatrix := [][]float64{
		[]float64{0.0, 0.333, 0.333, 0.334},
		[]float64{0.334, 0.0, 0.333, 0.333},
		[]float64{0.333, 0.334, 0.0, 0.333},
		[]float64{0.333, 0.333, 0.334, 0.0},
	}
	mutationRate := 0.9
	// Deepcopy ancArray
	evolvedSlice := utils.DeepCopyInts(ancSlice)
	rand.Seed(1)

	// EvolveExplicit for 10 rounds
	for i := 0; i < 10; i++ {
		MutateSeqFast(&evolvedSlice, mutationRate, rateMatrix)
	}
	sameSlices, _ := utils.CompareIntSlices(ancSlice, evolvedSlice)
	if sameSlices == true {
		t.Errorf("MutateSeqFast(seqArrayPtr, rateMatrix): expected not equal to %v", ancSlice)
	}
}

// Following tests different scenarios for the MutateSeqSpace function using
// the same tests for MutateSeqExplicitly

/*
func TestMutateSeqSpaceTo1(t *testing.T) {
	ancSeqSpace := [][]int{
		[]int{0, 2, 3, 0, 2, 3, 0, 2, 3},
		[]int{0, 2, 3, 0, 2, 3, 0, 2, 3},
		[]int{0, 2, 3, 0, 2, 3, 0, 2, 3},
		[]int{0, 2, 3, 0, 2, 3, 0, 2, 3},
		[]int{0, 2, 3, 0, 2, 3, 0, 2, 3},
	}
	// transition rate matrix guarantees change to 1
	rateMatrix := [][]float64{
		[]float64{0.0, 1.0, 0.0, 0.0},
		[]float64{0.0, 0.0, 0.0, 0.0},
		[]float64{0.0, 1.0, 0.0, 0.0},
		[]float64{0.0, 1.0, 0.0, 0.0},
	}
	mutationRate := 1.0 / float64(9)
	var evolvedSeqSpace [][]int
	rand.Seed(1)

	// MutateChar for 10 rounds
	for i := 0; i < 10; i++ {
		// Deepcopy ancArray
		evolvedSeqSpace = utils.DeepCopyInts2d(ancSeqSpace)
		fmt.Println(evolvedSeqSpace)
		MutateSeqSpace(&evolvedSeqSpace, mutationRate, rateMatrix)

		sameMatrices, _ := utils.CompareIntMatrix(ancSeqSpace, evolvedSeqSpace)
		if sameMatrices == true {
			t.Errorf("MutateSeqSpace(seqArrayPtr, rateMatrix): expected all values to be 1, actual %v", evolvedSeqSpace)
		}
	}
}


func TestMutateSeqSpaceTo1Or2(t *testing.T) {
	ancSeqSpace := []int{0, 3, 0, 3, 0, 3, 0, 3, 0, 3}
	// transition rate matrix guarantees change to 1 or 2
	rateMatrix := [][]float64{
		[]float64{0.0, 0.5, 0.5, 0.0},
		[]float64{0.0, 0.0, 0.0, 0.0},
		[]float64{0.0, 0.0, 0.0, 0.0},
		[]float64{0.0, 0.5, 0.5, 0.0},
	}
	mutationRate := 1.0
	var evolvedSlice []int
	rand.Seed(1)

	// MutateChar for 10 rounds
	for i := 0; i < 10; i++ {
		// Deepcopy ancArray
		evolvedSlice = utils.DeepCopyInts(ancSeqSpace)
		MutateSeqSpace(&evolvedSlice, mutationRate, rateMatrix)
		for i := range evolvedSlice {
			if evolvedSlice[i] != 1 && evolvedSlice[i] != 2 {
				t.Errorf("MutateSeqSpace(seqArrayPtr, rateMatrix): expected 1 or 2, actual (%d)", evolvedSlice[i])
			}
		}
	}
}

func TestMutateSeqSpaceNoChange(t *testing.T) {
	ancSeqSpace := []int{0, 1, 2, 3, 0, 1, 2, 3}
	// transition rate matrix guarantees change to 1
	rateMatrix := [][]float64{
		[]float64{1.0, 0.0, 0.0, 0.0},
		[]float64{0.0, 1.0, 0.0, 0.0},
		[]float64{0.0, 0.0, 1.0, 0.0},
		[]float64{0.0, 0.0, 0.0, 1.0},
	}
	mutationRate := 0.0
	var evolvedSlice []int
	rand.Seed(1)

	// MutateChar for 10 rounds
	for i := 0; i < 10; i++ {
		// Deepcopy ancArray
		evolvedSlice = utils.DeepCopyInts(ancSeqSpace)
		MutateSeqSpace(&evolvedSlice, mutationRate, rateMatrix)
		for i := range evolvedSlice {
			if evolvedSlice[i] != ancSeqSpace[i] {
				t.Errorf("MutateSeqSpace(seqArrayPtr, rateMatrix): expected %d, actual (%d)", ancSeqSpace[i], evolvedSlice[i])
			}
		}
	}
}

func TestMutateSeqSpaceMustChange(t *testing.T) {
	ancSeqSpace := []int{0, 1, 2, 3, 0, 1, 2, 3}
	// transition rate matrix guarantees a transition
	// always happens
	rateMatrix := [][]float64{
		[]float64{0.0, 0.3, 0.3, 0.4},
		[]float64{0.4, 0.0, 0.3, 0.3},
		[]float64{0.3, 0.4, 0.0, 0.3},
		[]float64{0.3, 0.3, 0.4, 0.0},
	}
	mutationRate := 1.0
	var evolvedSlice []int
	rand.Seed(1)

	// MutateChar for 10 rounds
	for i := 0; i < 10; i++ {
		// Deepcopy ancArray
		evolvedSlice = utils.DeepCopyInts(ancSeqSpace)
		MutateSeqSpace(&evolvedSlice, mutationRate, rateMatrix)
		for i := range evolvedSlice {
			if evolvedSlice[i] == ancSeqSpace[i] {
				t.Errorf("MutateSeqSpace(seqArrayPtr, rateMatrix): expected not equal to %d, actual (%d)", ancSeqSpace[i], evolvedSlice[i])
			}
		}
	}
}

func TestMutateSeqSpaceRandom(t *testing.T) {
	ancSeqSpace := []int{0, 1, 2, 3, 0, 1, 2, 3}
	rateMatrix := [][]float64{
		[]float64{0.0, 0.333, 0.333, 0.334},
		[]float64{0.334, 0.0, 0.333, 0.333},
		[]float64{0.333, 0.334, 0.0, 0.333},
		[]float64{0.333, 0.333, 0.334, 0.0},
	}
	mutationRate := 0.9
	// Deepcopy ancArray
	evolvedSlice := utils.DeepCopyInts(ancSeqSpace)
	rand.Seed(1)

	// EvolveExplicit for 10 rounds
	for i := 0; i < 10; i++ {
		MutateSeqSpace(&evolvedSlice, mutationRate, rateMatrix)
	}
	sameSlices, _ := utils.CompareIntSlices(ancSeqSpace, evolvedSlice)
	if sameSlices == true {
		t.Errorf("MutateSeqSpace(seqArrayPtr, rateMatrix): expected not equal to %v", ancSeqSpace)
	}
}
*/
