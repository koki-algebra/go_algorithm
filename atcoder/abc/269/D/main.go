package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Graph map[Point]map[Point]struct{}

type Point struct {
	X, Y int
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

var (
	graph = make(Graph)
	seen  = make(map[Point]bool)
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := NextInt(sc)

	for i := 0; i < n; i++ {
		x, y := NextInt(sc), NextInt(sc)
		point := NewPoint(x, y)

		graph[point] = make(map[Point]struct{})
		seen[point] = false
	}

	for v := range graph {
		x := v.X
		y := v.Y

		var (
			upperLeft  = NewPoint(x, y+1)
			upperRight = NewPoint(x+1, y+1)
			right      = NewPoint(x+1, y)
			lowerRight = NewPoint(x, y-1)
			lowerLeft  = NewPoint(x-1, y-1)
			left       = NewPoint(x-1, y)
		)

		if _, ok := graph[upperLeft]; ok {
			graph[v][upperLeft] = struct{}{}
			graph[upperLeft][v] = struct{}{}
		}
		if _, ok := graph[upperRight]; ok {
			graph[v][upperRight] = struct{}{}
			graph[upperRight][v] = struct{}{}
		}
		if _, ok := graph[right]; ok {
			graph[v][right] = struct{}{}
			graph[right][v] = struct{}{}
		}
		if _, ok := graph[lowerRight]; ok {
			graph[v][lowerRight] = struct{}{}
			graph[lowerRight][v] = struct{}{}
		}
		if _, ok := graph[lowerLeft]; ok {
			graph[v][lowerLeft] = struct{}{}
			graph[lowerLeft][v] = struct{}{}
		}
		if _, ok := graph[left]; ok {
			graph[v][left] = struct{}{}
			graph[left][v] = struct{}{}
		}
	}

	ans := 0
	for v := range graph {
		if !seen[v] {
			dfs(v)
			ans++
		}
	}
	fmt.Fprintln(w, ans)
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

func dfs(v Point) {
	seen[v] = true
	for nv := range graph[v] {
		if !seen[nv] {
			dfs(nv)
		}
	}
}
