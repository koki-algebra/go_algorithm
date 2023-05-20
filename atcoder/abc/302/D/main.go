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
	n, m, d := NextInt(), NextInt(), NextInt()
	a := make([]int, n)
	for i := range a {
		a[i] = NextInt()
	}
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })

	max := -1
	for j := 0; j < m; j++ {
		b := NextInt()
		index := sort.Search(n, func(i int) bool {
			return a[i] <= b+d
		})

		if index == n {
			continue
		}

		if a[index] >= b-d {
			chmax(&max, a[index]+b)
		}
	}
	fmt.Println(max)
}

func chmax(a *int, b int) {
	if *a < b {
		*a = b
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
