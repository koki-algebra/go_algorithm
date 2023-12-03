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

	n, m := NextInt(sc), NextInt(sc)
	a := make([]int, n)
	for i := range a {
		a[i] = NextInt(sc)
	}

	// 累積和
	sum := make([]int, n+1)
	sum[1] = a[0]
	for i := 1; i < n; i++ {
		sum[i+1] = sum[i] + a[i]
	}

	sumi := make([]int, n-m+1)
	for i := 0; i < m; i++ {
		sumi[0] += a[i] * (i + 1)
	}

	for i := 0; i < n-m; i++ {
		sumi[i+1] = sumi[i] + m*a[m+i] - (sum[m+i] - sum[i])
	}

	ans := -(1 << 60)
	for i := range sumi {
		Chmax(&ans, sumi[i])
	}
	fmt.Fprintln(w, ans)
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

func Chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}
