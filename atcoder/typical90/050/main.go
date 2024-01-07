package main

import "fmt"

const MOD = 1000000007

func main() {
	var n, l int
	fmt.Scan(&n, &l)

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		add(&dp[i], dp[i-1])
		if i >= l {
			add(&dp[i], dp[i-l])
		}
	}

	fmt.Println(dp[n])
}

func add(a *int, b int) {
	*a += b
	if *a > MOD {
		*a -= MOD
	}
}
