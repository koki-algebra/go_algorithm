package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type (
	Color int
	Graph [][]int
)

const (
	UNKNOWN Color = iota
	BLACK
	WHITE
)

var (
	colors         []Color
	whites, blacks []int
	graph          Graph
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := NextInt(sc)
	graph = make([][]int, n)
	for i := 0; i < n-1; i++ {
		a, b := NextInt(sc)-1, NextInt(sc)-1
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	start := 0
	colors = make([]Color, n)
	colors[start] = WHITE
	whites = append(whites, start)

	dfs(start)

	if len(whites) > len(blacks) {
		sort.Ints(whites)
		for i := 0; i < n/2; i++ {
			fmt.Fprintf(w, "%d ", whites[i]+1)
		}
	} else {
		sort.Ints(blacks)
		for i := 0; i < n/2; i++ {
			fmt.Fprintf(w, "%d ", blacks[i]+1)
		}
	}
	fmt.Fprintln(w)
}

func dfs(v int) {
	for _, nv := range graph[v] {
		if colors[nv] == UNKNOWN {
			colors[nv] = colors[v].Invert()

			if colors[nv] == WHITE {
				whites = append(whites, nv)
			} else if colors[nv] == BLACK {
				blacks = append(blacks, nv)
			}

			dfs(nv)
		}
	}
}

func (c Color) Invert() Color {
	switch c {
	case WHITE:
		return BLACK
	case BLACK:
		return WHITE
	default:
		return UNKNOWN
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
