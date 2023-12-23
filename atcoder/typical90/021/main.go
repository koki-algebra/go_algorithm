package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Graph [][]int

var (
	graph        Graph
	reverseGraph Graph
	seen         []bool
	components   [][]int
	history      []int
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n, m := NextInt(sc), NextInt(sc)
	graph = make(Graph, n)
	reverseGraph = make(Graph, n)
	for i := 0; i < m; i++ {
		a, b := NextInt(sc)-1, NextInt(sc)-1
		graph[a] = append(graph[a], b)
		reverseGraph[b] = append(reverseGraph[b], a)
	}

	// DFS
	seen = make([]bool, n)
	for i := range seen {
		if !seen[i] {
			dfs(i)
		}
	}

	// DFS
	seen = make([]bool, n)
	for i := len(history) - 1; i >= 0; i-- {
		if !seen[history[i]] {
			components = append(components, []int{})
			reverseDfs(history[i])
		}
	}

	ans := 0
	for i := range components {
		n := len(components[i])
		ans += n * (n - 1) / 2
	}
	fmt.Println(ans)
}

func dfs(v int) {
	seen[v] = true
	for _, nv := range graph[v] {
		if !seen[nv] {
			dfs(nv)
		}
	}

	history = append(history, v)
}

func reverseDfs(v int) {
	seen[v] = true
	last := len(components) - 1
	components[last] = append(components[last], v)

	for _, nv := range reverseGraph[v] {
		if !seen[nv] {
			reverseDfs(nv)
		}
	}
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
