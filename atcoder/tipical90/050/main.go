package main

import "fmt"

const (
	MOD = 1000000007
)

func main() {
	var n, l int
	fmt.Scanf("%d %d", &n, &l)

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if i >= l {
			dp[i] = dp[i-1] + dp[i-l]
		} else {
			dp[i] = dp[i-1]
		}
		if dp[i] >= MOD {
			dp[i] -= MOD
		}
	}

	fmt.Println(dp[n])
}
