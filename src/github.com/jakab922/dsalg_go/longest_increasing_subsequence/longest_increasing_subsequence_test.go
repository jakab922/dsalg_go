package dsalg

import (
	"math/rand"
	"testing"
	"time"
)

const (
	maxLen    = 50
	maxVal    = 30
	testCount = 10000
)

func randomSlice(length, low, high int) []int {
	ret := make([]int, length)
	for i := 0; i < length; i++ {
		ret[i] = rand.Int()%(high-low+1) + low
	}

	return ret
}

type Tuple struct {
	index, length int
}

func Brute(seq []int) int {
	l := len(seq)
	stack := make([]*Tuple, 0)
	for i := range seq {
		stack = append(stack, &Tuple{i, 1})
	}

	longest := 1
	for len(stack) > 0 {
		curr := stack[0]
		stack = stack[1:]
		if curr.length > longest {
			longest = curr.length
		}
		for i := curr.index + 1; i < l; i++ {
			if seq[curr.index] < seq[i] {
				stack = append(stack, &Tuple{i, curr.length + 1})
			}
		}
	}

	return longest
}

func check(seq []int, longest, last int, prev []int) bool {
	tmp := make([]int, 0)
	c := last

	for c != -1 {
		tmp = append(tmp, seq[c])
		if c <= prev[c] {
			return false
		}
		c = prev[c]
	}

	if len(tmp) != longest {
		return false
	}

	for i := 0; i < longest-1; i++ {
		if tmp[i] <= tmp[i+1] {
			return false
		}
	}

	return true
}

func TestLongestIncreasingSubsequence(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < testCount; i++ {
		seq := randomSlice(rand.Int()%maxLen+1, 0, rand.Int()%maxVal+1)

		expected := Brute(seq)

		actual, last, prev := LongestIncreasingSubsequence(seq)

		if expected != actual {
			t.Errorf("The length of the longest sequence is reported to be %v while it's %v", actual, expected)
			return
		}

		if !check(seq, actual, last, prev) {
			t.Errorf("The returned information doesn't define an increasing subsequence of length %v", actual)
			return
		}
	}
}
