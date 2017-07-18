package utils

func DivMod(numerator, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}

func Sum(input ...float64) (sum float64) {
	sum = 0
	for _, i := range input {
		sum += i
	}
	return
}

func ColSum(matrix [][]float64) []float64 {
	if len(matrix) < 1 {
		panic("Matrix must have one or more rows")
	} else {
		if len(matrix[0]) < 1 {
			panic("Matrix must have one or more columns")
		}
	}
	colSum := make([]float64, len(matrix[0]))
	for m := 0; m < len(matrix); m++ {
		for n := 0; n < len(matrix[0]); n++ {
			colSum[n] += matrix[m][n]
		}
	}
	return colSum
}

func ColMean(matrix [][]float64) []float64 {
	if len(matrix) < 1 {
		panic("Matrix must have one or more rows")
	} else {
		if len(matrix[0]) < 1 {
			panic("Matrix must have one or more columns")
		}
	}
	colMean := make([]float64, len(matrix[0]))
	sums := ColSum(matrix)
	for i, sum := range sums {
		colMean[i] = sum / float64(len(matrix))
	}
	return colMean
}
