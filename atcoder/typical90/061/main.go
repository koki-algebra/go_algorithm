package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func main() {
	sc.Split(bufio.ScanWords)
	q := NextInt()

	deck := list.New()
	for i := 0; i < q; i++ {
		t, x := NextInt(), NextInt()
		switch t {
		case 1:
			deck.PushFront(x)
		case 2:
			deck.PushBack(x)
		case 3:
			fmt.Println(find(x-1, deck))
		}
	}
}

func find(i int, l *list.List) int {
	n := l.Len()
	if i < n/2 {
		index := 0
		for e := l.Front(); e != nil; e = e.Next() {
			if index == i {
				return e.Value.(int)
			}
			index++
		}
	} else {
		index := n - 1
		for e := l.Back(); e != nil; e = e.Prev() {
			if index == i {
				return e.Value.(int)
			}
			index--
		}
	}

	return -1
}

func NextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}
