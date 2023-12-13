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
	a, b := make([]int, n), make([]int, n)
	for i := range a {
		a[i] = NextInt()
	}
	for i := range b {
		b[i] = NextInt()
	}

	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	ans := 0
	for i := 0; i < n; i++ {
		val := a[i] - b[i]
		if val < 0 {
			val = -val
		}
		ans += val
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
