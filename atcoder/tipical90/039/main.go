package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Graph [][]int

var (
	sc    = bufio.NewScanner(os.Stdin)
	graph Graph
	dp    []int
)

func main() {
	sc.Split(bufio.ScanWords)
	n := NextInt()
	graph = make(Graph, n)
	dp = make([]int, n)
	a := make([]int, n-1)
	b := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		a[i] = NextInt() - 1
		b[i] = NextInt() - 1
		graph[a[i]] = append(graph[a[i]], b[i])
		graph[b[i]] = append(graph[b[i]], a[i])
	}

	dfs(0, -1)

	ans := 0
	for i := 0; i < n-1; i++ {
		r := min(dp[a[i]], dp[b[i]])
		ans += r * (n - r)
	}

	fmt.Println(ans)
}

func dfs(pos, pre int) {
	dp[pos] = 1

	for _, v := range graph[pos] {
		if v == pre {
			continue
		}
		dfs(v, pos)
		dp[pos] += dp[v]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func NextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}
