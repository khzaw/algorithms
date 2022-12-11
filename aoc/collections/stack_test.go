package collections

import (
	"fmt"
	"testing"
)

func Test_Stack(t *testing.T) {
	s := Stack[int]{}
	empty := s.Peek()
	fmt.Println(empty)
}
