package dsalg

import (
	"errors"
	"fmt"
	euc "github.com/jakab922/dsalg_go/euclidean"
)

func ModInv(a, mod int) (int, error) {
	gcd, aCoeff, _ := euc.ExtendedEuclidean(a, mod)
	if gcd != 1 {
		return 0, errors.New(fmt.Sprintf("%v doesn't have an inverse modulo %v\n", a, mod))
	}
	ret := aCoeff % mod
	if ret < 0 {
		return mod + ret, nil
	} else {
		return ret, nil
	}
}
