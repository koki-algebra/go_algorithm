package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

const (
	INF = 1 << 60
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n, m := NextInt(sc), NextInt(sc)
	graph := make(Graph, n)
	for i := 0; i < m; i++ {
		a, b, c := NextInt(sc)-1, NextInt(sc)-1, NextInt(sc)
		graph[a] = append(graph[a], Edge{To: b, Weight: c})
		graph[b] = append(graph[b], Edge{To: a, Weight: c})
	}

	from1 := Dijkstra(graph, 0)
	fromN := Dijkstra(graph, n-1)
	for k := 0; k < n; k++ {
		fmt.Println(from1[k] + fromN[k])
	}
}

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
	pq := new(PriorityQueue)
	heap.Push(pq, Item{Value: start, Priority: dist[start]})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(Item)
		now := item.Value
		nowCost := item.Priority
		if dist[now] < nowCost {
			continue
		}
		for _, e := range graph[now] {
			next := e.To
			nextCost := nowCost + e.Weight
			if Chmin(&dist[e.To], nextCost) {
				heap.Push(pq, Item{Value: next, Priority: nextCost})
			}
		}
	}

	return dist
}

func Chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

type Item struct {
	Value    int
	Priority int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	res := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return res
}

func NextInt(sc *bufio.Scanner) int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}
