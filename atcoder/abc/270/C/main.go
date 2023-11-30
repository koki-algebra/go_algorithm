package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

var (
	graph [][]int
	seen  []bool
	deq   = list.New()
	stop  = false
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n, x, y := NextInt(sc), NextInt(sc)-1, NextInt(sc)-1
	graph = make([][]int, n)
	for i := 0; i < n-1; i++ {
		u, v := NextInt(sc)-1, NextInt(sc)-1
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	seen = make([]bool, n)

	dfs(x, y)

	for e := deq.Front(); e != nil; e = e.Next() {
		v := e.Value.(int)
		fmt.Fprintf(w, "%d ", v+1)
	}
	fmt.Fprintln(w)
}

func dfs(v, to int) {
	if !stop {
		deq.PushBack(v)
	}
	if v == to {
		stop = true
	}
	seen[v] = true

	for _, nv := range graph[v] {
		if !seen[nv] {
			dfs(nv, to)
		}
	}

	if !stop {
		deq.Remove(deq.Back())
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
