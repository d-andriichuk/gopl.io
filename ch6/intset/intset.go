package intset

import (
	"bytes"
	"fmt"
)

//
// IntSet struct
type IntSet struct {
	words []uint64
}

//
// Has method
//
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

//
// Add method
//
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

//
// UnionWith method
//
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//
// AddAll method
//
func (s *IntSet) AddAll(args ...int) {
	for _, a := range args {
		s.Add(a)
	}
}

//
// Clear method
//
func (s *IntSet) Clear() {
	s.words = make([]uint64, 0)
}

//
// Copy method
//
func (s *IntSet) Copy() *IntSet {
	r := &IntSet{}

	for _, w := range s.words {
		r.words = append(r.words, w)
	}

	return r
}

//
// Remove method
//
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word >= len(s.words) {
		return
	}
	s.words[word] ^= 1 << bit
}

//
// Len method
//
func (s *IntSet) Len() int {
	var len int
	for _, w := range s.words {
		for w != 0 {
			w &= (w - 1)
			len++
		}
	}
	return len
}

//
// Elems method
//
func (s *IntSet) Elems() []int {
	var elems []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, 64*i+j)
			}
		}
	}
	return elems
}
