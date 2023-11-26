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
	n := NextInt(sc)
	a, b, c := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i], c[i] = NextInt(sc), NextInt(sc), NextInt(sc)
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 3)
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= 2; j++ {
			switch j {
			case 0:
				dp[i][j] = max(dp[i-1][1], dp[i-1][2]) + a[i-1]
			case 1:
				dp[i][j] = max(dp[i-1][2], dp[i-1][0]) + b[i-1]
			case 2:
				dp[i][2] = max(dp[i-1][0], dp[i-1][1]) + c[i-1]
			}
		}
	}

	ans := max(dp[n][0], dp[n][1])
	ans = max(dp[n][2], ans)
	fmt.Println(ans)
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

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
