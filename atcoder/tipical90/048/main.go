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
	n, k := NextInt(), NextInt()
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = NextInt(), NextInt()
	}

	c := make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		if i < n {
			c[i] = b[i]
		} else {
			c[i] = a[i-n] - b[i-n]
		}
	}

	// sort in descending order
	sort.Slice(c, func(i, j int) bool { return c[i] > c[j] })

	ans := 0
	for i := 0; i < k; i++ {
		ans += c[i]
	}
	fmt.Println(ans)
}

func NextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}
