package libs

import "testing"

func TestStack(t *testing.T) {
	t.Parallel()

	s := NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if got := s.Pop(); got != 3 {
		t.Errorf("expected value is 3, but got %d", got)
	}
	if got := s.Pop(); got != 2 {
		t.Errorf("expected value is 2, but got %d", got)
	}
	if got := s.Pop(); got != 1 {
		t.Errorf("expected value is 1, but got %d", got)
	}
}
