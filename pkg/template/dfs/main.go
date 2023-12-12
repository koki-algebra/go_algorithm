package main

import (
	"fmt"
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
	dfs(graph, seen, 0)
}

type Graph [][]int

func dfs(graph Graph, seen []bool, v int) {
	seen[v] = true

	fmt.Println(v + 1)

	for _, nv := range graph[v] {
		if !seen[nv] {
			dfs(graph, seen, nv)
		}
	}
}
