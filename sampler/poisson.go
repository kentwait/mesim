package sampler

import (
	"math"
	"math/rand"
	utils "mesim/utils"
)

// PoissonMutArray samples from a Poisson distribution and
// returns a 2-d array of Poisson random variables
// whose rows represent individuals and columns represent sites.
func PoissonMutArray(mu float64, nSites, popSize int) (result [][]int) {
	tmp := make([]int, nSites*popSize)
	n := nSites * popSize
	blockSize := int(1e6) //TODO : Optimize block size

	if n > blockSize {
		workers := 0
		resultChan := make(chan []int)

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

	for i := 0; i < n; i += nSites {
		result = append(result, tmp[i:i+nSites])
	}
	return result
}

// poissonMutArray is the base function of PoissonMutArray.
// It calls the PoissonSampler function to generate a set of Poisson
// random variables.
func poissonMutArray(mu float64, n int) []int {
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = PoissonSampler(mu)
	}
	return result
}

//PoissonMutCoords
func PoissonMutCoordsFromArray(arr [][]int) [][]int {
	var xArray, yArray, value []int

	for m, col := range arr {
		for n, v := range col {
			if v > 0 {
				xArray = append(xArray, m)
				yArray = append(yArray, n)
				value = append(value, v)
			}
		}
	}
	result := [][]int{xArray, yArray, value}
	return result
}

//PoissonMutCoords
func PoissonMutCoords(mu float64, nSites, popSize int) [][]int {
	var xArray, yArray, value []int
	n := nSites * popSize
	var v int
	for i := 0; i < n; i++ {
		v = PoissonSampler(mu)
		if v > 0 {
			var q, r = utils.DivMod(i, nSites)
			xArray = append(xArray, q)
			yArray = append(yArray, r)
			value = append(value, v)
		}
	}
	result := [][]int{xArray, yArray, value}
	return result
}

// PoissonSampler return a pseudorandom sample from a Poisson
// distribution of lambda using the Knuth algorithm.
func PoissonSampler(lambda float64) int {
	L := math.Exp(-1 * lambda)
	k := 0
	p := 1.
	for p > L {
		k++
		p *= rand.Float64()
	}
	return int(k - 1)
}
