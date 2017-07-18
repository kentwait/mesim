package mesim

import (
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
		t.Error("ancChar should not be equal to newChar")
		t.Error(ancChar, newChar)
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
		t.Error("ancArray should not be equal to evolvedArray")
		t.Error(ancArray, evolvedArray)
	}
}
