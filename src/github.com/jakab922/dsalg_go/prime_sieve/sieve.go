package dsalg

import (
	"errors"
	"fmt"
)

type Sieve struct {
	n             int
	first, primes []int
}

type Factor struct {
	Base, Exp int
}

func NewSieve(n int) *Sieve {
	first := make([]int, n+1)
	primes := make([]int, 0)

	for i := 2; i <= n; i++ {
		if first[i] == 0 {
			first[i] = i
			primes = append(primes, i)
		}

		for j := 0; j < len(primes) && primes[j] <= first[i] && i*primes[j] <= n; j++ {
			first[i*primes[j]] = primes[j]
		}
	}

	ret := Sieve{n, first, primes}

	return &ret
}

func (s *Sieve) IsPrime(x int) (bool, error) {
	if s.n < x {
		return false, errors.New(fmt.Sprintf("The largest supported number by this sieve is %v", s.n))
	}
	if x < s.primes[0] || x > s.primes[len(s.primes)-1] {
		return false, nil
	} else if x == s.primes[0] {
		return true, nil
	}
	low, high := 0, len(s.primes)-1
	for high-low > 1 {
		mid := (low + high) / 2
		if s.primes[mid] < x {
			low = mid
		} else {
			high = mid
		}
	}
	return s.primes[high] == x, nil
}

func (s *Sieve) Factors(x int) ([]Factor, error) {
	factors := make([]Factor, 0)

	if s.n < x {
		return factors, errors.New(fmt.Sprintf("The largest supported number by this sieve is %v", s.n))
	}

	for x > 1 {
		pi := s.first[x]
		exp := 0
		for x > 1 && s.first[x] == pi {
			exp += 1
			x /= pi
		}
		factors = append(factors, Factor{pi, exp})
	}

	return factors, nil
}
