package main

import "math"

// PGivenF is (P|F i, n) -- 1
func PGivenF(i, n float64) float64 {
	return math.Pow(1+i, -n)
}

// FGivenP is (F|P i, n) -- 2
func FGivenP(i, n float64) float64 {
	return 1 / PGivenF(i, n)
}

// PGivenA is (P|A i, n) -- 3
func PGivenA(i, n float64) float64 {
	numerator := math.Pow(1+i, n) - 1
	denominator := i * math.Pow(1+i, n)
	return numerator / denominator
}

// AGivenP is (A|P i, n) -- 4
func AGivenP(i, n float64) float64 {
	return 1 / PGivenA(i, n)
}

// FGivenA is (F|A i, n) -- 5
func FGivenA(i, n float64) float64 {
	return (math.Pow(1+i, n) - 1) / i
}

// AGivenF is (A|F i, n) -- 6
func AGivenF(i, n float64) float64 {
	return 1 / FGivenA(i, n)
}

// PGivenG is (P|G i, n) -- 7
func PGivenG(i, n float64) float64 {
	numerator := 1 - (1+n*i)*math.Pow(1+i, -n)
	denominator := math.Pow(i, 2)
	return numerator / denominator
}

// AGivenG is (A|G i, n) -- 8
func AGivenG(i, n float64) float64 {
	numerator := math.Pow(1+i, n) - (1 + n*i)
	denominator := i * (math.Pow(1+i, n) - 1)
	return numerator / denominator
}

// PGivenAWithJ is (P|A i, j, n) -- 9
func PGivenAWithJ(i, j, n float64) float64 {
	numerator := 1 - math.Pow(1+j, n)*math.Pow(1+i, -n)
	denominator := i - j
	return numerator / denominator
}

// FGivenAWithJ is (F|A i, j, n) -- 10
func FGivenAWithJ(i, j, n float64) float64 {
	numerator := math.Pow(1+i, n) - math.Pow(1+j, n)
	denominator := i - j
	return numerator / denominator
}
