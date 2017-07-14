// Copyright Kent Kawashima 2017
// All Rights Reserved

package main

import (
	"fmt"
	"math/rand"
	"time"

	sampler "mesim/sampler"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(sampler.PoissonMutCoords(0.0001, 10000, 10))
}
