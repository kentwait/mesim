package sampler

import (
	"math"
	"math/rand"
)

// Multinomial draws a sample from a multinomial distribution.
func Multinomial(n int, p []float64) (result []int) {
	result = generalMultinomial(n, p, false)
	return result
}

// MultinomialLog draws a sample from a multinomial distribution.
func MultinomialLog(n int, p []float64) (result []int) {
	result = generalMultinomial(n, p, true)
	return result
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

func generalMultinomial(n int, p []float64, isLogP bool) []int {
	trials := 10000

	// If n * len(p) > 1000, uses concurrency
	if n*len(p) > trials {
		workers := 0
		resultChan := make(chan []int)

		for n > trials {
			go func() {
				if isLogP == true {
					resultChan <- multinomialLog(trials, p)
				} else {
					resultChan <- multinomial(trials, p)
				}
			}()
			n -= trials
			workers++
		}
		go func() {
			if isLogP == true {
				resultChan <- multinomialLog(n, p)
			} else {
				resultChan <- multinomial(n, p)
			}
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
	}
	if isLogP == true {
		return multinomialLog(n, p)
	}
	return multinomial(n, p)

}

// multinomial is the base function of Multinomial.
func multinomial(n int, p []float64) []int {
	result := make([]int, len(p))
	cumP := make([]float64, len(p))
	lastIdx := len(p) - 1

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
			} else if x > cumP[lastIdx] && x <= 1.0 {
				result[lastIdx]++
				break
			}
		}
	}
	return result
}

func multinomialLog(n int, logP []float64) []int {
	result := make([]int, len(logP))
	cumP := make([]float64, len(logP))
	lastIdx := len(logP) - 1

	// Create a cummulative distribution of p
	cumP[0] = math.Exp(logP[0])
	for i := 1; i < len(logP); i++ {
		cumP[i] = cumP[i-1] + math.Exp(logP[i])
	}

	// fmt.Println(cumP)
	for i := 0; i < n; i++ {
		// Generate pseudorandom number
		x := rand.Float64()

		// for j; e := range cumP {
		for j := 0; j < len(cumP); j++ {
			if x < cumP[j] {
				result[j]++
				break
			} else if x > cumP[lastIdx] && x <= 1.0 {
				result[lastIdx]++
				break
			}
		}
	}
	return result
}
