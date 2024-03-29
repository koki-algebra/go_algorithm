package libs

type Item[T any] struct {
	Value    T
	Priority int
}

type PriorityQueue[T any] []Item[T]

func (pq PriorityQueue[T]) Len() int {
	return len(pq)
}

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue[T]) Push(x any) {
	*pq = append(*pq, x.(Item[T]))
}

func (pq *PriorityQueue[T]) Pop() any {
	n := len(*pq)
	res := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return res
}
