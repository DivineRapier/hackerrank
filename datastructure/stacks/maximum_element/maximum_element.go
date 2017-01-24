package main


import "fmt"

// Stack 类型
type Stack struct {
	Element []int
	max     []int
	count   int
}

// Push ...
func (s *Stack) Push(e int) {
	if s.count <= len(s.Element) {
		s.Element = append(s.Element, e)
	} else {
		s.Element[s.count] = e
	}
	max := e

	if s.count != 0 && (max < s.max[s.count-1]) {
		max = s.max[s.count-1]
	}
	if s.count < len(s.max) {
		s.max[s.count] = max
	} else {
		s.max = append(s.max, max)
	}

	s.count++
}

// Delete ...
func (s *Stack) Delete() {
	if s.count == 0 {
		return
	}
	s.count--
}

// Print ...
func (s *Stack) Print() {
	fmt.Println(s.max[s.count-1])
}
