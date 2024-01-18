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

	n := NextInt(sc)
	xs, ys := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		xs[i], ys[i] = NextInt(sc), NextInt(sc)
	}

	sort.Ints(xs)
	sort.Ints(ys)

	ans := 0
	for i := 0; i < n; i++ {
		ans += abs(xs[i] - xs[n/2])
		ans += abs(ys[i] - ys[n/2])
	}

	fmt.Println(ans)
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
