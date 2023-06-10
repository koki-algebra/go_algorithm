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

	sleeps := make([]int, n)
	for i := 1; i < n; i++ {
		if i%2 == 0 {
			sleeps[i] = sleeps[i-1] + a[i] - a[i-1]
		} else {
			sleeps[i] = sleeps[i-1]
		}
	}

	q := NextInt()
	for i := 0; i < q; i++ {
		ans := 0
		l, r := NextInt(), NextInt()
		indexL := sort.Search(n, func(i int) bool {
			return a[i] >= l
		})
		indexR := sort.Search(n, func(i int) bool {
			return a[i] >= r
		})

		if indexL%2 == 0 {
			ans += a[indexL] - l
		}
		if indexR%2 == 0 {
			ans += r - a[indexR-1]
		}
		ans += sleeps[indexR-1] - sleeps[indexL]

		fmt.Println(ans)
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
