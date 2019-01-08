package dsalg

import (
	"math/rand"
	"testing"
	"time"
)

const (
	maxLength = 22
	testCount = 100000
	alphabet  = "abcdefghijklmnopqrstuvwxyz"
)

func randomString(length int) string {
	la := len(alphabet)
	tmp := make([]byte, length)

	for i := 0; i < length; i++ {
		tmp[i] = alphabet[rand.Int()%la]
	}

	return string(tmp)
}

func TestSuffixArray(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < testCount; i++ {
		s := randomString(rand.Int()%maxLength + 1)
		r := []rune(s)
		r = append(r, DOLLAR)
		suffix := NewSuffixArray(s, 256)
		order := suffix.Order
		prev := string(r[order[0]:])
		for _, el := range order[1:] {
			curr := string(r[el:])
			if curr < prev {
				t.Errorf("The suffix %q should be bigger than %q", curr, prev)
			}
		}
	}
}
