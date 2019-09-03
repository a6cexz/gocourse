package statistic

// Average returns average
func Average(xs []float64) float64 {
	total := Sum(xs)
	return total / float64(len(xs))
}

// Sum returns sum
func Sum(xs []float64) float64 {
	sum := float64(0)
	for _, x := range xs {
		sum += x
	}
	return sum
}
