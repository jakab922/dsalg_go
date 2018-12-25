package dsalg

import (
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

const (
	TestCount = 1000
	MaxVal    = int64(10000)
	DataFile  = "prime_data"
)

func pow(base, exp int) int {
	ret := 1
	for exp > 0 {
		ret *= base
		exp -= 1
	}

	return ret
}

func readPrimes(t *testing.T) []int {
	data, err := ioutil.ReadFile(DataFile)
	if err != nil {
		t.Error(err)
	}

	fields := strings.Fields(string(data))

	ret := make([]int, 0)
	for _, elem := range fields {
		tmp, err := strconv.ParseInt(elem, 10, 64)
		if err != nil {
			t.Error(err)
		}

		ret = append(ret, int(tmp))
	}

	return ret
}

func TestIsPrime(t *testing.T) {
	sieve := NewSieve(int(MaxVal))
	cp := 0
	primes := readPrimes(t)

	for i := 1; i < int(MaxVal); i++ {
		if cp < len(primes) && primes[cp] == i {
			cp += 1
			if isPrime, _ := sieve.IsPrime(i); !isPrime {
				t.Errorf("%v is a prime but the sieve says otherwise.", i)
			}
		} else {
			if isPrime, _ := sieve.IsPrime(i); isPrime {
				t.Errorf("%v is not a prime while the sieve says it is.", i)
			}
		}
	}
}

func TestFactors(t *testing.T) {
	sieve := NewSieve(int(MaxVal))
	for i := 0; i < TestCount; i++ {
		expected := int(rand.Int63n(MaxVal))
		factors, err := sieve.Factors(expected)
		if err != nil {
			t.Error(err)
		}
		actual := 1
		for _, factor := range factors {
			actual *= pow(factor.Base, factor.Exp)
		}

		if expected != actual {
			t.Errorf("After factoring %v != %v", expected, actual)
		}
	}
}
