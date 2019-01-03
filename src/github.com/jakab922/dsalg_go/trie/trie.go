package dsalg

import (
	"fmt"
	"strings"
)

const ()

type Node struct {
	followers  map[rune]*Node
	val        []rune
	prev       *Node
	terminator bool
}

func NewNode(val []rune, prev *Node, terminator bool, followers map[rune]*Node) *Node {
	if followers == nil {
		followers = make(map[rune]*Node)
	}
	node := Node{followers, val, prev, terminator}
	if prev != nil {
		prev.followers[val[0]] = &node
	}
	for _, follower := range followers {
		follower.prev = &node
	}
	return &node
}

func merge(a, b []rune) []rune {
	ret := make([]rune, 0)
	for _, el := range a {
		ret = append(ret, el)
	}

	for _, el := range b {
		ret = append(ret, el)
	}

	return ret
}

type Trie struct {
	root *Node
	size int
}

type TrieIterState struct {
	node *Node
	val  []rune
}

func NewTrie() *Trie {
	root := NewNode([]rune{}, nil, false, nil)
	return &Trie{root, 0}
}

func (t *Trie) find(word []rune) (*Node, int, int) {
	wLen := len(word)
	wp := 0
	cNode := t.root
	mismatch := false

	for true {
		cVal, cLen, cp := cNode.val, len(cNode.val), 0
		for cp < cLen && wp < wLen {
			if cVal[cp] == word[wp] {
				cp += 1
				wp += 1
			} else {
				mismatch = true
				break
			}
		}

		ok := false
		if wp != wLen {
			_, ok = cNode.followers[word[wp]]
		}
		if !mismatch && cp == cLen && wp != wLen && ok {
			cNode = cNode.followers[word[wp]]
		} else {
			return cNode, cp, wp
		}
	}
	return nil, 0, 0 // We will never get here so this is just for the compiler.
}

func (t *Trie) Insert(word string) {
	if word == "" {
		return
	}

	rWord := []rune(word)
	cNode, cp, wp := t.find(rWord)
	cVal := cNode.val
	cLen := len(cVal)
	wLen := len(rWord)

	if cp == cLen {
		if wp == wLen {
			// The word is either already in the Trie or the current
			// node is a good node word the suffix of the word
			// so we set the terminator flag
			if !cNode.terminator {
				t.size += 1
			}
			cNode.terminator = true
		} else {
			// The current word is a continuation of a word in the Trie
			NewNode(rWord[wp:], cNode, true, nil)
			t.size += 1
		}
	} else {
		t.size += 1
		if wp == wLen {
			// One of the words in the Trie is a continuation of the
			// current word thus we need to split the current node.
			firstNode := NewNode(cVal[:cp], cNode.prev, true, nil)
			NewNode(cVal[cp:], firstNode, cNode.terminator, cNode.followers)
		} else {
			// This is a mismatch
			// The node from which we branch out. The first half of the original value/word
			splitNode := NewNode(cVal[:cp], cNode.prev, false, nil)
			// The node that we get from the original node with the second half of the original value
			NewNode(cVal[cp:], splitNode, cNode.terminator, cNode.followers)
			// The node we create to store the second half of the current word
			NewNode(rWord[wp:], splitNode, true, nil)
		}
	}
}

func (t *Trie) Remove(word string) {
	if word == "" {
		return
	}

	rWord := []rune(word)
	cNode, cp, wp := t.find(rWord)
	var toCompress *Node = nil

	if wp == len(rWord) && cp == len(cNode.val) && cNode.terminator {
		t.size -= 1
		// The word is in the trie
		fCount := len(cNode.followers)
		if fCount == 0 {
			// We can delete the current node and possibly compress the previous node
			pNode := cNode.prev
			delete(pNode.followers, cNode.val[0])

			if len(pNode.followers) == 1 && !pNode.terminator {
				// We can compress the only follower of the previous node into the previous one.
				toCompress = pNode
			}
		} else if fCount == 1 && !cNode.terminator {
			// We can compress the follower of the current node into the current one.
			toCompress = cNode
		} else {
			// Remove the terminator flag. Compression or removal of this node is not
			// possible since it has more than one follower.
			cNode.terminator = false
		}

		if toCompress != nil {
			var nNode *Node = nil // The only follower of the node we want to merge into its ancestor
			for _, el := range toCompress.followers {
				nNode = el
				break
			}
			// We create a new node where we concatenate the values in the 2 nodes.
			NewNode(merge(toCompress.val, nNode.val), toCompress.prev, nNode.terminator, nNode.followers)
		}
	}
}

func (t *Trie) Find(word string) bool {
	if word == "" {
		return true
	}

	rWord := []rune(word)
	cNode, cp, wp := t.find(rWord)
	return wp == len(rWord) && cp == len(cNode.val) && cNode.terminator
}

func (t *Trie) Iterate() <-chan string {
	c := make(chan string)
	go func() {
		stack := []*TrieIterState{
			&TrieIterState{t.root, []rune{}},
		}

		for len(stack) > 0 {
			cState := stack[0]
			stack = stack[1:]
			cNode := cState.node

			for _, nNode := range cNode.followers {
				stack = append(stack, &TrieIterState{nNode, merge(cState.val, nNode.val)})
			}
			if cNode.terminator {
				c <- string(cState.val)
			}
		}

		close(c)
	}()
	return c
}

type ShowState struct {
	indent int
	node   *Node
}

func (t *Trie) Show() {
	stack := []*ShowState{
		&ShowState{0, t.root},
	}

	for len(stack) != 0 {
		ls := len(stack)
		curr := stack[ls-1]
		stack = stack[:ls-1]

		fmt.Printf("%v: %v(%p, %p)\n", strings.Repeat("*", curr.indent), string(curr.node.val), curr.node, curr.node.prev)

		for _, next := range curr.node.followers {
			stack = append(stack, &ShowState{curr.indent + 1, next})
		}
	}
}

func (t *Trie) Len() int {
	return t.size
}

func main() {
	words := []string{
		"bubba",
		"bela",
		"asdasd",
		"asd",
	}
	t := NewTrie()

	for _, word := range words {
		t.Insert(word)
		for el := range t.Iterate() {
			fmt.Println(el)
		}
		fmt.Println("")
	}
}
