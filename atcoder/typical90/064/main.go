package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n, q := NextInt(sc), NextInt(sc)
	a := make([]int, n)
	for i := range a {
		a[i] = NextInt(sc)
	}

	ans := 0
	b := make([]int, n-1)
	for i := range b {
		b[i] = a[i+1] - a[i]
		ans += abs(b[i])
	}

	for i := 0; i < q; i++ {
		l, r, v := NextInt(sc)-1, NextInt(sc)-1, NextInt(sc)

		var before, after int
		if l > 0 {
			before += abs(b[l-1])
			b[l-1] += v
			after += abs(b[l-1])
		}
		if r < n-1 {
			before += abs(b[r])
			b[r] -= v
			after += abs(b[r])
		}
		ans += after - before

		fmt.Fprintln(w, ans)
	}
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
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
