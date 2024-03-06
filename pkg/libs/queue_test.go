package libs

import "testing"

func TestQueue(t *testing.T) {
	t.Parallel()

	q := NewQueue()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if got := q.Dequeue(); got != 1 {
		t.Errorf("expected value is 1, but got %d", got)
	}
	if got := q.Dequeue(); got != 2 {
		t.Errorf("expected value is 2, but got %d", got)
	}
	if got := q.Dequeue(); got != 3 {
		t.Errorf("expected value is 3, but got %d", got)
	}
}
