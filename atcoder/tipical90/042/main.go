package main

import "fmt"

const (
	MOD = 1000000007
)

func main() {
	var k int
	fmt.Scan(&k)
	if k%9 != 0 {
		fmt.Println(0)
		return
	}

	dp := make([]int, k+1)
	dp[0] = 1
	for i := 1; i <= k; i++ {
		b := min(i, 9)
		for j := 1; j <= b; j++ {
			dp[i] += dp[i-j]
			if dp[i] >= MOD {
				dp[i] -= MOD
			}
		}
	}
	fmt.Println(dp[k])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
