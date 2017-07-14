package utils

func DivMod(numerator, denominator int64) (quotient, remainder int64) {
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
