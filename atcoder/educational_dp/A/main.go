package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := NextInt(sc)
	h := make([]int, n)
	for i := range h {
		h[i] = NextInt(sc)
	}

	// dp[i]: 足場 i に辿り着くまでに支払うコストの総和の最小値 (i = 0, 1, ..., n-1)
	dp := make([]int, n)
	dp[1] = abc(h[1] - h[0])
	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-2]+abc(h[i]-h[i-2]), dp[i-1]+abc(h[i]-h[i-1]))
	}

	fmt.Println(dp[n-1])
}

func abc(a int) int {
	return int(math.Abs(float64(a)))
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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
