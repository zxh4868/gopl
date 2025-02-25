package main

import (
	"bytes"
	"fmt"
)

// IntSet是一个包含非负整数的集合
// 零值代表空的集合
type IntSet struct {
	units []uint64
}

func (s *IntSet) Has(x int) bool {
	idx, bit := x/64, x%64
	return idx < len(s.units) && s.units[idx]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	idx, bit := x/64, x%64
	for idx >= len(s.units) {
		s.units = append(s.units, 0)
	}
	s.units[idx] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for idx, tw := range t.units {
		if idx < len(s.units) {
			s.units[idx] |= tw
		} else {
			s.units = append(s.units, tw)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("{")
	for i, u := range s.units {
		if u == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if u&(1<<uint(j)) != 0 {
				if buf.Len() > 1 {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}

		}
	}
	buf.WriteString("}")
	return buf.String()
}

func (s *IntSet) Len() int {
	return len(s.units)
}

func (s *IntSet) Remove(x int) {
	idx, bit := x/64, x%64
	s.units[idx] &^= 1 << bit
}

func (s *IntSet) Clear() {
	s.units = s.units[:0]
}
func (s *IntSet) Copy() *IntSet {
	t := &IntSet{
		units: make([]uint64, len(s.units)),
	}
	copy(t.units, s.units)
	return t
}

func (s *IntSet) Elements() []uint64 {
	return s.units
}

func (s *IntSet) AddAll(ints ...int) {
	for _, i := range ints {
		s.Add(i)
	}
}

func (s *IntSet) IntersectionWith(t *IntSet) {
	
}

func (s *IntSet) DifferenceWith(t *IntSet) {

}

func (s *IntSet) SymDifference(t *IntSet) {

}

func main() {

	var x, y IntSet
	x.Add(12)
	y.Add(24)
	x.Add(13)
	x.Add(14)
	fmt.Println(&x)
	fmt.Println(x.String())
	fmt.Println(x)
}
