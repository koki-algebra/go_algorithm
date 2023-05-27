package main

import "fmt"

const (
	MOD = 100000
)

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	nxt := make([]int, MOD)
	for i := 0; i < MOD; i++ {
		nxt[i] = (i + digitsSum(i)) % MOD
	}

	timeStamps := make([]int, MOD)
	for i := range timeStamps {
		timeStamps[i] = -1
	}

	pos := n
	cnt := 0
	for timeStamps[pos] == -1 {
		timeStamps[pos] = cnt
		pos = nxt[pos]
		cnt++
	}

	cycle := cnt - timeStamps[pos]
	if k >= timeStamps[pos] {
		k = (k-timeStamps[pos])%cycle + timeStamps[pos]
	}

	ans := -1
	for i := 0; i < MOD; i++ {
		if timeStamps[i] == k {
			ans = i
		}
	}
	fmt.Println(ans)
}

// digitsSum 各桁の和を計算する O(logN)
func digitsSum(n int) (sum int) {
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return
}
