package collections

import "fmt"

type Stack[T any] struct {
	items []T
}

// Push adds a value on top of the stack
func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

// Pop removes the top element of a stack and a boolean indicate whether it is succesful at retrival
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var empty T
		return empty, false

	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

// Peek gets the top element of a stack without removal
func (s *Stack[T]) Peek() T {
	var last T
	if len(s.items) == 0 {
		return last
	}
	return s.items[(len(s.items) - 1)]
}

func (s *Stack[T]) String() string {
	str := fmt.Sprint("[ ")
	for i := range s.items {
		str += fmt.Sprint(s.items[i])
		str += fmt.Sprint(", ")
	}
	str += fmt.Sprint(" ]")
	return str
}
