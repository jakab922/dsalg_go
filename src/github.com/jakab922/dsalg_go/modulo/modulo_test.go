package dsalg

import (
	euc "github.com/jakab922/dsalg_go/euclidean"
	"math/rand"
	"testing"
)

const (
	MaxVal    = 1000000000
	TestCount = 100000
)

func TestModInvRelativePrime(t *testing.T) {
	for i := 0; i < TestCount; i++ {
		a := int(rand.Int63n(int64(MaxVal))) + 2
		mod := int(rand.Int63n(int64(MaxVal))) + 2

		for euc.Euclidean(a, mod) != 1 {
			mod = int(rand.Int63n(int64(MaxVal))) + 2
		}

		inv, err := ModInv(a, mod)
		if err != nil {
			t.Errorf(err.Error())
		}

		if (a*inv)%mod != 1 {
			t.Errorf("The inverse of %v modulo %v is not %v\n", a, mod, inv)
		}
	}
}

func TestModInv(t *testing.T) {
	for i := 0; i < TestCount; i++ {
		a := int(rand.Int63n(int64(MaxVal))) + 2
		mod := int(rand.Int63n(int64(MaxVal))) + 2

		inv, err := ModInv(a, mod)
		if err != nil {
			if euc.Euclidean(a, mod) == 1 {
				t.Errorf("The function signalled an error for %v mod %v while the inverse exists\n", a, mod)
			}
		} else if (a*inv)%mod != 1 {
			t.Errorf("The inverse of %v modulo %v is not %v\n", a, mod, inv)
		}
	}
}
