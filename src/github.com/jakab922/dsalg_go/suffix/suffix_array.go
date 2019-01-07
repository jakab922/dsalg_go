package dsalg

import (
	"fmt"
)

const (
	DOLLAR = '$'
)

type SuffixArray struct {
	r     []rune
	order []int
}

func initialOrders(r []rune, alphabet_size int) []int {
	// Basically a stable sort for the one element substrings of the string
	// Runs in O(len(r) + alphabet_size) thus if the size of the alphabet is small
	// this runs in linear time in the length of the string.
	count := make([]rune, alphabet_size)

	for _, el := range r {
		count[el] += 1
	}

	for i := 1; i < alphabet_size; i++ {
		count[i] += count[i-1] // count[i] is the number of occurances of all characters not bigger than i + minVal
	}

	lr := len(r)
	order := make([]int, lr)

	for i := lr - 1; i > -1; i-- {
		c := r[i]
		count[c] -= 1
		order[count[c]] = i
	}

	return order
}

func updateOrders(r []rune, l int, order, class []int) []int {
	// This is again a stable sort. We have cyclic shifts(CS) of length l
	// already sorted and know their equivalence class.
	// We want to calculate the order of CS of length 2l. Let's call the first
	// half of this prefix and the second suffix.
	// The running time again is O(len(r))
	lr := len(r)
	count := make([]int, lr)
	for i := 0; i < lr; i++ {
		count[class[i]] += 1
	}
	for i := 1; i < lr; i++ {
		count[i] += count[i-1] // count[i] holds the number of prefixes which have class <= i
	}

	newOrder := make([]int, lr)
	for i := lr - 1; i > -1; i-- {
		start := (order[i] - l + lr) % lr // We take the index of the prefix part of the suffix of order[i]
		cl := class[start]
		count[cl] -= 1
		newOrder[count[cl]] = start // The new order of the 2l CS starting at index "start" is the number of remaining prefixes that have a <= class of the current one thus will come before it since also their suffix order is smaller.
	}

	return newOrder
}

func initialClasses(r []rune, order []int) []int {
	// Computes the equivalence class for a given rune slice and
	// the current order of suffixes.
	// Suppose we have a range i..j. The order order[i]..order[j]
	// are in the same equivalence class if and only if their
	// associated cyclic shift starts with the same letter.
	lr := len(r)
	class := make([]int, lr)
	class[order[0]] = 0

	for i := 1; i < lr; i++ {
		delta := 0
		if r[order[i]] != r[order[i-1]] {
			delta = 1
		}
		class[order[i]] = class[order[i-1]] + delta
	}

	return class
}

func updateClasses(newOrder, class []int, l int) []int {
	// We take the newOrder calculated in the updateOrders method and
	// calculate the prefix and suffix class of each cyclic shift of length
	// 2 * l in the order defined by newOrder. If the pair is different
	// for adjacent elements then they will be in a different class
	// otherwise they stay in the same class. Run in O(len(newOrder))
	// which is also the length of the string.
	n := len(newOrder)
	newClass := make([]int, n)
	newClass[newOrder[0]] = 0

	for i := 1; i < n; i++ {
		ff, fs := newOrder[i], newOrder[i-1]
		sf, ss := (ff+l)%n, (fs+l)%n
		delta := 0
		if class[ff] != class[fs] || class[sf] != class[ss] {
			delta = 1
		}
		newClass[ff] = newClass[fs] + delta
	}

	return newClass
}

func NewSuffixArray(s string, alphabet_size int) *SuffixArray {
	r := []rune(s)
	r = append(r, DOLLAR)
	lr := len(r)

	order := initialOrders(r, alphabet_size)
	class := initialClasses(r, order)
	l := 1

	// Note that cyclic shifts of length len(r) have the same ordering as the
	// suffixes of the string.
	for l < lr {
		order = updateOrders(r, l, order, class)
		class = updateClasses(order, class, l)
		l <<= 1
	}

	return &SuffixArray{r, order}
}

func (sa *SuffixArray) Order() []int {
	return sa.order
}

func main() {
	s := "baaabbaaa"
	r := []rune(s)
	r = append(r, DOLLAR)

	suffix := NewSuffixArray(s, 256)
	fmt.Printf("suffix.order: %v\n", suffix.order)
	for _, el := range suffix.order {
		fmt.Println(string(r[el:]))
	}
}
