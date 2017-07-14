package sampler

import (
	"math/rand"
)

// Multinomial draws a sample from a multinomial distribution.
func Multinomial(n int64, p []float64) []int64 {
	// If n * len(p) > 1000, uses concurrency
	if n*int64(len(p)) > int64(1000) {
		workers := 0
		resultChan := make(chan []int64)

		for n > 1000 {
			go func() {
				resultChan <- multinomial(1000, p)
			}()
			n -= 1000
			workers++
		}
		go func() {
			resultChan <- multinomial(n, p)
		}()
		workers++

		result := make([]int64, len(p))
		for i := 0; i < workers; i++ {
			tmp := <-resultChan
			for j := 0; j < len(result); j++ {
				result[j] += tmp[j]
			}
		}
		return result
	} else {
		return multinomial(n, p)
	}
}

// MultinomialWhere returns the coordinates equal to the given value
func MultinomialWhere(n int64, p []float64, cnt int64) (result []int64) {
	for i, hit := range Multinomial(n, p) {
		if hit == cnt {
			result = append(result, int64(i))
		}
	}
	return
}

// multinomial is the base function of Multinomial.
func multinomial(n int64, p []float64) []int64 {
	result := make([]int64, len(p))
	cumP := make([]float64, len(p))

	// Create a cummulative distribution of p
	cumP[0] = p[0]
	for i := 1; i < len(p); i++ {
		cumP[i] = cumP[i-1] + p[i]
	}

	for i := int64(0); i < n; i++ {
		// Generate pseudorandom number
		x := rand.Float64()

		// for j; e := range cumP {
		for j := 0; j < len(cumP); j++ {
			if x < cumP[j] {
				result[j]++
				break
			} else if x == 1.0 {
				result[len(p)]++
				break
			}
		}
	}
	return result
}
