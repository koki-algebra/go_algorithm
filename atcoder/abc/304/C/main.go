package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Graph [][]int

type Person struct {
	id         int
	x, y       int
	isInfected bool
}

var (
	sc     = bufio.NewScanner(os.Stdin)
	people []Person
	graph  Graph
)

func main() {
	sc.Split(bufio.ScanWords)
	n, d := NextInt(), NextFloat()
	people = make([]Person, n)
	for i := 0; i < n; i++ {
		x, y := NextInt(), NextInt()
		people[i] = Person{id: i, x: x, y: y}
	}
	people[0].isInfected = true

	graph = make(Graph, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && calcDist(people[i], people[j]) <= d {
				graph[people[i].id] = append(graph[people[i].id], people[j].id)
			}
		}
	}

	dfs(0)

	for i := range people {
		if people[i].isInfected {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func dfs(v int) {
	people[v].isInfected = true
	for _, nv := range graph[v] {
		if people[nv].isInfected {
			continue
		}
		dfs(nv)
	}
}

func calcDist(p1, p2 Person) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(float64(dx*dx + dy*dy))
}

func NextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func NextFloat() float64 {
	sc.Scan()
	n, err := strconv.ParseFloat(sc.Text(), 64)
	if err != nil {
		panic(err)
	}
	return n
}
