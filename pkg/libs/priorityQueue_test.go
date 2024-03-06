package libs

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	t.Parallel()

	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue[string], len(items))
	i := 0
	for value, priority := range items {
		pq[i] = Item[string]{
			Value:    value,
			Priority: priority,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item
	item := Item[string]{
		Value:    "orange",
		Priority: 5,
	}
	heap.Push(&pq, item)

	exps := []Item[string]{
		{Value: "apple", Priority: 2},
		{Value: "banana", Priority: 3},
		{Value: "pear", Priority: 4},
		{Value: "orange", Priority: 5},
	}

	for _, exp := range exps {
		got := heap.Pop(&pq).(Item[string])
		assert.Equal(t, exp.Value, got.Value)
		assert.Equal(t, exp.Priority, got.Priority)
	}
}
