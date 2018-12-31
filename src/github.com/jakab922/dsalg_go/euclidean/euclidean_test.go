package dsalg

import (
	"math/rand"
	"testing"
	"time"
)

const (
	maxVal    = 10000
	testCount = 10000
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func TestEuclidean(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < testCount; i++ {
		a := int(rand.Int63n(int64(maxVal)))
		b := int(rand.Int63n(int64(maxVal)))

		calculated := Euclidean(a, b)
		expected := 1

		for i := max(a, b); i > 0; i-- {
			if a%i == 0 && b%i == 0 {
				expected = i
				break
			}
		}

		if expected != calculated {
			t.Errorf("The expected value for gcd(%v, %v) is: %v while the calculated is: %v", a, b, expected, calculated)
		}
	}
}

func TestExtendedEuclidean(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < testCount; i++ {
		a := int(rand.Int63n(int64(maxVal)))
		b := int(rand.Int63n(int64(maxVal)))

		calculated, aCoeff, bCoeff := ExtendedEuclidean(a, b)
		expected := 1

		for i := max(a, b); i > 0; i-- {
			if a%i == 0 && b%i == 0 {
				expected = i
				break
			}
		}

		if expected != calculated {
			t.Errorf("The expected value for gcd(%v, %v) is: %v while the calculated is: %v", a, b, expected, calculated)
		}
		if a*aCoeff+b*bCoeff != calculated {
			t.Errorf("%v * %v + %v * %v != %v", a, aCoeff, b, bCoeff, calculated)
		}
	}
}
