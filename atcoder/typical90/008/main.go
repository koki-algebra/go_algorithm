package main

import "fmt"

const (
	MOD = 1000000007
	T   = "atcoder"
)

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n, &s)
	m := len(T)

	// dp[i][j]: 文字列 S の最初の i 文字から文字列かを抜き出してつなげる方法のうち, それが T の最初の j 文字まで一致するような方法の個数
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	dp[0][0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j <= m; j++ {
			add(&dp[i+1][j], dp[i][j])
			if j < m && s[i] == T[j] {
				add(&dp[i+1][j+1], dp[i][j])
			}
		}
	}

	fmt.Println(dp[n][m])
}

func add(a *int, b int) {
	*a += b
	if *a >= MOD {
		*a %= MOD
	}
}
