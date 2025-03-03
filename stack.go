package main

import (
	"fmt"
)

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
func (s *Stack) Push(data int) {
	*s = append(*s, data)
}
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	id := len(*s) - 1
	val := (*s)[id]
	*s = (*s)[:id]
	return val, true
}
func main() {
	var s Stack
	val, ok := s.Pop()
	fmt.Println(val, ok)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	for !s.IsEmpty() {
		val, ok = s.Pop()
		if ok {
			fmt.Println(val)
		}

	}
}
