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
	for i := 0; i < n; i++ {
		if seen[i] {
			continue
		}
		DFS(i)
	}

	// 帰りがけ順を降順に
	reverseSlice(history)

	// 強連結成分分解
	seen = make([]bool, n)

	for _, v := range history {
		if seen[v] {
			continue
		}
		components = append(components, []int{})
		ReverseDFS(v)
	}

	ans := 0
	for i := range components {
		n := len(components[i])
		ans += n * (n - 1) / 2
	}
	fmt.Println(ans)
}

func DFS(v int) {
	seen[v] = true
	for _, nv := range graph[v] {
		if seen[nv] {
			continue
		}
		DFS(nv)
	}
	history = append(history, v)
}

func ReverseDFS(v int) {
	seen[v] = true
	last := len(components) - 1
	components[last] = append(components[last], v)
	for _, nv := range reverseGraph[v] {
		if seen[nv] {
			continue
		}
		ReverseDFS(nv)
	}
}

func reverseSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func NextInt(sc *bufio.Scanner) int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}
