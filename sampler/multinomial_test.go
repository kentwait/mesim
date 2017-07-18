package sampler

import (
	"math"
	"math/rand"
	"testing"
)

func TestMultinomial(t *testing.T) {
	n := 1000
	p := []float64{0.4, 0.3, 0.2, 0.1}
	times := 10000
	samples := make([][]int, times)
	seed := 0
	rand.Seed(int64(seed))

	for i := 0; i < times; i++ {
		samples[i] = MultinomialSample(n, p)
	}

	// Average
	colSum := make([]int, len(samples[0]))
	colMean := make([]float64, len(samples[0]))
	for _, row := range samples {
		for j, v := range row {
			colSum[j] += v
		}
	}
	for i, v := range colSum {
		colMean[i] = float64(v) / float64(times)
	}

	// Check
	for i, v := range colMean {
		if math.Abs(v-(p[i]*float64(n))) > 1 {
			t.Errorf("Multinomial(%d, %v): mean error of %d repeated trials is greater than 1, %v",
				n, p, times, colMean,
			)
		}
	}

}

func TestMultinomialLog(t *testing.T) {
	n := 1000
	p := []float64{
		0.4,
		0.3,
		0.2,
		0.1,
	}
	logP := []float64{
		math.Log(0.4),
		math.Log(0.3),
		math.Log(0.2),
		math.Log(0.1),
	}
	times := 10000
	samples := make([][]int, times)
	seed := 0
	rand.Seed(int64(seed))

	for i := 0; i < times; i++ {
		samples[i] = MultinomialLogSample(n, logP)
	}

	// Average
	colSum := make([]int, len(samples[0]))
	colMean := make([]float64, len(samples[0]))
	for _, row := range samples {
		for j, v := range row {
			colSum[j] += v
		}
	}
	for i, v := range colSum {
		colMean[i] = float64(v) / float64(times)
	}

	// Check
	for i, v := range colMean {
		if math.Abs(v-(p[i]*float64(n))) > 1 {
			t.Errorf("Multinomial(%d, %v): mean error of %d repeated trials is greater than 1, %v",
				n, p, times, colMean,
			)
		}
	}

}
