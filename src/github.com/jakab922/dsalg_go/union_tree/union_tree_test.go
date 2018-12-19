package dsalg

import (
	"math/rand"
	"testing"
	"time"
)

const (
	COUNT    = 1000
	MAX_SIZE = int32(500)
)

func findRoute(edges []map[int32]bool, c int32, d int32) bool {
	if c == d {
		return true
	}
	stack := make([]int32, 0)
	was := make([]bool, len(edges))

	stack = append(stack, c)
	was[c] = true

	for len(stack) != 0 {
		curr := stack[0]
		stack = stack[1:]

		for k, _ := range edges[curr] {
			if k == d {
				return true
			}

			if !was[k] {
				was[k] = true
				stack = append(stack, k)
			}
		}
	}

	return false
}

func TestItWorks(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < COUNT; i++ {
		size := rand.Int31n(MAX_SIZE) + int32(1)
		edges := make([]map[int32]bool, size)
		for i := 0; i < int(size); i++ {
			edges[i] = make(map[int32]bool)
		}
		was := make(map[int32]map[int32]bool)
		ut := NewUtree(int(size))

		for j := 0; j < int(size-1); j++ {
			a, b := rand.Int31n(size), rand.Int31n(size)

			for {
				if a == b {
					a, b = rand.Int31n(size), rand.Int31n(size)
					continue
				}

				if a > b {
					a, b = b, a
				}

				if _, ok := was[a]; ok {
					if _, ok = was[a][b]; ok {
						a, b = rand.Int31n(size), rand.Int31n(size)
						continue
					} else {
						was[a][b] = true
						break
					}
				} else {
					was[a] = make(map[int32]bool)
					was[a][b] = true
					break
				}
			}

			edges[a][b] = true
			edges[b][a] = true

			ut.Union(int(a), int(b))
			c, d := rand.Int31n(size), rand.Int31n(size)

			if findRoute(edges, c, d) != (ut.Find(int(c)) == ut.Find(int(d))) {
				t.Errorf("Failed for the above graph with edge %v-%v", c, d)
				return
			}
		}
	}
}
