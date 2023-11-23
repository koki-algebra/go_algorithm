package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i], &b[i])
	}

	// dp[i][j]: A[i] 円の硬貨 B[i] 枚を用いてちょうど j 円を払えるならば　 true, 払えないならば false
	// dp[i][j] = dp[i-1][j] || dp[i-1][j - A[i]] || dp[i-1][j - 2*A[i]] || ... dp[i-1][j - B[i]*A[i]]
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, x+1)
	}
	dp[0][0] = true

	for i := 0; i < n; i++ {
		for j := 0; j <= x; j++ {
			for k := 0; k <= b[i]; k++ {
				if j >= k*a[i] {
					if dp[i][j-k*a[i]] {
						dp[i+1][j] = true
					}
				}
			}
		}
	}

	if dp[n][x] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
