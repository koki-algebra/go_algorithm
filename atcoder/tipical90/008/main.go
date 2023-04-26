package main

import "fmt"

const (
	MOD = 1000000007
	T   = "atcoder"
)

// add a に b を足して MOD を取る
func add(a *int, b int) {
	*a += b
	if *a >= MOD {
		*a %= MOD
	}
}

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&n)
	fmt.Scan(&s)

	// dp[i][j]: S の i 文字目までの部分文字列のうち, "atcoder" の j 文字目まで一致する方法の個数
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, len(T)+1)
	}
	dp[0][0] = 1

	// dp
	for i := 0; i < n; i++ {
		for j := 0; j <= len(T); j++ {
			// S[i] を選ばない場合
			add(&dp[i+1][j], dp[i][j])

			// s[i] を選ぶ場合
			if j < len(T) && s[i] == T[j] {
				add(&dp[i+1][j+1], dp[i][j])
			}
		}
	}

	fmt.Println(dp[n][len(T)])
}
