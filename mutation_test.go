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
		newChar = MutateChar(newChar, rateMatrix)
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