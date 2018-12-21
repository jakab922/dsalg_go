package dsalg

import (
	"log"
	"math"
	"math/rand"
	"testing"
)

const (
	maxBase   = int64(2 << 10)
	maxExp    = int64(10)
	maxExpMod = int64(10000)
	maxMod    = int64(2 << 28)
	testCount = 1000
)

func TestFastpow(t *testing.T) {
	for i := 0; i < testCount; i++ {
		base := rand.Int63n(maxBase)
		exp := rand.Int63n(maxExp)
		exp2 := exp
		expected := One
		bad := false
		for exp2 > Zero {
			if math.MaxInt64/base < expected {
				bad = true
				log.Printf("%v^%v are too big to fit on 64 bits\n", base, exp)
				break
			}
			expected *= base
			exp2 -= 1
		}
		if bad {
			continue
		}
		calculated := Fastpow(base, exp)
		if expected != calculated {
			t.Errorf("The values for %v^%v differ. Expected: %v. Calculated: %v.", base, exp, expected, calculated)
		}
	}
}

func TestFastpowMod(t *testing.T) {
	for i := 0; i < testCount; i++ {
		base := rand.Int63n(maxBase)
		exp := rand.Int63n(maxExpMod)
		mod := rand.Int63n(maxMod)
		exp2 := exp
		expected := One
		for exp2 > Zero {
			expected = (expected * base) % mod
			exp2 -= 1
		}
		calculated := FastpowMod(base, exp, mod)
		if expected != calculated {
			t.Errorf("The values for %v^%v %% %v differ. Expected: %v. Calculated: %v.", base, exp, mod, expected, calculated)
		}
	}

}
