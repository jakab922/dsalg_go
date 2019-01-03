package dsalg

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

const (
	LowerCase = "abcdefghijklmnopqrstuvwxyz"
	INSERT    = 0
	DELETE    = 1
)

type Operation struct {
	ty   int
	word string
}

func RandomLowerCase(length, letterCount int) string {
	tmp := make([]byte, length)

	for i := 0; i < length; i++ {
		tmp[i] = LowerCase[rand.Int()%letterCount]
	}

	return string(tmp)
}

func TestSmallAlphabet(t *testing.T) {
	MaxOp := 100
	MaxWordLen := 10
	LetterCount := 2
	InsertProb := 0.5
	TestCount := 10000
	testCommon(t, MaxOp, MaxWordLen, LetterCount, InsertProb, TestCount)
}

func TestNormalAlphabet(t *testing.T) {
	MaxOp := 100
	MaxWordLen := 10
	LetterCount := 26
	InsertProb := 0.5
	TestCount := 10000
	testCommon(t, MaxOp, MaxWordLen, LetterCount, InsertProb, TestCount)
}

func TestLotOfOperations(t *testing.T) {
	MaxOp := 10000
	MaxWordLen := 10
	LetterCount := 26
	InsertProb := 0.8
	TestCount := 5
	testCommon(t, MaxOp, MaxWordLen, LetterCount, InsertProb, TestCount)
}

func testCommon(t *testing.T, MaxOp, MaxWordLen, LetterCount int, InsertProb float64, TestCount int) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < TestCount; i++ {
		opCount := rand.Int() % MaxOp
		operations := make([]*Operation, opCount)
		for j := 0; j < opCount; j++ {
			word := RandomLowerCase(rand.Int()%MaxWordLen+1, LetterCount)
			if rand.Float64() < InsertProb {
				operations[j] = &Operation{INSERT, word}
			} else {
				operations[j] = &Operation{DELETE, ""}
			}
		}
		trie := NewTrie()

		coll := make(map[string]bool)
		for _, operation := range operations {
			if operation.ty == INSERT {
				coll[operation.word] = true
				trie.Insert(operation.word)
			} else {
				if trie.Len() == 0 {
					continue
				}
				var word string
				index := rand.Int() % trie.Len()
				k := 0
				for cword := range trie.Iterate() {
					if k == index {
						word = cword
					}
					k += 1
				}
				coll[word] = false
				trie.Remove(word)
			}

			expected := make([]string, 0)
			for word, isIn := range coll {
				if isIn {
					expected = append(expected, word)
				}
				if trie.Find(word) != isIn {
					t.Errorf("The word %q is in the trie while it shouldn't be!", word)
				}
			}
			sort.Slice(expected, func(i, j int) bool {
				return expected[i] < expected[j]
			})

			actual := make([]string, 0)
			for word := range trie.Iterate() {
				actual = append(actual, word)
			}

			sort.Slice(actual, func(i, j int) bool {
				return actual[i] < actual[j]
			})

			if len(expected) != trie.Len() {
				t.Errorf("The length of expected and actual differs: %v != %v", len(expected), len(actual))
			}

			for i := range actual {
				if expected[i] != actual[i] {
					t.Errorf("The elements at index %v differ: %v != %v", i, expected[i], actual[i])
					return
				}
			}
		}
	}
}
