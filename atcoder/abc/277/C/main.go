package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	seen  map[int]struct{}
	graph *Graph
	max   int
)

type Graph struct {
	Nodes map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(a, b int) {
	g.Nodes[a] = append(g.Nodes[a], b)
	g.Nodes[b] = append(g.Nodes[b], a)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := getInt(sc)

	graph = NewGraph()

	for i := 0; i < n; i++ {
		a, b := getInt(sc), getInt(sc)
		graph.AddEdge(a, b)
	}

	seen = make(map[int]struct{})

	dfs(1)

	fmt.Println(max)
}

func getInt(sc *bufio.Scanner) int {
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

func dfs(v int) {
	seen[v] = struct{}{}
	for _, nv := range graph.Nodes[v] {
		if _, ok := seen[nv]; ok {
			continue
		}
		dfs(nv)
	}
	chmax(&max, v)
}

func chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}
