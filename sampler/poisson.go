// Copyright Kent Kawashima 2017
// All Rights Reserved

package sampler

import (
	"math"
	"math/rand"
	utils "mesim/utils"
)

// PoissonMutArray samples from a Poisson distribution and
// returns a 2-d array of Poisson random variables
// whose rows represent individuals and columns represent sites.
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

// poissonMutArray is the base function of PoissonMutArray.
// It calls the PoissonSampler function to generate a set of Poisson
// random variables.
func poissonMutArray(mu float64, n int64) []int64 {
	result := make([]int64, n)

	for i := int64(0); i < n; i++ {
		result[i] = PoissonSampler(mu)
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
		if PoissonSampler(mu) > 0 {
			var q, r = utils.DivMod(int64(i), int64(nSites))
			coords := []int64{q, r}
			result = append(result, coords)
		}
	}
	return result
}

// PoissonSampler return a pseudorandom sample from a Poisson
// distribution of lambda using the Knuth algorithm.
func PoissonSampler(lambda float64) int64 {
	L := math.Exp(-1 * lambda)
	k := 0
	p := 1.
	for p > L {
		k++
		p *= rand.Float64()
	}
	return int64(k - 1)
}
