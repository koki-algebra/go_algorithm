package main

import (
	"fmt"

	"github.com/koki-algebra/go_algorithm/pkg/libs"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	graph := make(Graph, n)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		u--
		v--
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	seen := make([]bool, n)

	queue := libs.NewQueue()
	queue.Enqueue(0)

	bfs(graph, seen, queue)
}

type Graph [][]int

func bfs(graph Graph, seen []bool, queue *libs.Queue) {
	for !queue.IsEmpty() {
		v := queue.Dequeue().(int)
		seen[v] = true

		fmt.Println(v + 1)

		for _, nv := range graph[v] {
			if !seen[nv] {
				queue.Enqueue(nv)
			}
		}
	}
}
