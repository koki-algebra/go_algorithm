package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

const (
	INF = 1 << 60
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := NextInt(sc)
	graph := make([][]int, n)
	for i := 0; i < n-1; i++ {
		a, b := NextInt(sc)-1, NextInt(sc)-1
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	dist0 := dfs(graph, 0)
	index, _ := max(dist0)
	dist := dfs(graph, index)
	_, val := max(dist)
	fmt.Println(val + 1)
}

func max(dist []int) (index, val int) {
	for i := range dist {
		if val < dist[i] {
			val = dist[i]
			index = i
		}
	}

	return
}

func dfs(graph [][]int, s int) []int {
	n := len(graph)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = -1
	}
	dist[s] = 0

	st := NewStack()
	st.Push(s)

	for !st.IsEmpty() {
		v := st.Pop()
		for _, nv := range graph[v] {
			if dist[nv] == -1 {
				st.Push(nv)
				dist[nv] = dist[v] + 1
			}
		}
	}

	return dist
}

type Stack struct {
	l *list.List
}

func NewStack() *Stack {
	return &Stack{
		l: list.New(),
	}
}

func (s *Stack) Push(v int) {
	s.l.PushFront(v)
}

func (s *Stack) Pop() int {
	e := s.l.Front()
	return s.l.Remove(e).(int)
}

func (s *Stack) IsEmpty() bool {
	return s.l.Len() == 0
}

func NextInt(sc *bufio.Scanner) int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func NextLine(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
