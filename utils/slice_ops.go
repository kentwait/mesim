package utils

func DeepCopyInts(s []int) []int {
	newCopy := make([]int, len(s))
	for i := range newCopy {
		newCopy[i] = s[i]
	}
	return newCopy
}

func DeepCopyInts2d(s [][]int) [][]int {
	newCopy := make([][]int, len(s))
	for i := range newCopy {
		newCopy[i] = DeepCopyInts(s[i])
	}
	return newCopy
}

func CompareIntSlices(slice1, slice2 []int) (same bool, diffCoords []int) {
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			diffCoords = append(diffCoords, i)
			same = false
		}
	}
	return
}

func CompareIntMatrix(matrix1, matrix2 [][]int) (same bool, diffCoords [][]int) {
	var diffCoordsX, diffCoordsY []int

	for i, row := range matrix1 {
		for j := range row {
			if matrix1[i][j] != matrix2[i][j] {
				diffCoordsX = append(diffCoordsX, i)
				diffCoordsY = append(diffCoordsY, j)
				same = false
			}
		}
	}
	diffCoords = [][]int{diffCoordsX, diffCoordsY}
	return
}
