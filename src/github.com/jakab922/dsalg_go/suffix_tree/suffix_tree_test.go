package dsalg

import (
	sarr "github.com/jakab922/dsalg_go/suffix_array"
	"math/rand"
	"strings"
	"testing"
	"time"
)

const (
	maxLengthText = 500
	maxGenerated  = 500
	testCount     = 10000
	alphabet      = "abcdefghijklmnopqrstuvwxyz"
	debug         = false
)

func randomString(length int) string {
	la := len(alphabet)
	tmp := make([]byte, length)

	for i := 0; i < length; i++ {
		tmp[i] = alphabet[rand.Int()%la]
	}

	return string(tmp)
}

func TestLCP(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < testCount; i++ {
		s := randomString(rand.Int()%maxLengthText + 1)
		suffixArray := sarr.NewSuffixArray(s, 256)
		r := suffixArray.R
		lr := len(r)
		order := suffixArray.Order

		lcp := getLCP(r, order)

		for i := 0; i < lr-1; i++ {
			fp, sp := order[i], order[i+1]
			total := 0
			for fp+total < lr && sp+total < lr && r[fp+total] == r[sp+total] {
				total += 1
			}

			if total != lcp[i] {
				t.Errorf("The longest common prefix length of %q and %q is %v while it's reported to be %v", string(r[order[i]:]), string(r[order[i+1]:]), total, lcp[i])
			}
		}
	}
}

func TestSuffixTree(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < testCount; i++ {
		s := randomString(rand.Int()%maxLengthText + 1)
		r := []rune(s)
		lr := len(r)
		suffixTree := NewSuffixTree(s, 256)
		if debug {
			suffixTree.Show()
		}

		for j := 0; j < rand.Int()%maxGenerated+1; j++ {
			if rand.Int()%2 == 1 {
				notSubstring := randomString(rand.Int()%(2*lr) + 1)
				for strings.Contains(s, notSubstring) {
					notSubstring = randomString(rand.Int()%(2*lr) + 1)
				}

				if suffixTree.Contains(notSubstring) {
					t.Errorf("%q is not a substring of %q", notSubstring, s)
				}
			} else {
				low := rand.Int() % lr
				high := rand.Int() % lr
				if high < low {
					low, high = high, low
				}
				substring := string(r[low : high+1])

				if !suffixTree.Contains(substring) {
					t.Errorf("%q is a substring of %q", substring, s)
				}
			}
		}
	}
}

func BenchmarkSuffixTree(b *testing.B) {
	s := randomString(100000)
	for i := 0; i < b.N; i++ {
		NewSuffixTree(s, 256)
	}
}
