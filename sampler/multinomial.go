package sampler

import (
	"math/rand"
)

// Multinomial draws a sample from a multinomial distribution.
func Multinomial(n int, p []float64) []int {
	// If n * len(p) > 1000, uses concurrency
	if n*len(p) > 1000 {
		workers := 0
		resultChan := make(chan []int)

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

		result := make([]int, len(p))
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
func MultinomialWhere(n int, p []float64, cnt int) (result []int) {
	for i, hit := range Multinomial(n, p) {
		if hit == cnt {
			result = append(result, i)
		}
	}
	return
}

// multinomial is the base function of Multinomial.
func multinomial(n int, p []float64) []int {
	result := make([]int, len(p))
	cumP := make([]float64, len(p))

	// Create a cummulative distribution of p
	cumP[0] = p[0]
	for i := 1; i < len(p); i++ {
		cumP[i] = cumP[i-1] + p[i]
	}

	for i := 0; i < n; i++ {
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

func multinomialLog(n int, logP []float64) []int {
	result := make([]int, len(logP))

	for i := 0; i < n; i++ {

	}
	return result
}
