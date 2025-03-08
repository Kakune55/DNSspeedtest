package utils

import (
	"math"
)

// Stats holds statistical data for DNS response times.
type Stats struct {
	Mean   float64
	StdDev float64
	Max    float64
	Min    float64
}

// CalculateStats computes the statistics for a given slice of response times.
func CalculateStats(times []float64) Stats {
	if len(times) == 0 {
		return Stats{}
	}

	var sum, sumSq float64
	min := times[0]
	max := times[0]

	for _, time := range times {
		sum += time
		sumSq += time * time
		if time < min {
			min = time
		}
		if time > max {
			max = time
		}
	}

	mean := sum / float64(len(times))
	variance := (sumSq / float64(len(times))) - (mean * mean)
	stdDev := math.Sqrt(variance)

	return Stats{
		Mean:   mean,
		StdDev: stdDev,
		Max:    max,
		Min:    min,
	}
}