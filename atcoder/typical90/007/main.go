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
	// sort slice: O(NlogN)
	sort.Ints(a)

	q := NextInt(sc)
	for i := 0; i < q; i++ {
		b := NextInt(sc)

		// binary search: O(logN)
		index := sort.SearchInts(a, b)

		ans := 1 << 60
		if index >= 0 && index <= n-1 {
			Chmin(&ans, abc(a[index]-b))
		}
		if index >= 1 {
			Chmin(&ans, abc(a[index-1]-b))
		}

		fmt.Fprintln(w, ans)
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

func abc(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}
