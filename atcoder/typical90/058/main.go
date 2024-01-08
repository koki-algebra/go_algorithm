package main

import (
	"fmt"
)

const MOD = 100000

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	nexts := make([]int, MOD)
	timestamps := make([]int, MOD)
	for i := 0; i < MOD; i++ {
		nexts[i] = (i + digitsSum(i)) % MOD
		timestamps[i] = -1
	}

	pos := n
	cnt := 0
	for timestamps[pos] == -1 {
		timestamps[pos] = cnt
		pos = nexts[pos]
		cnt++
	}

	cycle := cnt - timestamps[pos]
	if k >= timestamps[pos] {
		k = (k-timestamps[pos])%cycle + timestamps[pos]
	}

	ans := -1
	for i := 0; i < MOD; i++ {
		if timestamps[i] == k {
			ans = i
		}
	}
	fmt.Println(ans)
}

// digitsSum calculates the sum of each digit O(logN)
func digitsSum(x int) (sum int) {
	for x > 0 {
		sum += x % 10
		x /= 10
	}
	return
}
