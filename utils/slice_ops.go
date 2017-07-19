package utils

func DeepCopyInts(s []int) []int {
	newCopy := make([]int, len(s))
	for i := range newCopy {
		newCopy[i] = s[i]
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
