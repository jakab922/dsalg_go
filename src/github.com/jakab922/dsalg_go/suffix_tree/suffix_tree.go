package dsalg

import (
	"fmt"
	sarr "github.com/jakab922/dsalg_go/suffix_array"
)

type SuffixTreeNode struct {
	first, last, length int
	followers           map[rune]*SuffixTreeNode
	prev                *SuffixTreeNode
}

func NewSuffixTreeNode(first, last, length int, followers map[rune]*SuffixTreeNode, prev *SuffixTreeNode) *SuffixTreeNode {
	if followers == nil {
		followers = make(map[rune]*SuffixTreeNode)
	}

	return &SuffixTreeNode{first, last, length, followers, prev}
}

type SuffixTree struct {
	root *SuffixTreeNode
	r    []rune
}

// Let's notate the longest common prefix of r[i:] and r[j:] with lcp(i, j)
// Then it's easy to prove that lcp(i, i + 1) >= lcp(i, j) for all j > i

func getInverseOrder(order []int) []int {
	// Calculates the inverse permutation of the order permutation
	lo := len(order)
	ret := make([]int, lo)

	for i := 0; i < lo; i++ {
		ret[order[i]] = i
	}

	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getAdjacentLCP(r []rune, i, j, minVal int) int {
	// Returns the least common prefix of 2 suffixes
	// adjacent by the suffix ordering.
	lr := len(r)
	ret := max(0, minVal)
	for i+ret < lr && j+ret < lr {
		if r[i+ret] != r[j+ret] {
			break
		}
		ret += 1
	}

	return ret
}

func getLCP(r []rune, order []int) []int {
	// Return the least common prefix of the suffixes
	// of the rune slice r where order defines the order
	// of the suffixes. The algorithm works in linear time
	// since in each iteration cLCP can get smaller by
	// at most 1 thus we can't make more than linear amount
	// of steps in getAdjacentLCP.
	lr := len(r)
	lcp := make([]int, lr-1)
	cLCP := 0

	iOrder := getInverseOrder(order)
	suffix := order[0]
	for i := 0; i < lr; i++ {
		sIndex := iOrder[suffix]
		if sIndex == lr-1 {
			cLCP = 0
			suffix = (suffix + 1) % lr
			continue
		}
		nSuffix := order[sIndex+1]
		// We have to argue why cLCP-1 is a good lower bound on the
		// next LCP. By adding one to the suffix in the previous
		// iteration of the loop we basically stripped away the first
		// character so stripping away the first character from the nSuffix
		// from the previous iteration gives a suffix which fulfils this
		// criteria and this proves the assumption on the lower bound.
		cLCP = getAdjacentLCP(r, suffix, nSuffix, cLCP-1)
		lcp[sIndex] = cLCP
		suffix = (suffix + 1) % lr
	}

	return lcp
}

func splitNode(r []rune, node *SuffixTreeNode, length int) *SuffixTreeNode {
	// node.length - length == node.last - newLast
	newLast := node.last + length - node.length
	ret := NewSuffixTreeNode(node.first, newLast, length, nil, node.prev)

	// Setting the link from the previous node
	ret.prev.followers[r[ret.first]] = ret

	// Repairing the second node
	node.prev = ret
	node.first = newLast + 1

	// Setting the only follower of the first node
	ret.followers[r[node.first]] = node

	return ret
}

func NewSuffixTree(s string, alphabet_size int) *SuffixTree {
	// This function builds a suffix tree from the input string
	// The complexity of the function depends on how fast we
	// can build a suffix array. The current suffix array builder
	// works in O(n * log n) where n is the length of the string.
	// Given a suffix array we can build a suffix tree in linear
	// time since if we used a "prev" link to get out from a node
	// we will never get back to that node again since in each
	// iteration of the loop we build a new node which is
	// the node associated with the current suffix.
	suffixArray := sarr.NewSuffixArray(s, alphabet_size)
	r, order := suffixArray.R, suffixArray.Order
	lr := len(r)

	// The least common prefixes tell us where we should branch off
	// from the tree if we go through the suffixes in order.
	lcp := getLCP(r, order)

	root := NewSuffixTreeNode(-1, -1, 0, nil, nil)
	empty := NewSuffixTreeNode(lr-1, lr-1, 1, nil, root)
	root.followers[sarr.DOLLAR] = empty
	curr := empty

	for i := 0; i < lr-1; i++ {
		if lcp[i] < curr.length {
			// lcp[i] < curr.length since lcp[i] > curr.length is not possible and for
			// curr.length == lcp[i] we don't need to do anything.
			// Being here means that the longest common prefix is shorter than the current suffix.
			for curr.prev != nil && lcp[i] < curr.prev.length+1 {
				curr = curr.prev
			}

			if lcp[i] < curr.length { // We need to split the node
				curr = splitNode(r, curr, lcp[i])
			}
		}
		// lcp[i] == curr.length so we can continue from the current node
		next := NewSuffixTreeNode(order[i+1]+lcp[i], lr-1, lr-order[i+1], nil, curr)
		curr.followers[r[order[i+1]+lcp[i]]] = next
		curr = next
	}

	return &SuffixTree{root, r}
}

type suffixTreeShowTuple struct {
	node   *SuffixTreeNode
	indent int
}

func (st *SuffixTree) Show() {
	stack := make([]*suffixTreeShowTuple, 0)
	stack = append(stack, &suffixTreeShowTuple{st.root, 0})

	for len(stack) > 0 {
		ls := len(stack)
		curr := stack[ls-1]
		stack = stack[:ls-1]
		pref := make([]byte, curr.indent)
		for i := 0; i < curr.indent; i++ {
			pref[i] = '*'
		}
		if curr.node.prev != nil {
			fmt.Printf("%v %q\n", string(pref), string(st.r[curr.node.first:curr.node.last+1]))
		}

		for _, follower := range curr.node.followers {
			stack = append(stack, &suffixTreeShowTuple{follower, curr.indent + 1})
		}
	}
}

func (st *SuffixTree) Contains(s string) bool {
	cnode, cp := st.root, 0

	for _, ch := range s {
		if cp > cnode.last {
			if nnode, ok := cnode.followers[ch]; ok {
				cnode = nnode
				cp = cnode.first
			} else {
				return false
			}
		}
		if st.r[cp] != ch {
			return false
		}
		cp += 1
	}

	return true
}
