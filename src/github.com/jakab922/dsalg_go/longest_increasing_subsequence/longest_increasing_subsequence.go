package dsalg

import (
	"sort"
)

func LongestIncreasingSubsequence(seq []int) (int, int, []int) {
	n := len(seq)
	indexes, prev := make([]int, n+1), make([]int, n+1)

	for i := 0; i < n+1; i++ {
		indexes[i] = -1
		prev[i] = -1
	}

	indexes[1] = 0
	longest := 1

	for i, el := range seq {
		first := sort.Search(longest+1, func(j int) bool {
			return j != 0 && el < seq[indexes[j]]
		})
		if indexes[first-1] != -1 && el <= seq[indexes[first-1]] {
			// Not bigger than the value pointed by the previous index
			continue
		}

		prev[i] = indexes[first-1]
		indexes[first] = i

		if first > longest {
			longest = first
		}
	}

	return longest, indexes[longest], prev
}
