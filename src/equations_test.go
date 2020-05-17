package main

import (
	"math"
	"strconv"
	"testing"
)

var testForPGivenF = []struct {
	i        float64
	n        float64
	expected float64
}{
	{0.1, 5, 0.6209},
	{0.2, 1.2, 0.8035},
	{2, 1.2, 0.2676},
	{2, 3, 0.0370},
	{0.9999999999999999999, 1.9999999999999999999999999, 0.25},
	{99999, 99999, 0},
}

var testForFGivenP = []struct {
	i        float64
	n        float64
	expected float64
}{
	{0.1, 5, 1.6105},
	{0.2, 1.2, 1.2446},
	{2, 1.2, 3.7372},
	{2, 3, 27},
	{0.9999999999999999999, 1.9999999999999999999999999, 4},
	{99999, 99999, math.Inf(1)},
}

var testForPGivenA = []struct {
	i        float64
	n        float64
	expected float64
}{
	{0.1, 5, 3.7908},
	{0.2, 1.2, 0.9825},
	{2, 1.2, 0.3662},
	{2, 3, 0.4815},
	{0.9999999999999999999, 1.9999999999999999999999999, 0.75},
	{99999, 99999, math.NaN()},
}

var testForAGivenP = []struct {
	i        float64
	n        float64
	expected float64
}{
	{0.1, 5, 0.2638},
	{0.2, 1.2, 1.0178},
	{2, 1.2, 2.7307},
	{2, 3, 2.0769},
	{0.9999999999999999999, 1.9999999999999999999999999, 1.3333},
	{99999, 99999, math.NaN()},
}

var testForFGivenA = []struct {
	i        float64
	n        float64
	expected float64
}{
	{0.1, 5, 6.1051},
	{0.2, 1.2, 1.2228},
	{2, 1.2, 1.3686},
	{2, 3, 13},
	{0.9999999999999999999, 1.9999999999999999999999999, 3.0000},
	{99999, 99999, math.Inf(1)},
}

var testForAGivenF = []struct {
	i        float64
	n        float64
	expected float64
}{
	{0.1, 5, 0.1638},
	{0.2, 1.2, 0.8178},
	{2, 1.2, 0.7307},
	{2, 3, 0.0769},
	{0.9999999999999999999, 1.9999999999999999999999999, 0.3333},
	{99999, 99999, 0},
}

var testForPGivenG = []struct {
	i        float64
	n        float64
	expected float64
}{
	{0.1, 5, 6.8618},
	{0.2, 1.2, 0.0917},
	{2, 1.2, 0.0226},
	{2, 3, 0.1852},
	{0.9999999999999999999, 1.9999999999999999999999999, 0.2500},
	{99999, 99999, 0},
}

var testForAGivenG = []struct {
	i        float64
	n        float64
	expected float64
}{
	{0.1, 5, 1.8101},
	{0.2, 1.2, 0.0933},
	{2, 1.2, 0.0616},
	{2, 3, 0.3846},
	{0.9999999999999999999, 1.9999999999999999999999999, 0.3333},
	{99999, 99999, math.NaN()},
}

var testForPGivenAWithJ = []struct {
	i        float64
	j        float64
	n        float64
	expected float64
}{
	{0.1, 0.2, 5, 5.4505},
	{0.1, -0.2, 5, 2.6551},
	{-0.1, 0.2, 5, 10.7133},
	{99999, 99999, 99999, math.NaN()},
}

var testForFGivenAWithJ = []struct {
	i        float64
	j        float64
	n        float64
	expected float64
}{
	{0.1, 0.2, 5, 8.7781},
	{0.1, -0.2, 5, 4.2761},
	{-0.1, 0.2, 5, 6.3261},
	{99999, 99999, 99999, math.NaN()},
}

func TestPGivenF(t *testing.T) {
	for i, testCase := range testForPGivenF {
		testName := "TestPGivenF_" + strconv.Itoa(i+1)

		t.Run(testName, func(t *testing.T) {
			actual := roundTo4Decimal(PGivenF(testCase.i, testCase.n))

			if actual != testCase.expected {
				t.Errorf("got %f, want %f", actual, testCase.expected)
			}
		})
	}
}

func TestFGivenP(t *testing.T) {
	for i, testCase := range testForFGivenP {
		testName := "TestFGivenP_" + strconv.Itoa(i+1)

		t.Run(testName, func(t *testing.T) {
			actual := roundTo4Decimal(FGivenP(testCase.i, testCase.n))

			if actual != testCase.expected {
				t.Errorf("got %f, want %f", actual, testCase.expected)
			}
		})
	}
}

