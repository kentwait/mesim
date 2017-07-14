// Copyright Kent Kawashima 2017
// All Rights Reserved

package main

import (
	"log"
	"math/rand"
	sampler "mesim/sampler"
	"time"
)

// Evolve
func Evolve(ancArray *[]int64, rateMatrix [][]float64) {
	for i, char := range *ancArray {
		(*ancArray)[i] = sampler.MultinomialWhere(int64(1), rateMatrix[char], int64(1))[0]
	}
}

func EvolveAlt(ancArray *[]int64, rateMatrix [][]float64) {
	mu := float64(0)
	for x, row := range rateMatrix {
		for y, v := range row {
			if x != y {
				mu += v
			}
		}
	}

	mutCoords := sampler.PoissonMutCoords(mu, int64(len(*ancArray)), int64(1))
	var tmpArray []int64
	if len(mutCoords) > 0 {
		for _, yPos := range mutCoords[1] {
			tmpArray = append(tmpArray, (*ancArray)[yPos])
		}
		Evolve(&tmpArray, rateMatrix)
		for i, xPos := range mutCoords[0] {
			(*ancArray)[xPos] = tmpArray[i]
		}
	}

}

func main() {
	start := time.Now()
	rand.Seed(time.Now().UTC().UnixNano())
	a := 0.999991
	b := 0.000003
	rateMatrix := [][]float64{
		[]float64{a, b, b, b},
		[]float64{b, a, b, b},
		[]float64{b, b, a, b},
		[]float64{b, b, b, a},
	}
	arr := make([]int64, 1e8)
	// fmt.Println(arr)
	Evolve(&arr, rateMatrix)
	// fmt.Println(arr)
	log.Printf("Time: %s", time.Since(start))

	start = time.Now()
	arr = make([]int64, 1e8)
	// fmt.Println(arr)
	EvolveAlt(&arr, rateMatrix)
	// fmt.Println(arr)
	log.Printf("Time: %s", time.Since(start))
}
