package stack

import (
	"testing"
)

func Test(t *testing.T) {
	s := New(5)

	if s.Len() != 0 {
		t.Errorf("Length of an empty stack should be 0")
	}

	s.Push(1)

	if s.Len() != 1 {
		t.Errorf("Length should be 0")
	}

	if val, _ := s.Peek(); val != 1 {
		t.Errorf("Top item on the stack should be 1")
	}

	if val, _ := s.Pop(); val != 1 {
		t.Errorf("Top item should have been 1")
	}

	if s.Len() != 0 {
		t.Errorf("Stack should be empty")
	}

	s.Push(1)
	s.Push(2)

	if s.Len() != 2 {
		t.Errorf("Length should be 2")
	}

	if val, _ := s.Peek(); val != 2 {
		t.Errorf("Top of the stack should be 2")
	}
}
