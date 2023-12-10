package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Graph [][]int

var (
	graph Graph
	seen  []bool
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n, m := NextInt(sc), NextInt(sc)
	graph = make(Graph, n)
	for i := 0; i < m; i++ {
		u, v := NextInt(sc)-1, NextInt(sc)-1
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	seen = make([]bool, n)

	// エッジの数がちょうど n-1 本かどうか
	if m != n-1 {
		fmt.Println("No")
		return
	}

	// 全てのノードの次数が 2 以下かどうか
	for i := 0; i < n; i++ {
		if len(graph[i]) > 2 {
			fmt.Println("No")
			return
		}
	}

	// 連結グラフかどうか
	count := 0 // 連結成分の個数
	for v := range graph {
		if !seen[v] {
			dfs(v)
			count++
		}
	}
	if count != 1 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
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

func dfs(v int) {
	seen[v] = true
	for _, nv := range graph[v] {
		if !seen[nv] {
			dfs(nv)
		}
	}
}
