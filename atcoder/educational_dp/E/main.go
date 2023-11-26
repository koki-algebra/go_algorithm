package main

import "fmt"

const (
	INF = 1 << 60
)

func main() {
	var n, w int
	fmt.Scan(&n, &w)
	weight, value := make([]int, n), make([]int, n)
	valueMax := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&weight[i], &value[i])
		valueMax += value[i]
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, valueMax+1)
		for v := range dp[i] {
			dp[i][v] = INF
		}
	}
	dp[0][0] = 0

	for i := 0; i < n; i++ {
		for v := 0; v <= valueMax; v++ {
			if v >= value[i] {
				Chmin(&dp[i+1][v], dp[i][v-value[i]]+weight[i])
			}
			Chmin(&dp[i+1][v], dp[i][v])
		}
	}

	ans := 0
	for v := range dp[n] {
		if dp[n][v] <= w {
			ans = v
		}
	}
	fmt.Println(ans)
}

func Chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}
