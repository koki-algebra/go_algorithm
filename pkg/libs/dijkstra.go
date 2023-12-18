package libs

import "container/heap"

type Edge struct {
	To     int
	Weight int
}

type Graph [][]Edge

func Dijkstra(graph Graph, start int) []int {
	dist := make([]int, len(graph))
	for i := range dist {
		dist[i] = INF
	}
	dist[start] = 0

	// initialize priority queue
	pq := new(PriorityQueue[int])
	heap.Init(pq)
	heap.Push(pq, Item[int]{Value: start, Priority: dist[start]})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(Item[int])
		now := item.Value
		nowCost := item.Priority

		if dist[now] < nowCost {
			continue
		}

		for _, edge := range graph[now] {
			next := edge.To
			nextCost := nowCost + edge.Weight
			if Chmin(&dist[edge.To], nextCost) {
				heap.Push(pq, Item[int]{Value: next, Priority: nextCost})
			}
		}
	}

	return dist
}
