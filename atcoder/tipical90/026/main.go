package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Color int

var (
	sc     = bufio.NewScanner(os.Stdin)
	whites []int
	blacks []int
)

const (
	unknown Color = iota
	black
	white
)

type Graph [][]int

func main() {
	sc.Split(bufio.ScanWords)
	n := NextInt()
	graph := make(Graph, n)
	for i := 0; i < n-1; i++ {
		a, b := NextInt()-1, NextInt()-1
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	dfs(graph, 0)

	if len(whites) > len(blacks) {
		sort.Slice(whites, func(i, j int) bool { return whites[i] < whites[j] })
		for i := 0; i < n/2; i++ {
			fmt.Printf("%d ", whites[i]+1)
		}
	} else {
		sort.Slice(blacks, func(i, j int) bool { return blacks[i] < blacks[j] })
		for i := 0; i < n/2; i++ {
			fmt.Printf("%d ", blacks[i]+1)
		}
	}
	fmt.Print("\n")
}

func dfs(graph Graph, s int) bool {
	colors := make([]Color, len(graph))
	for i := range colors {
		colors[i] = unknown
	}
	colors[s] = white
	st := NewStack()
	st.Push(s)

	whites = append(whites, s)

	for !st.IsEmpty() {
		v := st.Pop().(int)
		for _, nv := range graph[v] {
			if colors[nv] != unknown {
				if colors[nv] == colors[v] {
					return false
				}
			} else {
				st.Push(nv)
				colors[nv] = invertColor(colors[v])

				if colors[nv] == white {
					whites = append(whites, nv)
				} else {
					blacks = append(blacks, nv)
				}
			}
		}
	}

	return true
}

func invertColor(color Color) Color {
	if color == white {
		return black
	} else if color == black {
		return white
	}
	return unknown
}

func NextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

type Stack struct {
	l *list.List
}

func NewStack() *Stack {
	return &Stack{
		l: list.New(),
	}
}

func (s *Stack) Push(v interface{}) {
	s.l.PushFront(v)
}

func (s *Stack) Pop() interface{} {
	if s.l.Len() == 0 {
		panic("stack is empty")
	}
	e := s.l.Front()
	return s.l.Remove(e)
}

func (s *Stack) IsEmpty() bool {
	return s.l.Len() == 0
}
