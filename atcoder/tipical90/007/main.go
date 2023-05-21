package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func main() {
	sc.Split(bufio.ScanWords)
	n := NextInt()
	a := make([]int, n)
	for i := range a {
		a[i] = NextInt()
	}

	sort.Ints(a)

	q := NextInt()
	for j := 0; j < q; j++ {
		b := NextInt()

		// binary search
		index := sort.Search(n, func(i int) bool { return a[i] >= b })

		if index > 0 && index < n {
			fmt.Println(min(abs(a[index]-b), abs(a[index-1]-b)))
		} else if index == 0 {
			fmt.Println(abs(a[index] - b))
		} else {
			fmt.Println(abs(a[index-1] - b))
		}
	}
}

func NextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func abs(v int) int {
	if v < 0 {
		v = -v
	}
	return v
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
