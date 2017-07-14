package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//PoissonMutArray
func PoissonMutArray(mu float64, nSites, popSize int64) [][]int64 {
	var result [][]int64
	tmp := make([]int64, nSites*popSize)
	n := nSites * popSize
	blockSize := int64(1e6) //TODO : Optimize block size

	if n > blockSize {
		workers := 0
		resultChan := make(chan []int64)

		for n > blockSize {
			go func() {
				resultChan <- poissonMutArray(mu, 1e5)
			}()
			n -= blockSize
			workers++
		}
		go func() {
			resultChan <- poissonMutArray(mu, n)
		}()
		workers++

		for i := 0; i < workers; i++ {
			tmp = append(tmp, <-resultChan...)
		}
	} else {
		tmp = poissonMutArray(mu, nSites*popSize)
	}

	for i := int64(0); i < n; i += nSites {
		result = append(result, tmp[i:i+nSites])
	}
	return result
}

func poissonMutArray(mu float64, n int64) []int64 {
	result := make([]int64, n)

	for i := int64(0); i < n; i++ {
		result[i] = PoissonSample(mu)
	}
	return result
}

//PoissonMutCoords
func PoissonMutCoordsFromArray(arr [][]int64) [][]int64 {
	var result [][]int64

	for m, col := range arr {
		for n, v := range col {
			if v > 0 {
				coords := []int64{int64(m), int64(n)}
				result = append(result, coords)
			}
		}
	}
	return result
}

//PoissonMutCoords
func PoissonMutCoords(mu float64, nSites, popSize int64) [][]int64 {
	var result [][]int64
	n := nSites * popSize

	for i := int64(0); i < n; i++ {
		if PoissonSample(mu) > 0 {
			var q, r = divmod(int64(i), int64(nSites))
			coords := []int64{q, r}
			result = append(result, coords)
		}
	}
	return result
}

func divmod(numerator, denominator int64) (quotient, remainder int64) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}

func PoissonSample(lambda float64) int64 {
	L := math.Exp(-1 * lambda)
	k := 0
	p := 1.
	for p > L {
		k++
		p *= rand.Float64()
	}
	return int64(k - 1)
}

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

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(PoissonMutCoords(0.0001, 10000, 10))
}
