package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := NextInt(sc)
	graph := make(Graph, n)
	for i := 0; i < n-1; i++ {
		u, v := NextInt(sc)-1, NextInt(sc)-1
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	s := 0

	dist := make([]int, n)
	for i := range dist {
		dist[i] = -1
	}
	dist[s] = 0

	// ノードから s から最も遠いノード v を求める
	dfs(graph, dist, s)
	v, _ := max(dist)

	dist = make([]int, n)
	for i := range dist {
		dist[i] = -1
	}
	dist[v] = 0

	// ノード v から最も通りノードを求める
	dfs(graph, dist, v)
	_, val := max(dist)

	fmt.Println(val + 1)
}

type Graph [][]int

func dfs(graph Graph, dist []int, v int) {
	for _, nv := range graph[v] {
		if dist[nv] == -1 {
			dist[nv] = dist[v] + 1
			dfs(graph, dist, nv)
		}
	}
}

func max(dist []int) (index, value int) {
	value = -1
	for i := range dist {
		if Chmax(&value, dist[i]) {
			index = i
		}
	}
	return
}

func Chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}

func NextInt(sc *bufio.Scanner) int {
	if sc.Scan() {
		n, err := strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}
		return n
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}

	return -1
}
