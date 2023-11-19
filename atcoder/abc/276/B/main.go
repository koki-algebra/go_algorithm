package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n, m := getInt(sc), getInt(sc)

	graph := make([]heap.Interface, n)
	for i := 0; i < n; i++ {
		graph[i] = &IntHeap{}
		heap.Init(graph[i])
	}

	for i := 0; i < m; i++ {
		a, b := getInt(sc)-1, getInt(sc)-1
		heap.Push(graph[a], b)
		heap.Push(graph[b], a)
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(w, "%d ", graph[i].Len())
		for g := graph[i]; g.Len() > 0; {
			fmt.Fprintf(w, "%d ", heap.Pop(g).(int)+1)
		}
		fmt.Fprintln(w)
	}
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
