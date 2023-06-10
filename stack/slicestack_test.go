package stack_test

import (
	"github.com/rupertw/goat/stack"
	"testing"
)

func TestSliceStack(t *testing.T) {
	s := stack.NewSliceStack[int](0)

	if !s.IsEmpty() {
		t.Error(`s.IsEmpty() = false`)
	}

	e, b := s.Peek()
	if b {
		t.Errorf(`s.Peek() returns: %v`, e)
	}

	s.Push(5)

	e, b = s.Pop()
	if !b || e != 5 {
		t.Error(`s.Pop() fails`)
	}

	l := s.Len()
	if l > 0 {
		t.Errorf(`s.Len() returns: %v`, l)
	}
}
