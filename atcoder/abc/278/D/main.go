package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := getInt(sc)
	a := make([]int, n)
	for i := range a {
		a[i] = getInt(sc)
	}
	q := getInt(sc)

	// 差分が 0 でないインデックス
	l := list.New()
	for i := 0; i < n; i++ {
		l.PushBack(i)
	}

	base := 0

	for j := 0; j < q; j++ {
		t := getInt(sc)
		switch t {
		case 1:
			x := getInt(sc)
			for l.Len() != 0 {
				i := l.Back().Value.(int)
				a[i] = 0
				l.Remove(l.Back())
			}
			base = x
		case 2:
			i, x := getInt(sc)-1, getInt(sc)
			a[i] += x
			l.PushBack(i)
		case 3:
			i := getInt(sc) - 1
			fmt.Fprintln(w, a[i]+base)
		}
	}

	if err := w.Flush(); err != nil {
		panic(err)
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
