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

var (
	sc = bufio.NewScanner(os.Stdin)
)

type Node struct {
	value    int
	distance int
}

type Edge struct {
	to     int
	weight int
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
	pq := make(PriorityQueue, 1)
	pq[0] = &Node{value: start, distance: dist[start]}
	heap.Init(&pq)

	for pq.Len() > 0 {
		now := heap.Pop(&pq).(*Node)
		if dist[now.value] < now.distance {
			continue
		}
		for _, e := range graph[now.value] {
			if chmin(&dist[e.to], now.distance+e.weight) {
				heap.Push(&pq, &Node{value: e.to, distance: now.distance + e.weight})
			}
		}
	}

	return dist
}

func chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	n, m := NextInt(), NextInt()
	graph := make(Graph, n)
	for i := 0; i < m; i++ {
		a, b, c := NextInt()-1, NextInt()-1, NextInt()
		graph[a] = append(graph[a], Edge{to: b, weight: c})
		graph[b] = append(graph[b], Edge{to: a, weight: c})
	}

	from1 := Dijkstra(graph, 0)
	fromN := Dijkstra(graph, n-1)
	for k := 0; k < n; k++ {
		fmt.Println(from1[k] + fromN[k])
	}
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance > pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*Node)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(*pq)
	node := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return node
}

func NextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}
