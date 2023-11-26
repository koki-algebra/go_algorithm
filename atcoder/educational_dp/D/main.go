package main

import "fmt"

func main() {
	var n, w int
	fmt.Scan(&n, &w)
	weight, value := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&weight[i], &value[i])
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, w+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= w; j++ {
			if j >= weight[i] {
				dp[i+1][j] = max(dp[i][j-weight[i]]+value[i], dp[i][j])
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}

	fmt.Println(dp[n][w])
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
