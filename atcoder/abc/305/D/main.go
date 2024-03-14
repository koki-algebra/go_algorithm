package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := NextInt(sc)
	a := make([]int, n)
	for i := range a {
		a[i] = NextInt(sc)
	}
	q := NextInt(sc)

	// sleeps[i]: a[i] 分までに何分寝たか
	sleeps := make([]int, n)
	for i := 1; i < n; i++ {
		if i%2 == 0 {
			sleeps[i] = sleeps[i-1] + a[i] - a[i-1]
		} else {
			sleeps[i] = sleeps[i-1]
		}
	}

	// x 分までに何分寝たか
	f := func(x int) int {
		j := sort.Search(n, func(i int) bool {
			return a[i] > x
		})

		if j == 0 {
			return 0
		} else if j == n {
			return sleeps[n-1]
		}

		return sleeps[j-1] + (sleeps[j]-sleeps[j-1])/(a[j]-a[j-1])*(x-a[j-1])
	}

	for i := 0; i < q; i++ {
		l := NextInt(sc)
		r := NextInt(sc)
		fmt.Fprintln(w, f(r)-f(l))
	}
}

func NextInt(sc *bufio.Scanner) int {
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
