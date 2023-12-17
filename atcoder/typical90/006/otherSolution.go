package main

import (
	"container/heap"
	"fmt"
)

func OtherSolution() {
	var (
		n, k int
		s    string
		next = 0
		ans  []byte
		h    = &CharHeap{}
	)
	fmt.Scan(&n, &k, &s)

	heap.Init(h)

	for i := range s {
		heap.Push(h, CharWithIndex{Char: s[i], Index: i})

		if n <= i+k {
			for {
				d := heap.Pop(h).(CharWithIndex)
				if next <= d.Index {
					next = d.Index + 1
					ans = append(ans, d.Char)
					break
				}
			}
		}
	}

	fmt.Println(string(ans))
}

type CharWithIndex struct {
	Char  byte
	Index int
}

type CharHeap []CharWithIndex

func (h CharHeap) Len() int {
	return len(h)
}

func (h CharHeap) Less(i, j int) bool {
	if h[i].Char == h[j].Char {
		return h[i].Index < h[j].Index
	}

	return h[i].Char < h[j].Char
}

func (h CharHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *CharHeap) Push(x any) {
	*h = append(*h, x.(CharWithIndex))
}

func (h *CharHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