func TestPGivenA(t *testing.T) {
	for i, testCase := range testForPGivenA {
		testName := "TestPGivenA_" + strconv.Itoa(i+1)

		t.Run(testName, func(t *testing.T) {
			actual := roundTo4Decimal(PGivenA(testCase.i, testCase.n))

			// NaN == NaN will always return false in Go, so I have to do this
			if math.IsNaN(testCase.expected) {
				if !math.IsNaN(actual) {
					t.Errorf("got %f, want %f", actual, testCase.expected)
				}
			} else if actual != testCase.expected {
				t.Errorf("got %f, want %f", actual, testCase.expected)
			}
		})
	}
}

func TestAGivenP(t *testing.T) {
	for i, testCase := range testForAGivenP {
		testName := "TestAGivenP_" + strconv.Itoa(i+1)

		t.Run(testName, func(t *testing.T) {
			actual := roundTo4Decimal(AGivenP(testCase.i, testCase.n))

			// NaN == NaN will always return false in Go, so I have to do this
			if math.IsNaN(testCase.expected) {
				if !math.IsNaN(actual) {
					t.Errorf("got %f, want %f", actual, testCase.expected)
				}
			} else if actual != testCase.expected {
				t.Errorf("got %f, want %f", actual, testCase.expected)
			}
		})
	}
}

func TestFGivenA(t *testing.T) {
	for i, testCase := range testForFGivenA {
		testName := "TestFGivenA_" + strconv.Itoa(i+1)

		t.Run(testName, func(t *testing.T) {
			actual := roundTo4Decimal(FGivenA(testCase.i, testCase.n))

			if actual != testCase.expected {
				t.Errorf("got %f, want %f", actual, testCase.expected)
			}
		})
	}
}

func TestAGivenF(t *testing.T) {
	for i, testCase := range testForAGivenF {
		testName := "TestAGivenF_" + strconv.Itoa(i+1)

		t.Run(testName, func(t *testing.T) {
			actual := roundTo4Decimal(AGivenF(testCase.i, testCase.n))

			if actual != testCase.expected {
				t.Errorf("got %f, want %f", actual, testCase.expected)
			}
		})
	}
}

func TestPGivenG(t *testing.T) {
	for i, testCase := range testForPGivenG {
		testName := "TestPGivenG_" + strconv.Itoa(i+1)

		t.Run(testName, func(t *testing.T) {
			actual := roundTo4Decimal(PGivenG(testCase.i, testCase.n))

			if actual != testCase.expected {
				t.Errorf("got %f, want %f", actual, testCase.expected)
			}
		})
	}
}

func TestAGivenG(t *testing.T) {
	for i, testCase := range testForAGivenG {
		testName := "TestAGivenG_" + strconv.Itoa(i+1)

		t.Run(testName, func(t *testing.T) {
			actual := roundTo4Decimal(AGivenG(testCase.i, testCase.n))

			// NaN == NaN will always return false in Go, so I have to do this
			if math.IsNaN(testCase.expected) {
				if !math.IsNaN(actual) {
					t.Errorf("got %f, want %f", actual, testCase.expected)
				}
			} else if actual != testCase.expected {
				t.Errorf("got %f, want %f", actual, testCase.expected)
			}
		})
	}
}

func TestPGivenAWithJ(t *testing.T) {
	for i, testCase := range testForPGivenAWithJ {
		testName := "TestPGivenAWithJ_" + strconv.Itoa(i+1)

		t.Run(testName, func(t *testing.T) {
			actual := roundTo4Decimal(PGivenAWithJ(testCase.i, testCase.j, testCase.n))

			// NaN == NaN will always return false in Go, so I have to do this
			if math.IsNaN(testCase.expected) {
				if !math.IsNaN(actual) {
					t.Errorf("got %f, want %f", actual, testCase.expected)
				}
			} else if actual != testCase.expected {
				t.Errorf("got %f, want %f", actual, testCase.expected)
			}
		})
	}
}

func TestFGivenAWithJ(t *testing.T) {
	for i, testCase := range testForFGivenAWithJ {
		testName := "TestFGivenAWithJ_" + strconv.Itoa(i+1)

		t.Run(testName, func(t *testing.T) {
			actual := roundTo4Decimal(FGivenAWithJ(testCase.i, testCase.j, testCase.n))

			// NaN == NaN will always return false in Go, so I have to do this
			if math.IsNaN(testCase.expected) {
				if !math.IsNaN(actual) {
					t.Errorf("got %f, want %f", actual, testCase.expected)
				}
			} else if actual != testCase.expected {
				t.Errorf("got %f, want %f", actual, testCase.expected)
			}
		})
	}
}

// a helper function that helps me to round up to 4 d.p.
// since for ece472, 4 d.p. is all that matters
func roundTo4Decimal(result float64) float64 {
	return math.Round(result*10000) / 10000
}
