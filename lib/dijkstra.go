package lib

import "container/heap"

type Edge struct {
	To     int
	Weight int
}

type Graph [][]Edge

func Dijkstra(graph Graph, start int) []int {

	n := len(graph)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = INF
	}
	dist[start] = 0

	// priority queue
	pq := new(PriorityQueue[int])
	heap.Push(pq, Item[int]{Value: start, Priority: dist[start]})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(Item[int])
		now := item.Value
		nowCost := item.Priority
		if dist[now] < nowCost {
			continue
		}
		for _, e := range graph[now] {
			next := e.To
			nextCost := nowCost + e.Weight
			if Chmin(&dist[e.To], nextCost) {
				heap.Push(pq, Item[int]{Value: next, Priority: nextCost})
			}
		}
	}

	return dist
}
