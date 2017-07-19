package utils

import "testing"

func TestCompareIntSlicesSame(t *testing.T) {
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	same, _ := CompareIntSlices(s1, s1)
	if same == false {
		t.Errorf("CompareIntSlices(%v, %v): expected true, actual %v", s1, s1, same)
	}
}

func TestCompareIntSlicesDiff(t *testing.T) {
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	same, _ := CompareIntSlices(s1, s2)
	if same == true {
		t.Errorf("CompareIntSlices(%v, %v): expected false, actual %v", s1, s2, same)
	}
}

func TestCompareIntMatrixSame(t *testing.T) {
	s1 := [][]int{
		[]int{0, 1, 2, 3},
		[]int{0, 1, 2, 3},
		[]int{0, 1, 2, 3},
	}

	same, _ := CompareIntMatrix(s1, s1)
	if same == false {
		t.Errorf("CompareIntMatrix(%v, %v): expected true, actual %v", s1, s1, same)
	}
}

func TestCompareIntMatrixDiff(t *testing.T) {
	s1 := [][]int{
		[]int{0, 1, 2, 3},
		[]int{0, 1, 2, 3},
		[]int{0, 1, 2, 3},
	}
	s2 := [][]int{
		[]int{3, 2, 1, 0},
		[]int{3, 2, 1, 0},
		[]int{3, 2, 1, 0},
	}

	same, _ := CompareIntMatrix(s1, s2)
	if same == true {
		t.Errorf("CompareIntMatrix(%v, %v): expected true, actual %v", s1, s2, same)
	}
}
